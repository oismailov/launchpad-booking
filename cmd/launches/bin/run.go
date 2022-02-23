package main

import (
	"github.com/oismailov/launchpad-booking/cmd/launches"
	"github.com/oismailov/launchpad-booking/config"
	"log"
)

func main() {
	config.LoadConfig()
	parsedLaunches := launches.GetParsedLaunches()
	err := launches.SaveLaunchesToDB(parsedLaunches)
	if err != nil {
		log.Fatal(err.Error())
	}
}
