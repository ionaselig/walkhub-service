package walkhub

import (
	"errors"

	"github.com/tamasd/ab"
)

// AUTOGENERATED DO NOT EDIT

func NewLog() *Log {
	e := &Log{}

	// HOOK: newLog()

	return e
}

func EmptyLog() *Log {
	return &Log{}
}

var _ ab.Validator = &Log{}

func (e *Log) Validate() error {
	var err error

	// HOOK: validateLog()

	return err
}

func (e *Log) GetID() string {
	return e.UUID
}

var LogNotFoundError = errors.New("log not found")

const logFields = "l.uuid, l.type, l.message, l.created"

func selectLogFromQuery(db ab.DB, query string, args ...interface{}) ([]*Log, error) {
	// HOOK: beforeLogSelect()

	entities := []*Log{}

	rows, err := db.Query(query, args...)

	if err != nil {
		return entities, err
	}

	for rows.Next() {
		e := EmptyLog()

		if err = rows.Scan(&e.UUID, &e.Type, &e.Message, &e.Created); err != nil {
			return []*Log{}, err
		}

		entities = append(entities, e)
	}

	// HOOK: afterLogSelect()

	return entities, err
}

func selectSingleLogFromQuery(db ab.DB, query string, args ...interface{}) (*Log, error) {
	entities, err := selectLogFromQuery(db, query, args...)
	if err != nil {
		return nil, err
	}

	if len(entities) > 0 {
		return entities[0], nil
	}

	return nil, nil
}

func (e *Log) Insert(db ab.DB) error {
	// HOOK: beforeLogInsert()

	err := db.QueryRow("INSERT INTO \"log\"(type, message, created) VALUES($1, $2, $3) RETURNING uuid", e.Type, e.Message, e.Created).Scan(&e.UUID)

	// HOOK: afterLogInsert()

	return err
}

func LoadLog(db ab.DB, UUID string) (*Log, error) {
	// HOOK: beforeLogLoad()

	e, err := selectSingleLogFromQuery(db, "SELECT "+logFields+" FROM \"log\" l WHERE l.uuid = $1", UUID)

	// HOOK: afterLogLoad()

	return e, err
}

func LoadAllLog(db ab.DB, start, limit int) ([]*Log, error) {
	// HOOK: beforeLogLoadAll()

	entities, err := selectLogFromQuery(db, "SELECT "+logFields+" FROM \"log\" l ORDER BY UUID DESC LIMIT $1 OFFSET $2", limit, start)

	// HOOK: afterLogLoadAll()

	return entities, err
}

func (s *LogService) SchemaInstalled(db ab.DB) bool {
	found := ab.TableExists(db, "log")

	// HOOK: afterLogSchemaInstalled()

	return found
}

func (s *LogService) SchemaSQL() string {
	sql := "CREATE TABLE \"log\" (\n" +
		"\t\"uuid\" uuid DEFAULT uuid_generate_v4() NOT NULL,\n" +
		"\t\"type\" character varying NOT NULL,\n" +
		"\t\"message\" character varying NOT NULL,\n" +
		"\t\"created\" timestamp with time zone DEFAULT now() NOT NULL,\n" +
		"\tCONSTRAINT log_pkey PRIMARY KEY (uuid)\n);\n"

	// HOOK: afterLogSchemaSQL()

	return sql
}