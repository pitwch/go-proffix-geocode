[![Build Status](https://travis-ci.org/pitwch/go-proffix-geocode.svg?branch=master)](https://travis-ci.org/pitwch/go-proffix-geocode)
[![Coverage Status](https://coveralls.io/repos/github/pitwch/go-proffix-geocode/badge.svg?branch=master)](https://coveralls.io/github/pitwch/go-proffix-geocode?branch=master)

# Geocode PROFFIX Adressen

Mit diesem Tool können Adressen aus PROFFIX automatisch geocodiert werden.

Das Tool liest sämtliche Adressen ohne Longitude / Latitude, geocodiert diese wahlweise über Google oder Mapquest
und schreibt die Longitude / Latitude Werte zurück in die Datenbank.

## Verwendung

Sämtliche Konfiguration erfolgt über eine **config.json** welche sich im selben Verzeichnis wie das Tool befinden muss.

Beispiel:

```json
{
  "Database": {
    "Username": "sa",
    "Password": "1234",
    "Database": "PX000001",
    "Hostname": "SERVER1",
    "Port": 1433,
    "Parameter": "?connection+timeout=30"
  },
  "Settings": {
    "UseGeocoder": "Google",
    "MapQuestAPIKey": "",
    "GoogleAPIKey": "AgOleyXsXAbzEuSPzRYf_vbdT943s44w1nRkW"
  }
}

```

| Parameter    | Typ    | Bemerkung                                                              |
|--------------|--------|------------------------------------------------------------------------|
| Username     | string | Benutzername SQL - Server                                              |
| Password     | string | Passwort SQL - Server                                                  |
| Database     | string | PROFFIX Datenbank                                                      |
| Hostname     | string | Hostname des SQL - Servers; kann auch IP sein                          |
| Port         | int    | Port des SQL Servers                                                   |
| Parameter    | string | Div. Parameter für SQL-Server                                          |
| UseGeocoder  | string | "Google" / "Mapquest" ; legt fest mit welchem Provider gearbeitet wird |
| MapQuestAPI  | string | API - Key für Mapquest                                                 |
| GoogleAPIKey | string | API - Key für Google                                                   |

