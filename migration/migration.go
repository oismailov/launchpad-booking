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

	if err != nil {
		fmt.Printf("unable to run migrations : %v", err)
	}
}
