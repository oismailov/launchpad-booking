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

	err := dbInstance.AutoMigrate(&model.SpaseXLaunches{})

	dbInstance.Exec("CREATE UNIQUE INDEX spase_x_launches_date_local_unique ON spase_x_launches (launchpad_uuid, date_local)")

	if err != nil {
		fmt.Printf("unable to run migrations : %v", err)
	}
}
