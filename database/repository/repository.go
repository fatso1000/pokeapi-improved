package repository

import (
	"database/sql"
	"log"
	"main/database"
	"main/types"
	"strconv"
)

var (
	id              int
	identifier      string
	species_id      int
	height          int
	weight          int
	base_experience sql.NullInt64
	order_number    sql.NullInt64
	is_default      int
)

func GetAllPokemons(pageNumber string, pageSize string) []types.Record {
	db := database.DbService()

	records := []types.Record{}

	query := "SELECT * FROM pokemon ORDER BY id OFFSET (" + pageNumber + " - 1) * " + pageSize + " FETCH NEXT " + pageSize + " ROWS ONLY"
	rows, err := db.Query(query)
	check(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &identifier, &species_id, &height, &weight, &base_experience, &order_number, &is_default)
		check(err)
		records = append(records, types.Record{Id: id, Identifier: identifier, Species_id: species_id, Height: height, Weight: weight, Base_experience: base_experience, Order_number: order_number, Is_default: is_default, Url: "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(id)})
	}

	return records
}

func GetPokemonByName(name string) []types.Record {
	db := database.DbService()

	records := []types.Record{}
	query := "SELECT * FROM pokemon WHERE identifier LIKE '%" + name + "%'"
	rows, err := db.Query(query)
	check(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &identifier, &species_id, &height, &weight, &base_experience, &order_number, &is_default)
		check(err)
		records = append(records, types.Record{Id: id, Identifier: identifier, Species_id: species_id, Height: height, Weight: weight, Base_experience: base_experience, Order_number: order_number, Is_default: is_default})
	}

	return records
}

func SaveUser(username string) string {
	db := database.DbService()
	query := "INSERT INTO users (username) VALUES ('" + username + "')"
	_, err := db.Exec(query)
	check(err)
	defer db.Close()

	return "User created successfully"
}

func SavePokemonToUser(body types.SavePokemonBody) string {
	db := database.DbService()
	query := "INSERT INTO user_pokemon (user_id, pokemon_id) VALUES ('" + strconv.Itoa(body.UserId) + "', '" + strconv.Itoa(body.PokemonId) + "')"
	_, err := db.Exec(query)
	check(err)
	defer db.Close()

	return "Pokemon to User saved successfully"
}

func GetSavedPokemons(userId string) []types.Record {
	db := database.DbService()

	records := []types.Record{}
	query := "SELECT pokemon.* FROM user_pokemon JOIN pokemon on pokemon.id = user_pokemon.pokemon_id JOIN users on users.id = user_pokemon.user_id WHERE user_id = '" + userId + "'"
	rows, err := db.Query(query)
	check(err)
	defer db.Close()

	for rows.Next() {
		err := rows.Scan(&id, &identifier, &species_id, &height, &weight, &base_experience, &order_number, &is_default)
		check(err)
		records = append(records, types.Record{Id: id, Identifier: identifier, Species_id: species_id, Height: height, Weight: weight, Base_experience: base_experience, Order_number: order_number, Is_default: is_default})
	}

	return records
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
