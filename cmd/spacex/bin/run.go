package main

import (
	"github.com/oismailov/launchpad-booking/cmd/spacex"
	"github.com/oismailov/launchpad-booking/config"
)

func main() {
	config.LoadConfig()
	launches := spacex.GetParsedData()
	spacex.SaveToDB(launches)
}
