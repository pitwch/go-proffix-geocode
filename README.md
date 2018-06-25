[![Build Status](https://travis-ci.org/pitwch/go-proffix-geocode.svg?branch=master)](https://travis-ci.org/pitwch/go-proffix-geocode)
[![Coverage Status](https://coveralls.io/repos/github/pitwch/go-proffix-geocode/badge.svg?branch=master)](https://coveralls.io/github/pitwch/go-proffix-geocode?branch=master)

# Geocode PROFFIX Adressen

Mit diesem Tool können Adressen aus PROFFIX automatisch geocodiert werden.

Das Tool liest sämtliche Adressen ohne Longitude / Latitude, geocodiert diese wahlweise über Google oder Mapquest
und schreibt die Longitude / Latitude Werte zurück in die Datenbank.

## Konfiguration

Sämtliche Konfiguration erfolgt über eine **config.json** welche sich **im selben Verzeichnis wie das Tool** befinden muss.

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
| GoogleAPIKey | string | API - Key für Google (empfohlen)                                       |


### API - Key ###

Sowohl für das Geocodieren über Google als auch über Mapquest **wird ein API - Key benötigt**. Diesen erhält man wie folgt:

**Google (empfohlen):** [https://developers.google.com/maps/documentation/geocoding/get-api-key](https://developers.google.com/maps/documentation/geocoding/get-api-key)

**Mapquest:** [https://developer.mapquest.com/documentation/](https://developer.mapquest.com/documentation/)


Die entsprechenden Lizenzbedingungen beachten!
(Im Falle von [pApp - dem App für PROFFIX](https://www.proffixapp.ch) ist alles ok)


## Verwendung

Das Tool kann anschliessen in Windows per CMD ausgeführt werden:

![alt text](https://raw.githubusercontent.com/pitwch/go-proffix-geocode/master/assets/img/cmd_geocode.jpg "Kommandozeile PROFFIX geocodieren")
