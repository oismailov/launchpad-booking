package main

import (
	"github.com/oismailov/launchpad-booking/cmd/launches"
	"github.com/oismailov/launchpad-booking/config"
	"log"
)

func main() {
	config.LoadConfig()
	parsedLaunches, err := launches.GetParsedLaunches()
	if err != nil {
		log.Fatal("unable to parse launches: ", err.Error())
		return
	}

	err = launches.SaveLaunchesToDB(parsedLaunches)
	if err != nil {
		log.Fatal("unable to save launches: ", err.Error())
		return
	}
}
