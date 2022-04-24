package main

import (
	"fmt"
	"merchandise-circulation-api/src/configs"
	"merchandise-circulation-api/src/routes"
)

func main() {
	c, _ := configs.LoadServerConfig(".")
	api := routes.New()
	deploy := fmt.Sprintf("%v:%v", c.ServerHost, c.ServerPort)
	api.Logger.Fatal(api.Start(deploy))
}
