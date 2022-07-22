package launches

import (
	"fmt"
	"github.com/oismailov/launchpad-booking/repository"
	"log"
	"strings"
)

func SaveLaunchpadsToDB(records []spacexLaunchPad) error {
	if len(records) <= 0 {
		log.Println("There are no records")
	}

	var valueStrings []string
	var valueArgs []interface{}

	for _, record := range records {
		valueStrings = append(valueStrings, "(?, NOW(), NOW())")

		valueArgs = append(valueArgs, record.ID)
	}

	smt := "INSERT INTO space_x_launchpads " +
		"(launchpad_id, created_at, updated_at) " +
		"VALUES %s " +
		"ON CONFLICT (launchpad_id) " +
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
