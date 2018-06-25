package Adresses

import (
	"database/sql"
	"errors"
	"github.com/pitwch/go-proffix-geocode/shared"
	"github.com/spf13/cast"
)

type Adresses struct {
	Adressnr  int            `json:"Adressnr" db:"AdressNrADR"`
	Street    sql.NullString `json:"Strasse" db:"Strasse"`
	PLZ       sql.NullString `json:"PLZ" db:"PLZ"`
	City      sql.NullString `json:"City" db:"Ort"`
	Country   sql.NullString `json:"Country" db:"LandPRO"`
	Longitude sql.NullString `json:"Longitude" db:"Longitude"`
	Latitude  sql.NullString `json:"Latitude" db:"Latitude"`
}

type Geocoded struct {
	Adressnr int
	Longitude string
	Latitude  string
}
var (
	ErrNoResult = errors.New("No non-geocoded Adresses")
)

const (
	readQueryAdrAdresse = "SELECT AdressNrADR,Strasse,PLZ,Ort,LandPRO,Longitude,Latitude FROM ADR_Adressen WHERE " +
		"Strasse IS NOT NULL AND PLZ IS NOT NULL AND Ort IS NOT NULL AND (Longitude IS NULL OR Longitude = 0) AND (Latitude IS NULL OR Latitude = 0)"
)

// Get all Adresses from PROFFIX where Long / Lat is not set and Street, Plz and City are set
func GetEmptyAdresses() ([]Adresses, error) {
	var results []Adresses

	err := database.SQL.Select(&results, readQueryAdrAdresse)

	if err == sql.ErrNoRows {
		err = ErrNoResult
	}
	return results, err
}

// Update just Long / Lat in PROFFIX Adresses
func UpdateEmptyAdresses(g Geocoded) (adressnr int, err error) {


	// Update the entity
	_, err = database.SQL.Exec(`
		UPDATE ADR_Adressen SET
		Longitude = @p1,
		Latitude = @p2
		WHERE AdressNrADR = @p3
		`,
		cast.ToFloat32(g.Longitude),
		cast.ToFloat32(g.Latitude),
		cast.ToInt(g.Adressnr))

	// If error occurred error
	if err != nil {
		return g.Adressnr, err
	}

	return g.Adressnr, nil
}