#!/usr/bin/env bash
#
# Verifica que las specs, el indice y el registro de las puertas digan lo mismo.
# Corre en CI, asi que no usa gh ni toca la red.
#
# Uso: scripts/coherencia.sh

set -u

INDICE="docs/specs/README.md"
FALLAS=0

fallo() {
  echo "  FALLA: $1"
  FALLAS=$((FALLAS + 1))
}

# Lo primero: marcas de conflicto sin resolver. Un merge mal resuelto puede dejarlas
# commiteadas, y el resto de los chequeos no las ve porque el archivo "parece" correcto.
con_marcas=$(grep -rl '^<<<<<<< \|^>>>>>>> \|^======= *$' --include='*.md' --include='*.go' --include='*.sh' --include='*.yml' . 2>/dev/null | grep -v '^./.git/' || true)
if [ -n "$con_marcas" ]; then
  echo "Hay marcas de conflicto sin resolver en:"
  echo "$con_marcas" | sed 's/^/  /'
  exit 1
fi

echo "Chequeando coherencia de las specs..."

# El Sprint 0 queda afuera a proposito: sus documentos registran decisiones de
# configuracion, no historias, y no llevan puertas. Ver docs/specs/sprint-0/README.md.
for spec in docs/specs/sprint-[1-9]*/[0-9]*.md; do
  [ -e "$spec" ] || continue
  nombre=$(basename "$spec")
  echo "- $spec"

  estado=$(grep -m1 '^- \*\*Estado:\*\*' "$spec" | sed 's/^- \*\*Estado:\*\* *//')
  [ -n "$estado" ] || fallo "$nombre no declara Estado"

  grep -q '^- \*\*Responsable:\*\*' "$spec" || fallo "$nombre no declara Responsable"
  grep -q '^- \*\*Historia:\*\*' "$spec" || fallo "$nombre no declara Historia"

  # si una puerta figura aprobada en el encabezado, su linea no puede estar pendiente
  for n in 1 2 3; do
    linea=$(grep -A1 "^> \*\*Puerta $n\*\*" "$spec" | grep '^> Aprobó:' || true)
    aprobada_en_linea="no"
    case "$linea" in
      *"[pendiente]"*|"") aprobada_en_linea="no" ;;
      *) aprobada_en_linea="si" ;;
    esac
    case "$estado" in
      *"Puerta $n aprobada"*)
        [ "$aprobada_en_linea" = "si" ] || \
          fallo "$nombre dice 'Puerta $n aprobada' en el encabezado pero la linea de la Puerta $n sigue en [pendiente]"
        ;;
    esac
  done

  # la fase siguiente no puede estar escrita si la puerta anterior no se aprobo
  if grep -q '^### Enfoque' "$spec"; then
    linea1=$(grep -A1 '^> \*\*Puerta 1\*\*' "$spec" | grep '^> Aprobó:' || true)
    case "$linea1" in
      *"[pendiente]"*|"") fallo "$nombre tiene la Fase 2 escrita pero la Puerta 1 no esta aprobada" ;;
    esac
  fi

  # el indice tiene que mencionar la spec
  grep -q "$nombre" "$INDICE" || fallo "$nombre no aparece en $INDICE"
done

# documentos de docs/ que nadie enlaza: se vuelven invisibles y quedan sin mantener
echo "- documentos enlazados"
for doc in docs/*.md docs/actas/*.md docs/specs/*.md; do
  [ -e "$doc" ] || continue
  nombre=$(basename "$doc")
  carpeta=$(dirname "$doc")/
  # alcanza con que se enlace el archivo, o la carpeta que lo contiene
  if ! grep -rq "$nombre" README.md CLAUDE.md AGENTS.md docs/specs/README.md 2>/dev/null &&
     ! grep -rq "($carpeta)" README.md CLAUDE.md AGENTS.md 2>/dev/null; then
    fallo "$doc no esta enlazado desde README.md, CLAUDE.md, AGENTS.md ni el indice de specs"
  fi
done

# specs mencionadas en el indice que ya no existen
while read -r referencia; do
  [ -e "docs/specs/$referencia" ] || fallo "$INDICE menciona $referencia, que no existe"
done < <(grep -o 'sprint-[0-9]*/[0-9][^)]*\.md' "$INDICE" | sort -u)

echo
if [ "$FALLAS" -eq 0 ]; then
  echo "Todo coherente."
else
  echo "$FALLAS problema(s) de coherencia."
  exit 1
fi
