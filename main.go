package main

import (
	"fmt"
	"merchandise-circulation-api/src/configs"
	"merchandise-circulation-api/src/database"
	"merchandise-circulation-api/src/routes"
)

func init() {
	database.Migrate()
}

func main() {
	c, _ := configs.LoadServerConfig(".")
	api := routes.New()
	server := fmt.Sprintf("%v:%v", c.ServerHost, c.ServerPort)
	api.Logger.Fatal(api.Start(server))
}
