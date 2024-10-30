package main

import (
	"main/database"
	"main/routes"
)

func main() {
	database.StartService()
	routes.StartService()
}
