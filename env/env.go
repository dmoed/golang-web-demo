package env

import (
	"apptastic/dashboard/view"
	"database/sql"
	"log"
)

type Env struct {
	DB     *sql.DB
	Logger *log.Logger
	View   *view.View
}
