package main

import (
	"github.com/oismailov/launchpad-booking/config"
	r "github.com/oismailov/launchpad-booking/router"
	"log"
)

func main() {
	config.LoadConfig()

	router := r.GetRouter()

	err := router.Run(r.GetPort())

	if err != nil {
		log.Fatal("unable to start service")
	}
}
