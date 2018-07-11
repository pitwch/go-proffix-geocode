package test

import (
	"database/sql"
	"github.com/pitwch/go-proffix-geocode/adresses"
	"github.com/pitwch/go-proffix-geocode/geocode"
	"testing"
)

func TestGeocodeGoogle(t *testing.T) {
	testaddresses := []Adresses.Adresses{
		{
			1,
			sql.NullString{"Murgtalstrasse 20", true},
			sql.NullString{"9542", true},
			sql.NullString{"Münchwilen", true},
			sql.NullString{"CH", true},
			sql.NullString{"", true},
			sql.NullString{"", true},
		},
	}
	geocode.GeocodeAdressesGoogle(testaddresses)
}

func TestGeocodeMapquest(t *testing.T) {
	testaddresses := []Adresses.Adresses{
		{
			1,
			sql.NullString{"Murgtalstrasse 20", true},
			sql.NullString{"9542", true},
			sql.NullString{"Münchwilen", true},
			sql.NullString{"CH", true},
			sql.NullString{"", true},
			sql.NullString{"", true},
		},
	}
	geocode.GeocodeAdressesGoogle(testaddresses)
}
