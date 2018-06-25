package configuration

import (
	"github.com/pitwch/go-proffix-geocode/shared"
	"os"
	"fmt"
	"encoding/json"
	"github.com/pitwch/go-proffix-geocode/geocode"
	"flag"
)

type Configuration struct {
	Database database.Database
	Settings geocode.Settings
}

// Read Config from config.json
func Config() {

	configuration := Configuration{}


	c := new(Configuration)
	flag.StringVar(&c.Database.Database, "Database", "", "PROFFIX Database Name")
	flag.IntVar(&c.Database.Port ,"Port",1433, "Database Port")
	flag.StringVar(&c.Database.Username ,"Username","", "Database Username")
	flag.StringVar(&c.Database.Password ,"Password","", "Database Password")
	flag.StringVar(&c.Database.Hostname ,"Hostname","", "Database Hostname")
	flag.StringVar(&c.Database.Parameter ,"Parameter","?connection+timeout=30", "Database Parameter")


	flag.StringVar(&c.Settings.UseGeocoder ,"UseGeocoder","Google", "Geocoding Provider")
	flag.StringVar(&c.Settings.GoogleAPIKey ,"GoogleAPIKey","", "Google API-Key")
	flag.StringVar(&c.Settings.MapQuestAPIKey ,"MapQuestAPIKey","", "Mapquest API-Key")


	// Read config.json. Config has to be next to executable.
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Printf("Error reading Config File: %v", err)
		os.Exit(1)
	}
	database.Connect(configuration.Database)
	geocode.SetSettings(configuration.Settings)






}
