package main

import (
	
	"pos-app/config"
	"pos-app/routes"
)

func main() {
	
	config.DatabaseInit()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":1313"))
}