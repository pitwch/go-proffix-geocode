[![Build Status](https://travis-ci.org/pitwch/go-proffix-geocode.svg?branch=master)](https://travis-ci.org/pitwch/go-proffix-geocode)
[![codecov](https://codecov.io/gh/pitwch/go-proffix-geocode/branch/master/graph/badge.svg)](https://codecov.io/gh/pitwch/go-proffix-geocode)

# Geocode PROFFIX Adressen

Mit diesem Tool können Adressen aus PROFFIX automatisch geocodiert werden.

Das Tool liest sämtliche Adressen ohne Longitude / Latitude aus der angegebenen PROFFIX - Datenbank, geocodiert diese wahlweise über Google oder Mapquest
und schreibt die Longitude / Latitude Werte zurück in die PROFFIX - Datenbank.

*Beispiel Ergebnis geocodierter PROFFIX - Adressen in [pApp](https://www.proffixapp.ch)*
![alt text](https://raw.githubusercontent.com/pitwch/go-proffix-geocode/master/assets/img/geocode_proffix_papp.jpg "Beispiel Verwendung mit pApp")

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

Die aktuellste Version des Tools findet sich immer unter [Release](https://github.com/pitwch/go-proffix-geocode/releases/latest).

Alternativ kann das [ZIP-File hier direkt heruntergeladen](https://github.com/pitwch/go-proffix-geocode/files/2132369/go-proffix-geocode.zip) werden.

Dann das ZIP-File entpacken, die API-Keys besorgen und die config.json editieren.

Das Tool kann anschliessen in **Windows per CMD ausgeführt** werden:

![alt text](https://raw.githubusercontent.com/pitwch/go-proffix-geocode/master/assets/img/cmd_geocode.jpg "Kommandozeile PROFFIX geocodieren")
