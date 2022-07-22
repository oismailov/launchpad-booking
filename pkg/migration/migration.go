package migration

import (
	"fmt"
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/repository"
)

func Migrate() {
	dbInstance := repository.GetConnection()

	err := dbInstance.AutoMigrate(
		&model.SpaceXLaunch{},
		&model.SpaceXLaunchpad{},
		&model.Booking{},
	)

	dbInstance.Exec("CREATE UNIQUE INDEX IF NOT EXISTS space_x_unique_launches ON space_x_launches (launch_id, launchpad_id, date_local)")
	dbInstance.Exec("CREATE UNIQUE INDEX IF NOT EXISTS space_x_unique_launchpads ON space_x_launchpads (launchpad_id)")

	if err != nil {
		fmt.Printf("unable to run migrations : %v", err)
	}
}
