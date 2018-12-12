package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Mysql MySQLConfig `json:"Mysql"`
}

type MySQLConfig struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	Hostname string `json:"Host"`
	Port     int    `json:"Port"`
}

//Connect and return the database
func Connect(d DatabaseConfig) *sql.DB {

	DB, err := sql.Open("mysql", DNS(d.Mysql))

	if err != nil {
		panic(err.Error())
	}

	if err := DB.Ping(); err != nil {
		panic(err.Error())
	}

	return DB
}

//DNS returns the Data Source Name
func DNS(m MySQLConfig) string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Hostname + ":" + fmt.Sprintf("%d", m.Port) + ")/" + m.Name
}
