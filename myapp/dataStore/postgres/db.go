// dataStore/postgres/postgres.go
package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "dpg-cp0p70o21fec7388htdg-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "fhsg7uJLWI83BNUIpzTujbY1iognxJIb"
	postgres_dbname   = "my_db_y3qn"
)

// create pointser var Db which points to sql driver
var Db *sql.DB

// init() is always called before main
func init() {
	// creating the connection string
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error
	// open the connection to database
	Db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database successfully configured")

}
