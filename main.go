//go:generate goversioninfo -icon=icon.ico -manifest=go-proffix-geocode.exe.manifest

package main

import (
	"log"
	"github.com/pitwch/go-proffix-geocode/configuration"
	"github.com/pitwch/go-proffix-geocode/geocode"
	"github.com/pitwch/go-proffix-geocode/adresses"
)



func main() {

	// Read Config
	configuration.Config()

	// Read Adresses from PROFFIX with non-set Long / Lat
	results, err := Adresses.GetEmptyAdresses()

	if len(results) < 1 {
		log.Println("All adresses are already geocoded")
	} else if
	err != nil {
		log.Println(err)


	}

	// Batch Geocode Adresses from PROFFIX


	geocode.GeocodeAdresses(results)

}



