package utils

import (
	"database/sql"
)

type AppState struct {
	Env *Env
	Db  *sql.DB
}
