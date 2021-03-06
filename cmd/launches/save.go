package launches

import (
	"errors"
	"fmt"
	"github.com/oismailov/launchpad-booking/repository"
	"strings"
)

func SaveLaunchesToDB(records []spacexLaunch) error {
	if len(records) <= 0 {
		return errors.New("there are no records")
	}

	var valueStrings []string
	var valueArgs []interface{}

	for _, record := range records {
		valueStrings = append(valueStrings, "(?, ?, ?, NOW(), NOW())")

		valueArgs = append(valueArgs, record.LaunchID)
		valueArgs = append(valueArgs, record.LaunchpadID)
		valueArgs = append(valueArgs, record.DateLocal)
	}

	smt := "INSERT INTO space_x_launches " +
		"(launch_id, launchpad_id, date_local, created_at, updated_at) " +
		"VALUES %s " +
		"ON CONFLICT (launch_id, launchpad_id, date_local) " +
		"DO NOTHING"

	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	tx := repository.GetConnection().Begin()

	if err := tx.Exec(smt, valueArgs...).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
