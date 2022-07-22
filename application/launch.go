package application

import (
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/repository"
	"time"
)

func GetLaunchByLaunchpadIdAndDate(launchpadID string, date time.Time) (model.SpaceXLaunch, error) {
	return repository.GetLaunchByLaunchpadIdAndDate(launchpadID, date)
}
