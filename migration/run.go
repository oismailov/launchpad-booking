package main

import (
	"github.com/oismailov/launchpad-booking/config"
	"github.com/oismailov/launchpad-booking/pkg/migration"
)

func main() {
	config.LoadConfig()
	migration.Migrate()
}
