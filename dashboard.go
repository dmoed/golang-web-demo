package main

import (
	"apptastic/dashboard/config"
	"apptastic/dashboard/database"
	"apptastic/dashboard/env"
	"apptastic/dashboard/server"
	"apptastic/dashboard/session"
	"log"
	"os"
)

var (
	environment = "dev"
)

func main() {

	environment = getenv("ENVIRONMENT", "dev")

	//load config
	config := config.Load(environment)

	//Configure session cookie store
	session.Configure(config.Session)

	//init env
	env := &env.Env{
		DB:     database.Connect(config.Database), //Connect to database
		Logger: log.New(os.Stdout, "", 0),         //Create new Logger
	}

	//start server
	s := server.NewServer(config.Server, env)
	s.Run()

}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
