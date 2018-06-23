package main

import (
	"log"
	"go-proffix-geocode/configuration"
	"go-proffix-geocode/geocode"
	"go-proffix-geocode/adresses"
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



