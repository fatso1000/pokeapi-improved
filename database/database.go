package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func StartService() {

	// defer db.Close()
	db := DbService()

	// INITIALIZE SERVICES
	createPokemonTable(db)
	// bulkDataToPokemonTable(db)
	createUsersTable(db)
	createSavedPokemonTable(db)
}

func DbService() (db *sql.DB) {
	// NEED TO ENTER CREDENTIALS
	connStr := "postgres://postgres:{{password}}@localhost:{{port}}/pokemon?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func createPokemonTable(db *sql.DB) {
	/*
			-	"id"  int
		    -	"identifier" varchar
		    -	"species_id" int
		    -	"height" int
		    -	"weight" int
		    -	"base_experience" int
		    -	"order" int
		    -	"is_default" int
	*/
	query := `CREATE TABLE IF NOT EXISTS pokemon (
		id INTEGER PRIMARY KEY,
		identifier VARCHAR(255) NOT NULL,
		species_id INTEGER,
		height INTEGER,
		weight INTEGER,
 		base_experience INTEGER,
 		order_number INTEGER,
		is_default INTEGER
	)`

	_, err := db.Exec(query)
	check(err, "pokemon")
}

func createUsersTable(db *sql.DB) {
	/*
			-	"id"  int
		    -	"identifier" varchar
		    -	"species_id" int
		    -	"height" int
		    -	"weight" int
		    -	"base_experience" int
		    -	"order" int
		    -	"is_default" int
	*/
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL
	)`

	_, err := db.Exec(query)
	check(err, "users")
}

func createSavedPokemonTable(db *sql.DB) {
	/*
			-	"id"  int
		    -	"identifier" varchar
		    -	"species_id" int
		    -	"height" int
		    -	"weight" int
		    -	"base_experience" int
		    -	"order" int
		    -	"is_default" int
	*/
	query := `CREATE TABLE IF NOT EXISTS user_pokemon (
		user_id INTEGER REFERENCES users(id),
		pokemon_id INTEGER REFERENCES pokemon(id),
		CONSTRAINT saved_pk PRIMARY KEY(pokemon_id,user_id)
	)`

	_, err := db.Exec(query)
	check(err, "user_pokemon")
}

// BULK NOT WORKING, TRY TO FIND A SOLUTION BY READING THE CSV FILE AND USING A FOR TO EXEC EACH ROW
// func bulkDataToPokemonTable(db *sql.DB) {
// 	filePath, errFilePath := filepath.Abs("assets/pokemon.csv")
// 	check(errFilePath)
// 	query := `\copy pokemon FROM '` + filePath + `' WITH (FORMAT CSV, HEADER)`

// 	_, err := db.Exec(query)
// 	check(err)
// }

func check(e error, space string) {
	if e != nil {
		log.Default().Print(space)
		log.Fatal(e)
	}
}
