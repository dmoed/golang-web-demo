package config

import (
	"apptastic/dashboard/database"
	"apptastic/dashboard/server"
	"apptastic/dashboard/session"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	filename  = "config"
	extension = ".json"
)

type Configuration struct {
	Database database.DatabaseConfig `json:"Database"`
	Session  session.SessionConfig   `json:"Session"`
	Server   server.ServerConfig     `json:"Server"`
}

//LoadConfiguration from config json
func Load(env string) *Configuration {

	var config = &Configuration{}
	var file = strings.ToLower(fmt.Sprintf("%v_%v%v", filename, env, extension))

	fmt.Println(fmt.Sprintf("Loading config file %v", file))

	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}
