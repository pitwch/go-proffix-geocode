package geocode

import (
	"go-proffix-geocode/adresses"
	"github.com/spf13/cast"
	"github.com/kelvins/geocoder"
	"fmt"
)



// Geocode PROFFIX Adressen with Google

func GeocodeAdressesGoogle(adresses []Adresses.Adresses) (g []Adresses.Geocoded) {

	for _, y := range adresses {

	address := geocoder.Address{
		Street:  y.Street.String,
		City:    y.City.String,
		Country: y.Country.String,
	}

		location, err := geocoder.Geocoding(address)

		if err != nil {
			fmt.Println("Could not get the location: ", err)
		} else {
			g = append(g, Adresses.Geocoded{y.Adressnr, cast.ToString(location.Longitude), cast.ToString(location.Latitude)})

		}
	}

	return g
}
