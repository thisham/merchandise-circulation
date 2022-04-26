package main

import (
	"fmt"
	"merchandise-circulation-api/src/configs"
	"merchandise-circulation-api/src/database"
	"merchandise-circulation-api/src/routes"
)

var srv configs.ServerConfig

func init() {
	srv, _ = configs.LoadServerConfig(".")
	var DB *database.DBConf
	DB.InitDB().Migrate()
}

func main() {
	api := routes.New()
	api.Logger.Fatal(api.Start(fmt.Sprintf(
		"%v:%v", srv.ServerHost, srv.ServerPort,
	)))
}
