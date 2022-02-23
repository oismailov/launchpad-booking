package main

import (
	"github.com/oismailov/launchpad-booking/cmd/launchpads"
	"github.com/oismailov/launchpad-booking/config"
	"log"
)

func main() {
	config.LoadConfig()
	parsedLaunchpads := launches.GetParsedLaunchpads()
	err := launches.SaveLaunchpadsToDB(parsedLaunchpads)
	if err != nil {
		log.Fatal(err.Error())
	}
}
