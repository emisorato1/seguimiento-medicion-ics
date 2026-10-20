package web

import (
	"fmt"
	"net/http"

	"github.com/emisorato1/seguimiento-medicion-ics/internal/domain"
)

// roleOptions son los roles que se ofrecen para elegir, con el nombre que les da el dominio
// (RN-11, D10). El dominio convierte el texto elegido con ParseRole.
func roleOptions() []string {
	return []string{domain.RoleAgileEnabler.String(), domain.RoleProductBuilder.String()}
}

// memberFields son los campos del formulario de integrante.
var memberFields = []string{"nombre", "rol"}

// listMembers muestra los integrantes del proyecto. Sin integrantes muestra que no hay y
// ofrece agregar el primero (RN-7).
func (s *Server) listMembers(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.render(w, r, http.StatusOK, "integrantes.html", page{
		Title:   "Integrantes de " + project.Name(),
		Project: project,
		Section: "integrantes",
		Data:    project,
	})
}

// renderMemberForm muestra el formulario de integrante con lo escrito y sus errores.
func (s *Server) renderMemberForm(w http.ResponseWriter, r *http.Request, status int, project *domain.Project, f *form) {
	data := struct {
		ProjectID int
		Roles     []string
		Form      *form
	}{
		ProjectID: project.ID(),
		Roles:     roleOptions(),
		Form:      f,
	}
	s.render(w, r, status, "integrante-form.html", page{
		Title:   "Agregar integrante",
		Project: project,
		Section: "integrantes",
		Data:    data,
	})
}

// newMember muestra el formulario para agregar un integrante.
func (s *Server) newMember(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	s.renderMemberForm(w, r, http.StatusOK, project, newForm())
}

// createMember agrega el integrante con el nombre y el rol elegido. El rol se convierte en
// el dominio y, si no es valido, lo rechaza AddMember en su propio orden: asi el mensaje es
// siempre el primero que informa el nucleo (punto 7 de la puerta).
func (s *Server) createMember(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}

	f := readForm(r, memberFields...)
	role, _ := domain.ParseRole(f.Values["rol"])
	member, err := project.AddMember(f.Values["nombre"], role)
	if err != nil {
		f.reject(err)
		s.renderMemberForm(w, r, http.StatusUnprocessableEntity, project, f)
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, "Se agregó a "+member.Name())
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d/integrantes", project.ID()), http.StatusSeeOther)
}

// findMember busca el integrante con el nombre que viene en el pedido. Compara el nombre
// exacto, porque llega del enlace que arma la propia lista; que no exista lo informa el
// dominio con su error.
func findMember(project *domain.Project, name string) (*domain.Member, error) {
	for _, member := range project.Members() {
		if member.Name() == name {
			return member, nil
		}
	}
	return nil, domain.ErrMemberNotFound
}

// confirmRemoveMember pide confirmar antes de quitar un integrante (RN-10, D9). Mostrar la
// confirmacion no cambia nada: recien el boton que confirma envia el pedido.
func (s *Server) confirmRemoveMember(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	member, err := findMember(project, r.FormValue("nombre"))
	if err != nil {
		s.notFound(w, r, project, err)
		return
	}

	s.render(w, r, http.StatusOK, "confirmar.html", page{
		Title:   "Quitar integrante",
		Project: project,
		Section: "integrantes",
		Data: confirmation{
			Question: fmt.Sprintf("¿Quitar a %s del proyecto?", member.Name()),
			Detail:   "Deja de figurar entre los integrantes. Se puede volver a agregar después.",
			Action:   fmt.Sprintf("/proyectos/%d/integrantes/quitar", project.ID()),
			Fields:   map[string]string{"nombre": member.Name()},
			Confirm:  "Sí, quitar",
			Cancel:   fmt.Sprintf("/proyectos/%d/integrantes", project.ID()),
		},
	})
}

// removeMember quita el integrante, despues de que se confirmo.
func (s *Server) removeMember(w http.ResponseWriter, r *http.Request) {
	project, err := s.findProject(r)
	if err != nil {
		s.notFound(w, r, nil, err)
		return
	}
	name := r.FormValue("nombre")
	if err := project.RemoveMember(name); err != nil {
		s.notFound(w, r, project, err)
		return
	}

	s.saver.Save(s.registry)
	setFlash(w, "Se quitó a "+name)
	http.Redirect(w, r, fmt.Sprintf("/proyectos/%d/integrantes", project.ID()), http.StatusSeeOther)
}
