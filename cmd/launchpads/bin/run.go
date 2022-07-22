package main

import (
	"github.com/oismailov/launchpad-booking/cmd/launchpads"
	"github.com/oismailov/launchpad-booking/config"
	"log"
)

func main() {
	config.LoadConfig()
	parsedLaunchpads, err := launches.GetParsedLaunchpads()
	if err != nil {
		log.Fatal("unable to parse launchpads: ", err.Error())
	}

	err = launches.SaveLaunchpadsToDB(parsedLaunchpads)
	if err != nil {
		log.Fatal("unable to save launchpads: ", err.Error())
	}
}
