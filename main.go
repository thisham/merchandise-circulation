package main

import (
	"merchandise-circulation-api/src/routes"
)

func main() {
	api := routes.New()
	api.Logger.Fatal(api.Start(":8000"))
}
