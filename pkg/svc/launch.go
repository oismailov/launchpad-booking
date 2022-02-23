package svc

import (
	"github.com/oismailov/launchpad-booking/db"
	"github.com/oismailov/launchpad-booking/model"
	"time"
)

func GetLaunchByLaunchpadIDAndDate(launchpadID string, date time.Time) (model.SpaseXLaunch, error) {
	return db.GetLaunchByLaunchpadIDAndDate(launchpadID, date)
}