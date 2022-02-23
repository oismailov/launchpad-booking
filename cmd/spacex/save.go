package spacex

import (
	"fmt"
	"github.com/oismailov/launchpad-booking/db"
	"log"
	"strings"
)

func SaveToDB(records []spacexLaunch) error {
	if len(records) <= 0 {
		log.Println("There are no records")
	}

	var valueStrings []string
	var valueArgs []interface{}

	for _, record := range records {
		valueStrings = append(valueStrings, "(?, ?, NOW(), NOW())")

		valueArgs = append(valueArgs, record.Launchpad)
		valueArgs = append(valueArgs, record.DateLocal)
	}

	smt := "INSERT INTO spase_x_launches " +
		"(launchpad_uuid, date_local, created_at, updated_at) " +
		"VALUES %s " +
		"ON CONFLICT (launchpad_uuid, date_local) " +
		"DO NOTHING"

	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	tx := db.GetInstance().Begin()

	if err := tx.Exec(smt, valueArgs...).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
