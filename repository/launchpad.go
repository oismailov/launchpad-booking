package repository

import "github.com/oismailov/launchpad-booking/model"

func GetLaunchpadByID(launchpadID string) (model.SpaceXLaunchpad, error) {
	launchpad := model.SpaceXLaunchpad{}

	if err := GetConnection().
		Where(&model.SpaceXLaunchpad{LaunchpadID: launchpadID}).
		Find(&launchpad).Error; err != nil {
		return launchpad, err
	}

	return launchpad, nil
}
