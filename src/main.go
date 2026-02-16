package main

import (
	"lppm/src/database"
	"lppm/src/routes"
)

func main() {
	database.InitDB()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
