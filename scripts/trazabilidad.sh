#!/usr/bin/env bash
#
# Genera docs/trazabilidad.md: el recorrido historia -> spec -> criterios -> escenario BDD ->
# test -> codigo de cada historia que ya tiene tests en el repositorio. Es lo que piden en la
# defensa: eligen una funcionalidad y hay que poder recorrerla entera.
#
# De donde sale cada eslabon:
#   - historia, criterios y escenarios: la Fase 1 de cada spec (docs/specs/sprint-N/NN-*.md)
#   - archivos de test y de codigo de la historia: la tabla "Trazabilidad" de la spec
#   - test de cada escenario: el comentario "// Escenario:" (o "// Esquema del escenario:")
#     arriba del test, como pide AGENTS.md
#   - codigo de cada test: el historial de TDD. Se busca el commit que agrego el test (RED) y
#     el commit siguiente de su rama que toca codigo (GREEN): los archivos que cambia el
#     GREEN son el codigo que hizo pasar el test
#
# Los huecos se marcan, no se tapan: escenario sin test, test sin escenario, test sin GREEN,
# test y codigo en el mismo commit.
#
# Necesita el historial completo de git: en un clon superficial no encuentra los commits.
#
# Uso: scripts/trazabilidad.sh            escribe docs/trazabilidad.md
#      scripts/trazabilidad.sh --stdout   lo muestra sin escribir nada

set -eu

cd "$(git rev-parse --show-toplevel)"

SALIDA="docs/trazabilidad.md"
TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

# --- 1. Specs: historia, titulo, escenarios, criterios y archivos de la tabla ------------

for spec in docs/specs/sprint-[1-9]*/[0-9]*.md; do
  [ -e "$spec" ] || continue
  awk -v spec="$spec" -v dir="$TMP" '
    function limpiar(s) { sub(/^[ \t]+/, "", s); sub(/[ \t]+$/, "", s); return s }
    # Contenido de los `...` de una fila de la tabla, separado por espacios
    function entre_comillas(s,    r, i, partes, n) {
      r = ""; n = split(s, partes, "`")
      for (i = 2; i <= n; i += 2) r = r (r == "" ? "" : " ") partes[i]
      return r
    }
    /^# Spec:/ && titulo == "" { titulo = limpiar(substr($0, 8)) }
    /^- \*\*Historia:\*\*/ && historia == "" {
      if (match($0, /#[0-9]+/)) historia = substr($0, RSTART + 1, RLENGTH - 1)
    }
    /^## Fase 2/ { fase1_terminada = 1 }
    /^## Trazabilidad/ { en_tabla = 1 }
    /^### / { en_criterios = ($0 ~ /^### Criterios de aceptaci/) }
    en_criterios && /^- \[[ xX]\]/ {
      n_crit++; crit[n_crit] = limpiar(substr($0, 6)); next
    }
    en_criterios && n_crit > 0 && /^[ \t]+[^ \t]/ { crit[n_crit] = crit[n_crit] " " limpiar($0); next }
    !fase1_terminada && /^[ \t]*(Esquema del escenario|Escenario):/ {
      t = $0; sub(/^[ \t]*(Esquema del escenario|Escenario):[ \t]*/, "", t)
      n_esc++; esc[n_esc] = limpiar(t); esc_linea[n_esc] = FNR
    }
    en_tabla && /^\| *Tests *\|/ { tests = entre_comillas($0) }
    en_tabla && /^\| *C(o|ó)digo Go *\|/ { codigo = entre_comillas($0) }
    END {
      if (historia == "") exit
      printf "%s\t%s\t%s\t%s\t%s\n", historia, spec, titulo, tests, codigo >> (dir "/specs.tsv")
      for (i = 1; i <= n_esc; i++)
        printf "%s\t%s\t%s\n", historia, esc[i], esc_linea[i] >> (dir "/escenarios.tsv")
      for (i = 1; i <= n_crit; i++)
        printf "%s\t%s\n", historia, crit[i] >> (dir "/criterios.tsv")
    }
  ' "$spec"
done
touch "$TMP/specs.tsv" "$TMP/escenarios.tsv" "$TMP/criterios.tsv"

# --- 2. Tests: nombre, archivo, linea, escenario que cita y reglas que cubre -------------

find . -name '*_test.go' -not -path './.git/*' | sed 's|^\./||' | sort > "$TMP/archivos-test.txt"
: > "$TMP/tests.tsv"
while IFS= read -r archivo; do
  awk -v archivo="$archivo" '
    /^\/\// { bloque = bloque "\n" $0; next }
    /^func Test[A-Za-z0-9_]*\(/ {
      nombre = $2; sub(/\(.*/, "", nombre)
      n_esc = 0; cubre = ""; declara = 0
      n = split(bloque, lineas, "\n")
      for (i = 1; i <= n; i++) {
        l = lineas[i]
        # Un test puede cubrir mas de un escenario: una linea "// Escenario:" por cada uno
        if (l ~ /^\/\/ (Esquema del escenario|Escenario):/) {
          sub(/^\/\/ (Esquema del escenario|Escenario):[ \t]*/, "", l)
          sub(/[ \t]+$/, "", l)
          # "Los datos guardados estan danados (el archivo)": el agregado entre parentesis
          # dice que parte del escenario cubre este test; el escenario es lo de antes
          sub(/ \([^()]*\)$/, "", l)
          escenarios[++n_esc] = l
        } else if (cubre == "" && l ~ /^\/\/ Cubre:/) {
          sub(/^\/\/ Cubre:[ \t]*/, "", l); cubre = l
        }
        if (l ~ /Sin escenario|[Nn]o tiene escenario|Caso l(i|í)mite/) declara = 1
      }
      if (n_esc == 0) escenarios[++n_esc] = ""
      for (i = 1; i <= n_esc; i++)
        printf "%s\t%s\t%s\t%s\t%s\t%s\n", nombre, archivo, FNR, escenarios[i], cubre, declara
      bloque = ""; next
    }
    { bloque = "" }
  ' "$archivo" >> "$TMP/tests.tsv"
done < "$TMP/archivos-test.txt"

# --- 3. Historial: commit que agrego cada test y el GREEN que lo sigue --------------------

# Primer commit donde aparece "func TestX(" en un archivo de test
git log --reverse --no-color --format='@@%H' -p -- '*_test.go' | awk '
  /^@@[0-9a-f]+$/ { commit = substr($0, 3); next }
  /^\+func Test[A-Za-z0-9_]*\(/ {
    nombre = $2; sub(/\(.*/, "", nombre)
    if (!(nombre in visto)) { visto[nombre] = 1; print nombre "\t" commit }
  }
' > "$TMP/primer-commit.tsv"

# Hijo de cada commit en su misma rama: el commit cuyo primer padre es ese. Si tiene varios,
# queda el mas viejo, que es el siguiente de la rama. Tambien se anota si el hijo es un merge.
git rev-list --parents HEAD | awk '{ print $2 "\t" $1 "\t" (NF - 1) }' \
  > "$TMP/hijos.tsv"

# Archivos de codigo que cambia cada commit, separados de los tests. Es codigo de produccion
# todo .go que no es test y, en la web, las plantillas (.html) y el CSS (.css) de internal/web
git log --no-color --format='@@%H' --name-only HEAD | awk '
  function cerrar() { if (commit != "") print commit "\t" codigo "\t" tests }
  /^@@[0-9a-f]+$/ { cerrar(); commit = substr($0, 3); codigo = ""; tests = 0; next }
  /_test\.go$/ { tests = 1; next }
  /\.go$/ || /^internal\/web\/.*\.(html|css)$/ { codigo = codigo (codigo == "" ? "" : " ") $0 }
  END { cerrar() }
' > "$TMP/archivos-commit.tsv"

# Commits que dicen "sin RED" o "(paso) en verde al escribirlo": el test paso con el codigo
# que ya habia, y no hay GREEN. Las dos frases dicen lo mismo; la segunda es la que usan la #8
# y la #9. git busca linea por linea, y en algunos mensajes "en" queda al final de la linea
# anterior: por eso se busca "verde al escribir". Las busquedas no ignoran mayusculas:
# "sin red" tambien esta en "sin redondear". "(pasa) sin codigo nuevo" es la tercera forma de
# decir lo mismo, con o sin tilde.
git log --no-color --format='%H' --grep='sin RED' --grep='verde al escribir' \
  --grep='sin codigo nuevo' --grep='sin código nuevo' HEAD > "$TMP/sin-red.txt"

# Commits "test:": si tocan codigo, es lo minimo para que el test compile (una constante, una
# firma), y la implementacion viene en el commit siguiente
git log --no-color --format='%H %s' HEAD | awk '$2 == "test:" { print $1 }' > "$TMP/commits-test.txt"

# Para cada test: commit RED, commit GREEN y archivos de codigo del GREEN
awk -F'\t' -v dir="$TMP" '
  FILENAME == dir "/hijos.tsv" { if ($1 != "") { hijo[$1] = $2; padres[$2] = $3 }; next }
  FILENAME == dir "/archivos-commit.tsv" { codigo[$1] = $2; toca_tests[$1] = $3; next }
  FILENAME == dir "/sin-red.txt" { sin_red[$1] = 1; next }
  FILENAME == dir "/commits-test.txt" { es_test[$1] = 1; next }
  {
    nombre = $1; red = $2; green = ""; archivos = ""; nota = ""
    if (red in sin_red) {
      nota = "sin-red"
    } else if (codigo[red] != "" && !(red in es_test)) {
      nota = "mismo-commit"; archivos = codigo[red]
    } else {
      actual = red
      for (paso = 0; paso < 8; paso++) {
        h = hijo[actual]
        if (h == "" || padres[h] > 1) break
        if (codigo[h] != "") { green = h; archivos = codigo[h]; break }
        if (toca_tests[h] == 1) break
        actual = h
      }
      if (green == "" && codigo[red] != "") { nota = "mismo-commit"; archivos = codigo[red] }
      else if (green == "") nota = "sin-green"
      else if (codigo[red] != "") nota = "andamiaje"
    }
    print nombre "\t" red "\t" green "\t" archivos "\t" nota
  }
' "$TMP/hijos.tsv" "$TMP/archivos-commit.tsv" "$TMP/sin-red.txt" "$TMP/commits-test.txt" \
  "$TMP/primer-commit.tsv" \
  > "$TMP/green.tsv"

# --- 4. A que historia va cada test -------------------------------------------------------
#
# Con escenario: a la historia cuya spec tiene ese escenario. Si el titulo esta en mas de una
# spec, a las que listan el archivo del test en su tabla de trazabilidad.
# Sin escenario: a la historia que nombra en "Cubre:" (#N), o si no, a la de numero mas bajo
# que lista el archivo del test. Si ninguna lo lista, queda sin historia.

awk -F'\t' -v dir="$TMP" '
  # Un patron de la tabla ("internal/web/*_test.go") pasado a expresion regular
  function patron(p,    r) {
    r = p; gsub(/\./, "\\.", r); gsub(/\*/, "[^/]*", r); return "^" r "$"
  }
  function lista_el_archivo(h, archivo,    i, n, ps) {
    n = split(tests_de[h], ps, " ")
    for (i = 1; i <= n; i++) if (archivo ~ patron(ps[i])) return 1
    return 0
  }
  FILENAME == dir "/specs.tsv" { historias[++nh] = $1; tests_de[$1] = $4; next }
  FILENAME == dir "/escenarios.tsv" { en_spec[$2] = en_spec[$2] " " $1; next }
  {
    nombre = $1; archivo = $2; escenario = $4; cubre = $5
    destino = ""; nota = ""
    if (escenario != "" && en_spec[escenario] != "") {
      n = split(en_spec[escenario], cand, " ")
      if (n == 1) destino = cand[1]
      else {
        for (i = 1; i <= n; i++) if (lista_el_archivo(cand[i], archivo)) destino = destino " " cand[i]
        if (destino == "") { destino = en_spec[escenario]; nota = "titulo-repetido" }
      }
    } else {
      if (escenario != "") nota = "escenario-fuera-de-spec"
      if (match(cubre, /#[0-9]+/)) {
        h = substr(cubre, RSTART + 1, RLENGTH - 1)
        if (tests_de[h] != "" || h in tests_de) destino = h
      }
      if (destino == "") {
        mejor = ""
        for (i = 1; i <= nh; i++)
          if (lista_el_archivo(historias[i], archivo) && (mejor == "" || historias[i] + 0 < mejor + 0))
            mejor = historias[i]
        destino = mejor
      }
    }
    if (destino == "") destino = "-"
    n = split(destino, ds, " ")
    for (i = 1; i <= n; i++) print ds[i] "\t" $0 "\t" nota
  }
' "$TMP/specs.tsv" "$TMP/escenarios.tsv" "$TMP/tests.tsv" > "$TMP/asignados.tsv"

# --- 5. El documento ----------------------------------------------------------------------

awk -F'\t' -v dir="$TMP" '
  function link_test(archivo, linea,    base) {
    base = archivo; sub(/.*\//, "", base)
    return "[" base ":" linea "](../" archivo "#L" linea ")"
  }
  function links_codigo(lista,    r, i, n, fs, base) {
    r = ""; n = split(lista, fs, " ")
    for (i = 1; i <= n; i++) {
      base = fs[i]; sub(/.*\//, "", base)
      r = r (r == "" ? "" : ", ") "[" base "](../" fs[i] ")"
    }
    return r
  }
  function commits(nombre,    r, g) {
    if (red[nombre] == "") return "—"
    r = "`" substr(red[nombre], 1, 7) "`"; g = "`" substr(green[nombre], 1, 7) "`"
    if (nota_green[nombre] == "mismo-commit") return r " ⚠️ test y código en el mismo commit"
    if (nota_green[nombre] == "sin-red") return r " (sin RED: ya pasaba)"
    if (green[nombre] == "") return r " → ⚠️ sin GREEN"
    if (nota_green[nombre] == "andamiaje") return r " → " g " (el test trae lo justo para compilar)"
    return r " → " g
  }
  function codigo_de(nombre) {
    if (archivos[nombre] != "") return links_codigo(archivos[nombre])
    if (nota_green[nombre] == "sin-red") return "el que ya estaba"
    if (red[nombre] == "") return "⚠️ no se encontró el commit del test"
    return "⚠️ ningún commit de código después del test"
  }

  FILENAME == dir "/specs.tsv" {
    nh++; orden[nh] = $1; spec[$1] = $2; titulo[$1] = $3; codigo_tabla[$1] = $5; next
  }
  FILENAME == dir "/escenarios.tsv" { ne[$1]++; esc[$1, ne[$1]] = $2; esc_linea[$1, ne[$1]] = $3; next }
  FILENAME == dir "/criterios.tsv" { nc[$1]++; crit[$1, nc[$1]] = $2; next }
  FILENAME == dir "/green.tsv" {
    red[$1] = $2; green[$1] = $3; archivos[$1] = $4; nota_green[$1] = $5; next
  }
  {
    # asignados.tsv: historia, nombre, archivo, linea, escenario, cubre, declara, nota
    h = $1; nt[h]++
    t_nombre[h, nt[h]] = $2; t_archivo[h, nt[h]] = $3; t_linea[h, nt[h]] = $4
    t_esc[h, nt[h]] = $5; t_cubre[h, nt[h]] = $6; t_declara[h, nt[h]] = $7; t_nota[h, nt[h]] = $8
    if ($5 != "" && $8 != "escenario-fuera-de-spec") {
      con_test[h, $5] = con_test[h, $5] " " nt[h]; valido[h, $2] = 1
    }
    if ($8 == "escenario-fuera-de-spec") fuera[h, $2] = fuera[h, $2] (fuera[h, $2] == "" ? "" : ", ") "«" $5 "»"
  }

  # Cuantos tests distintos de la historia h no cubren ningun escenario de su spec
  function contar_sin_escenario(h,    j, n, vistos) {
    n = 0; split("", vistos)
    for (j = 1; j <= nt[h]; j++)
      if (!valido[h, t_nombre[h, j]] && !(t_nombre[h, j] in vistos)) { vistos[t_nombre[h, j]] = 1; n++ }
    return n
  }

  # Tabla de los tests de h que no cubren ningun escenario de su spec
  function tabla_sin_escenario(h,    j, nombre, marca, vistos) {
    split("", vistos)
    print "| Test | Cubre | Código | RED → GREEN |"
    print "|---|---|---|---|"
    for (j = 1; j <= nt[h]; j++) {
      nombre = t_nombre[h, j]
      if (valido[h, nombre] || (nombre in vistos)) continue
      vistos[nombre] = 1; marca = ""
      if (fuera[h, nombre] != "") marca = " ⚠️ cita " fuera[h, nombre] ", que no es un escenario de la spec"
      else if (t_declara[h, j] != 1) marca = " ⚠️ no dice que no tiene escenario"
      printf "| `%s` %s%s | %s | %s | %s |\n", nombre, link_test(t_archivo[h, j], t_linea[h, j]), \
        marca, (t_cubre[h, j] == "" ? "⚠️ sin `Cubre:`" : t_cubre[h, j]), codigo_de(nombre), commits(nombre)
    }
    print ""
  }

  END {
    print "# Trazabilidad"
    print ""
    print "_Generado por `scripts/trazabilidad.sh`. No editar a mano: se regenera con el script cada vez"
    print "que entra a `dev` un pull request que toca tests o specs._"
    print ""
    print "En la defensa eligen una funcionalidad y piden el recorrido completo: **historia → spec →"
    print "criterios de aceptación → escenario BDD → test → código**. Acá está ese recorrido para cada"
    print "historia que ya tiene tests en el repositorio."
    print ""
    print "- **Escenarios y criterios**: salen de la Fase 1 de cada spec."
    print "- **Test de cada escenario**: el comentario `// Escenario:` arriba del test, como pide"
    print "  `AGENTS.md`. Si el título del escenario está en más de una spec, el test va a la historia"
    print "  cuya tabla de trazabilidad lista su archivo."
    print "- **Código de cada test**: sale del historial de TDD. *RED* es el commit que agregó el test y"
    print "  *GREEN* el commit siguiente de su rama que toca código; los archivos son los que cambió el"
    print "  GREEN."
    print "- Las specs no dicen qué escenario prueba cada criterio, así que los criterios se listan"
    print "  junto a la historia, sin cruzarlos."
    print ""
    print "Los huecos se marcan con ⚠️ y no se tapan."
    print ""

    # Historias que entran: las que tienen al menos un test asignado
    n_inc = 0
    for (i = 1; i <= nh; i++) if (nt[orden[i]] > 0) inc[++n_inc] = orden[i]

    print "## Resumen"
    print ""
    print "| Historia | Escenarios | Con test | ⚠️ Sin test | Tests sin escenario |"
    print "|---|---|---|---|---|"
    for (k = 1; k <= n_inc; k++) {
      h = inc[k]; si = 0; no = 0; sin_esc = contar_sin_escenario(h)
      for (j = 1; j <= ne[h]; j++) if (con_test[h, esc[h, j]] != "") si++; else no++
      printf "| [#%s %s](#%s) | %d | %d | %d | %d |\n", h, titulo[h], ancla(h), ne[h] + 0, si, no, sin_esc
    }
    if (nt["-"] > 0) printf "| [Tests de ninguna historia](#tests-de-ninguna-historia) | — | — | — | %d |\n", contar_sin_escenario("-")
    print ""

    # Las marcas del historial de TDD, contadas y explicadas
    m_verde = 0; m_sin_red = 0; m_juntos = 0; m_sin_green = 0; m_andamiaje = 0
    for (nombre in red) {
      if (nota_green[nombre] == "mismo-commit") m_juntos++
      else if (nota_green[nombre] == "andamiaje") m_andamiaje++
      else if (nota_green[nombre] == "sin-red") m_sin_red++
      else if (green[nombre] == "") m_sin_green++
      else m_verde++
    }
    print "## Qué dice el historial de TDD"
    print ""
    print "Cada test tiene en la columna *RED → GREEN* lo que muestra el historial de commits. Hay"
    print "cinco casos posibles:"
    print ""
    print "| Marca | Qué quiere decir | Tests |"
    print "|---|---|---|"
    printf "| `abc1234` → `def5678` | El ciclo completo: el test entró en un commit y el código que lo hizo pasar, en el siguiente. | %d |\n", m_verde
    printf "| `abc1234` → `def5678` (el test trae lo justo para compilar) | También es el ciclo completo. El commit del test agrega lo mínimo para que compile (una constante, una firma) y falla; la implementación viene en el commit siguiente. | %d |\n", m_andamiaje
    printf "| `abc1234` (sin RED: ya pasaba) | El test pasó con el código que ya estaba, y quien lo escribió lo dejó dicho en el mensaje del commit. | %d |\n", m_sin_red
    printf "| `abc1234` → ⚠️ sin GREEN | Después del test no hay un commit de código en su rama: pasó con lo que ya estaba. Puede ser que una tarea anterior ya cubriera el caso o que el código se haya escrito antes que el test; el historial no alcanza para saber cuál de las dos. | %d |\n", m_sin_green
    printf "| `abc1234` ⚠️ test y código en el mismo commit | El test y el código entraron juntos, así que el historial no muestra la fase en rojo. Es un desvío del TDD. | %d |\n", m_juntos
    print ""

    for (k = 1; k <= n_inc; k++) {
      h = inc[k]
      printf "## #%s %s\n\n", h, titulo[h]
      printf "- **Historia:** issue #%s\n", h
      printf "- **Spec:** [%s](%s)\n", spec[h], substr(spec[h], 6)
      if (codigo_tabla[h] != "") { c = codigo_tabla[h]; gsub(/ /, "`, `", c); printf "- **Código según la spec:** `%s`\n", c }
      print ""
      if (nc[h] > 0) {
        printf "<details><summary>Criterios de aceptación (%d)</summary>\n\n", nc[h]
        for (j = 1; j <= nc[h]; j++) print "- " crit[h, j]
        print "\n</details>\n"
      }
      print "| Escenario BDD | Test | Código | RED → GREEN |"
      print "|---|---|---|---|"
      for (j = 1; j <= ne[h]; j++) {
        e = esc[h, j]
        # ?plain=1 hace que GitHub abra el .md como texto y salte a la linea
        link_esc = "[" e "](" substr(spec[h], 6) "?plain=1#L" esc_linea[h, j] ")"
        if (con_test[h, e] == "") { printf "| ⚠️ %s | **sin test** | | |\n", link_esc; continue }
        n = split(con_test[h, e], idx, " ")
        for (m = 1; m <= n; m++) {
          x = idx[m]; nombre = t_nombre[h, x]
          extra = (t_nota[h, x] == "titulo-repetido") ? " ⚠️ título repetido en otra spec" : ""
          printf "| %s%s | `%s` %s | %s | %s |\n", (m == 1 ? link_esc : "″"), extra, nombre, \
            link_test(t_archivo[h, x], t_linea[h, x]), codigo_de(nombre), commits(nombre)
        }
      }
      print ""
      # Tests que cubren escenarios de la spec y ademas citan uno que no esta
      split("", vistos)
      for (j = 1; j <= nt[h]; j++) {
        nombre = t_nombre[h, j]
        if (valido[h, nombre] && fuera[h, nombre] != "" && !(nombre in vistos)) {
          vistos[nombre] = 1
          printf "- ⚠️ `%s` también cita %s, que no es un escenario de la spec.\n", nombre, fuera[h, nombre]
          otra = 1
        }
      }
      if (otra) { print ""; otra = 0 }
      hay = contar_sin_escenario(h)
      if (hay > 0) {
        printf "**Tests que no citan un escenario de la Fase 1 (%d).**\n\n", hay
        tabla_sin_escenario(h)
      }
    }

    if (nt["-"] > 0) {
      print "## Tests de ninguna historia"
      print ""
      print "Ninguna tabla de trazabilidad lista su archivo, y el test no cita un escenario ni nombra"
      print "una historia en `Cubre:`."
      print ""
      tabla_sin_escenario("-")
    }
  }

  # Ancla de GitHub para el titulo "## #N Titulo"
  function ancla(h,    a) {
    a = tolower("#" h " " titulo[h]); gsub(/[^a-z0-9áéíóúñü -]/, "", a); gsub(/ /, "-", a); return a
  }
' "$TMP/specs.tsv" "$TMP/escenarios.tsv" "$TMP/criterios.tsv" "$TMP/green.tsv" "$TMP/asignados.tsv" \
  > "$TMP/trazabilidad.md"

if [ "${1:-}" = "--stdout" ]; then
  cat "$TMP/trazabilidad.md"
else
  cp "$TMP/trazabilidad.md" "$SALIDA"
  echo "Escrito $SALIDA"
fi
