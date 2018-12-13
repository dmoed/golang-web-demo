package main

import (
	"apptastic/dashboard/config"
	"apptastic/dashboard/database"
	"apptastic/dashboard/env"
	"apptastic/dashboard/server"
	"apptastic/dashboard/session"
	"apptastic/dashboard/view"
	"log"
	"os"
)

var environment string
var port string

func main() {

	//Load Env
	environment = getenv("ENVIRONMENT", "dev")
	port = getenv("PORT", "5000")

	//load config
	config := config.Load(environment)

	//Configure session cookie store
	session.Configure(config.Session)

	//init env struct
	env := &env.Env{
		DB:     database.Connect(config.Database), //Connect to database
		Logger: log.New(os.Stdout, "", 0),         //Create new Logger
		View:   view.New(config.View),
	}

	//start server
	s := server.NewServer(config.Server, env)
	s.Run(port)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
