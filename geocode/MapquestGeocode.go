package geocode

import (
	"github.com/pitwch/go-proffix-geocode/adresses"
	"github.com/jasonwinn/geocoder"
	"log"
	"github.com/spf13/cast"
	"strings"
)

// Replace German Umlauts as Mapquest API has problems with them
func replaceUmlauts(s string) (r string) {
	var replacer = strings.NewReplacer("ä", "ae", "ö", "oe", "ü", "ue")
	r = replacer.Replace(s)
	return r
}

// Geocode PROFFIX Adressen with Mapquest Batch
func GeocodeAdressesMapquest(adresses []Adresses.Adressstruct) (g []Adresses.Geocoded) {

	// Set the API Key from Config
	geocoder.SetAPIKey(settings.MapQuestAPIKey)

	var Temp []string
	for _, y := range adresses {
		temp2 := replaceUmlauts(y.Street.String + " " + y.PLZ.String + " " + y.City.String)
		log.Println(temp2)
		Temp = append(Temp, temp2)
	}

	x, err := geocoder.BatchGeocode(Temp)
	if len(x) < 1 {
		log.Println("All Adresses are already geocoded.")
	} else if
	(err != nil) {
		log.Println(err)
	}

	for c, r := range adresses {
		lat := x[c].Lat
		long := x[c].Lng

		g = append(g, Adresses.Geocoded{r.Adressnr, cast.ToString(long), cast.ToString(lat)})
	}

	return g
}
