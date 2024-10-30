package v1_routes

import (
	"encoding/json"
	"main/database/repository"
	"main/types"

	"github.com/gofiber/fiber/v2"
)

func StartService(apiInstance fiber.Router) {
	v1 := apiInstance.Group("/v1")
	// POKEMONS
	v1.Get("/pokemons", func(c *fiber.Ctx) error {
		pageNumber := c.Query("page_number")
		pageSize := c.Query("page_size")
		pokemons := repository.GetAllPokemons(pageNumber, pageSize)

		return c.JSON(fiber.Map{"data": pokemons})
	})
	v1.Get("/findPokemon", func(c *fiber.Ctx) error {
		name := c.Query("name")
		pokemons := repository.GetPokemonByName(name)
		return c.JSON(fiber.Map{"data": pokemons})
	})
	// USER
	v1.Post("/user", func(c *fiber.Ctx) error {
		username := c.Query("username")
		repo := repository.SaveUser(username)
		return c.SendString(repo)
	})
	// SAVED
	v1.Get("/savePokemons", func(c *fiber.Ctx) error {
		userId := c.Query("userId")
		pokemons := repository.GetSavedPokemons(userId)
		return c.JSON(fiber.Map{"data": pokemons})
	})
	v1.Post("/savePokemon", func(c *fiber.Ctx) error {
		body := c.Body()
		var bodyPokemon types.SavePokemonBody
		json.Unmarshal(body, &bodyPokemon)
		repo := repository.SavePokemonToUser(bodyPokemon)
		return c.JSON(repo)
	})
}
