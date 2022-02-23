package db

import (
	"github.com/oismailov/launchpad-booking/model"
	"time"
)

func GetLaunchByLaunchpadIDAndDate(launchpadID string, date time.Time) (model.SpaseXLaunch, error) {
	launch := model.SpaseXLaunch{}

	if err := GetInstance().Where(&model.SpaseXLaunch{LaunchpadID: launchpadID, DateLocal: date}).Find(&launch).Error; err != nil {
		return launch, err
	}

	return launch, nil
}
