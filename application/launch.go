package application

import (
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/persistance"
	"time"
)

func GetLaunchByLaunchpadIdAndDate(launchpadID string, date time.Time) (model.SpaseXLaunch, error) {
	return persistance.GetLaunchByLaunchpadIdAndDate(launchpadID, date)
}
