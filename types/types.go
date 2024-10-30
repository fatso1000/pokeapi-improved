package types

import "database/sql"

type SavePokemonBody struct {
	PokemonId int `json:"pokemon_id"`
	UserId    int `json:"user_id"`
}

type Record struct {
	Id              int           `json:"id"`
	Identifier      string        `json:"identifier"`
	Species_id      int           `json:"species_id"`
	Height          int           `json:"height"`
	Weight          int           `json:"weight"`
	Base_experience sql.NullInt64 `json:"base_experience"`
	Order_number    sql.NullInt64 `json:"order_number"`
	Is_default      int           `json:"is_default"`
	Url             string        `json:"url"`
}
