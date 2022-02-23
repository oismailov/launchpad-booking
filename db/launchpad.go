package db

import "github.com/oismailov/launchpad-booking/model"

func GetLaunchpadByID(launchpadID string) (model.SpaseXLaunchpad, error) {
	launchpad := model.SpaseXLaunchpad{}

	if err := GetInstance().Where(&model.SpaseXLaunchpad{LaunchpadID: launchpadID}).Find(&launchpad).Error; err != nil {
		return launchpad, err
	}

	return launchpad, nil
}
