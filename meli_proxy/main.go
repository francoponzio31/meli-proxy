package main

import (
	"log"
	"meli_proxy/app"
)

func main() {
	config := app.LoadConfig()
	meli_proxy := app.CreateApp()

	log.Fatal(meli_proxy.Run(config.AppHost + ":" + config.AppPort))
}
