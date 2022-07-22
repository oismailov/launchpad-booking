package repository

import (
	"github.com/oismailov/launchpad-booking/model"
	"time"
)

func GetLaunchByLaunchpadIdAndDate(launchpadID string, date time.Time) (model.SpaceXLaunch, error) {
	launch := model.SpaceXLaunch{}

	if err := GetConnection().
		Where(&model.SpaceXLaunch{
			LaunchpadID: launchpadID,
			DateLocal:   date,
		}).
		Find(&launch).Error; err != nil {
		return launch, err
	}

	return launch, nil
}
