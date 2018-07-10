package database

import (
	_ "github.com/denisenkom/go-mssqldb" // MSSQL driver
	"github.com/jmoiron/sqlx"
	"log"

	"fmt"
)

var (
	// SQL wrapper
	SQL *sqlx.DB
)

type Database struct {
	Username  string
	Password  string
	Database  string
	Hostname  string
	Port      int
	Parameter string
}

// Connect to PROFFIX Database
func Connect(database Database) {

	var err error

	// Store the config
	//databases = d

	sqlx.NameMapper = func(s string) string { return s }

	// Connect to MSSQL
	if SQL, err = sqlx.Connect("sqlserver", DSNMS(database)); err != nil {
		log.Println("SQL Driver Error"+DSNMS(database), err)
	}

	// Check if is alive
	if err = SQL.Ping(); err != nil {
		log.Println("Database Error", err)
	}
	sqlx.NameMapper = func(s string) string { return s }

}

// Create Connect String
func DSNMS(ci Database) string {
	// Example: root:@tcp(localhost:3306)/test
	return "sqlserver://" +
		ci.Username +
		":" +
		ci.Password +
		"@" +
		ci.Hostname +
		":" +
		fmt.Sprintf("%d", ci.Port) +
		"?database=" +
		ci.Database
}
