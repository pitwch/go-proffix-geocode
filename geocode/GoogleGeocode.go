package geocode

import (
	"github.com/kelvins/geocoder"
	"github.com/pitwch/go-proffix-geocode/adresses"
	"github.com/spf13/cast"
	"log"
)

// Geocode PROFFIX Adressen with Google

func GeocodeAdressesGoogle(adresses []Adresses.Adresses) (g []Adresses.Geocoded) {

	// Set the API Key from Config

	geocoder.ApiKey = settings.GoogleAPIKey

	for _, y := range adresses {
		address := geocoder.Address{
			Street:  y.Street.String,
			City:    y.City.String,
			Country: y.Country.String,
		}

		location, err := geocoder.Geocoding(address)

		if err != nil {
			log.Println("Could not get the location for AdressNr "+cast.ToString(y.Adressnr), err)
		} else {
			g = append(g, Adresses.Geocoded{y.Adressnr, cast.ToString(location.Longitude), cast.ToString(location.Latitude)})

		}
	}

	return g
}
