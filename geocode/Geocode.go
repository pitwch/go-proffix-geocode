package geocode

import (
	"github.com/pitwch/go-proffix-geocode/adresses"
	"github.com/spf13/cast"
	"log"
)

type Settings struct {
	UseGeocoder    string
	MapQuestAPIKey string
	GoogleAPIKey   string
}

var settings Settings

func SetSettings(s Settings) {
	settings = s
}
func GeocodeAdresses(adresses []Adresses.Adressstruct) {

	var r []Adresses.Geocoded
	var err error

	if settings.UseGeocoder == "Mapquest" || settings.UseGeocoder == "mapquest" {
		r = GeocodeAdressesMapquest(adresses)
	} else {
		r = GeocodeAdressesGoogle(adresses)

		if err != nil {
			log.Println(err)
		}

	}

	for _, y := range r {
		adressnr, err := Adresses.UpdateEmptyAdresses(y)

		// If Error -> Log
		if err != nil {
			log.Println(err)
		}

		log.Println("Successfully geocoded PROFFIX AdressNr " + cast.ToString(adressnr))
	}
}
