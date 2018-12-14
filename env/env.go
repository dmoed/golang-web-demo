package env

import (
	"apptastic/dashboard/view"
	"database/sql"
	"log"
	"time"
)

type Env struct {
	DB               *sql.DB
	Logger           *log.Logger
	View             *view.View
	TimezoneLocation *time.Location
}
