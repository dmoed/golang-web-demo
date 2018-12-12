package env

import (
	"database/sql"
	"log"
)

type Env struct {
	DB     *sql.DB
	Logger *log.Logger
}
