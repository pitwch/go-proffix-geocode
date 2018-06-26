package geocode

import (
	"testing"
	"github.com/pitwch/go-proffix-geocode/adresses"
	"database/sql"
)

func TestGeocodeAdressesGoogle(t *testing.T) {
	testaddresses := []Adresses.Adressstruct{
		Adresses.Adressstruct{
			1,
			sql.NullString{"Murgtalstrasse 20", true},
			sql.NullString{"9542", true},
			sql.NullString{"Münchwilen", true},
			sql.NullString{"CH", true},
			sql.NullString{"", true},
			sql.NullString{"", true},
		},
	}
	GeocodeAdressesGoogle(testaddresses)
}