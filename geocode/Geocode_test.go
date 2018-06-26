package geocode

import (
	"testing"
)

func TestSetSettings(t *testing.T) {
	testSetting := Settings{"Google", "", ""}
	SetSettings(testSetting)
}

//func TestGeocodeAdresses (t *testing.T) {
//
//	testaddresses := []Adresses.Adressstruct{
//		Adresses.Adressstruct{
//			1,
//			sql.NullString{"Murgtalstrasse 20", true},
//			sql.NullString{"9542", true},
//			sql.NullString{"Münchwilen", true},
//			sql.NullString{"CH", true},
//			sql.NullString{"", true},
//			sql.NullString{"", true},
//		},
//	}
//	GeocodeAdresses(testaddresses)

