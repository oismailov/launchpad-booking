package main

import (
	"fmt"
	"github.com/oismailov/launchpad-booking/config"
	"github.com/oismailov/launchpad-booking/db"
	"github.com/oismailov/launchpad-booking/model"
)

func main() {
	config.LoadConfig()
	dbInstance := db.GetInstance()

	err := dbInstance.AutoMigrate(
		&model.SpaseXLaunch{},
		&model.SpaseXLaunchpad{},
		&model.Booking{},
	)

	dbInstance.Exec("CREATE UNIQUE INDEX IF NOT EXISTS spase_x_unique_launches ON spase_x_launches (launch_id, launchpad_id, date_local)")
	dbInstance.Exec("CREATE UNIQUE INDEX IF NOT EXISTS spase_x_unique_launchpads ON spase_x_launchpads (launchpad_id)")

	if err != nil {
		fmt.Printf("unable to run migrations : %v", err)
	}
}
