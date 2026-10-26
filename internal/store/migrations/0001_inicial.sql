-- Esquema inicial de la base de la web (historia #177, D2 y D6). Una tabla por entidad, y
-- tablas aparte para lo que una entidad tiene varias veces. Las columnas siguen el formato
-- del archivo de la #77 (internal/domain/registry_json.go): las fechas son dias "AAAA-MM-DD",
-- el rol, la prioridad y los estados van con su nombre, y "position" guarda el orden en que
-- estaban, que el dominio conserva. Las tablas son STRICT para que un tipo equivocado sea un
-- error y no un dato convertido en silencio.

-- El proximo numero de proyecto (D2): una sola fila.
CREATE TABLE registry (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	next_id INTEGER NOT NULL
) STRICT;

-- Los contadores de cada proyecto se guardan tal como los trae la conversion (D2): nunca se
-- deducen del numero mas alto guardado.
CREATE TABLE projects (
	id INTEGER PRIMARY KEY,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	start_day TEXT NOT NULL,
	end_day TEXT NOT NULL,
	next_story_id INTEGER NOT NULL,
	last_sprint_number INTEGER NOT NULL,
	last_defect_id INTEGER NOT NULL
) STRICT;

CREATE TABLE members (
	project_id INTEGER NOT NULL,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	role TEXT NOT NULL,
	PRIMARY KEY (project_id, position)
) STRICT;

CREATE TABLE stories (
	project_id INTEGER NOT NULL,
	id INTEGER NOT NULL,
	position INTEGER NOT NULL,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	priority TEXT NOT NULL,
	state TEXT NOT NULL,
	story_points INTEGER NOT NULL,
	estimated_hours REAL NOT NULL,
	PRIMARY KEY (project_id, id)
) STRICT;

CREATE TABLE acceptance_criteria (
	project_id INTEGER NOT NULL,
	story_id INTEGER NOT NULL,
	position INTEGER NOT NULL,
	text TEXT NOT NULL,
	PRIMARY KEY (project_id, story_id, position)
) STRICT;

CREATE TABLE efforts (
	project_id INTEGER NOT NULL,
	position INTEGER NOT NULL,
	story_id INTEGER NOT NULL,
	member_name TEXT NOT NULL,
	activity TEXT NOT NULL,
	hours REAL NOT NULL,
	day TEXT NOT NULL,
	PRIMARY KEY (project_id, position)
) STRICT;

-- closed_on vacio es un sprint sin cerrar.
CREATE TABLE sprints (
	project_id INTEGER NOT NULL,
	number INTEGER NOT NULL,
	position INTEGER NOT NULL,
	goal TEXT NOT NULL,
	start_day TEXT NOT NULL,
	end_day TEXT NOT NULL,
	state TEXT NOT NULL,
	closed_on TEXT NOT NULL,
	PRIMARY KEY (project_id, number)
) STRICT;

CREATE TABLE sprint_stories (
	project_id INTEGER NOT NULL,
	sprint_number INTEGER NOT NULL,
	position INTEGER NOT NULL,
	story_id INTEGER NOT NULL,
	PRIMARY KEY (project_id, sprint_number, position)
) STRICT;

CREATE TABLE sprint_completed_stories (
	project_id INTEGER NOT NULL,
	sprint_number INTEGER NOT NULL,
	position INTEGER NOT NULL,
	story_id INTEGER NOT NULL,
	PRIMARY KEY (project_id, sprint_number, position)
) STRICT;

CREATE TABLE sprint_points_at_close (
	project_id INTEGER NOT NULL,
	sprint_number INTEGER NOT NULL,
	story_id INTEGER NOT NULL,
	points INTEGER NOT NULL,
	PRIMARY KEY (project_id, sprint_number, story_id)
) STRICT;

-- resolved_in en 0 es un defecto sin resolver, como en el archivo.
CREATE TABLE defects (
	project_id INTEGER NOT NULL,
	id INTEGER NOT NULL,
	position INTEGER NOT NULL,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	severity TEXT NOT NULL,
	state TEXT NOT NULL,
	story_id INTEGER NOT NULL,
	detected_in INTEGER NOT NULL,
	resolved_in INTEGER NOT NULL,
	PRIMARY KEY (project_id, id)
) STRICT;

-- agreed en 0 es una ronda sin acuerdo (#73).
CREATE TABLE voting_rounds (
	project_id INTEGER NOT NULL,
	position INTEGER NOT NULL,
	story_id INTEGER NOT NULL,
	number INTEGER NOT NULL,
	revealed INTEGER NOT NULL,
	agreed INTEGER NOT NULL,
	PRIMARY KEY (project_id, position)
) STRICT;

CREATE TABLE voting_round_participants (
	project_id INTEGER NOT NULL,
	round_position INTEGER NOT NULL,
	position INTEGER NOT NULL,
	name TEXT NOT NULL,
	PRIMARY KEY (project_id, round_position, position)
) STRICT;

-- Los votos de una ronda sin revelar tambien se guardan: el secreto vale dentro del sistema,
-- como en el archivo (punto 8 de la puerta de la #13).
CREATE TABLE votes (
	project_id INTEGER NOT NULL,
	round_position INTEGER NOT NULL,
	position INTEGER NOT NULL,
	member_name TEXT NOT NULL,
	points INTEGER NOT NULL,
	PRIMARY KEY (project_id, round_position, position)
) STRICT;
