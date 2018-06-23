package main

import (
	"log"
	"go-proffix-geocode/configuration"
	"go-proffix-geocode/geocode"
	"go-proffix-geocode/adresses"
	"github.com/spf13/cast"
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

	x := geocode.GeocodeAdressesGoogle(results)

	for _,y := range x {
		adressnr,err := Adresses.UpdateEmptyAdresses(y)

		// If Error -> Log
		if err != nil {
			log.Println(err)
		}

		log.Println("Successfully geocoded PROFFIX AdressNr "+cast.ToString(adressnr))
	}

}
