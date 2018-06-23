# Geocode PROFFIX Adressen

Mit diesem Tool können Adressen aus PROFFIX automatisch geocodiert werden.

Das Tool liest sämtliche Adressen ohne Longitude / Latitude, geocodiert diese wahlweise über Google oder Mapquest
und schreibt die Longitude / Latitude Werte zurück in die Datenbank.

## Verwendung

Sämtliche Konfiguration erfolgt über eine **config.json** welche sich im selben Verzeichnis wie das Tool befinden muss.

Beispiel:


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


