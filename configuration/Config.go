package configuration

import (
	"go-proffix-geocode/shared"
	"os"
	"fmt"
	"encoding/json"
)

type Configuration struct {
	Database database.Database
	Google   MapQuest
}

type MapQuest struct {
	APIKey string
}


// Read Config from config.json
func Config() {
	//Read config.json. Config has to be next to executable.
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Printf("Error reading Config File: %v", err)
		os.Exit(1)
	}

	database.Connect(configuration.Database)
}
