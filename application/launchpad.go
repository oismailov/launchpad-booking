package application

import (
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/repository"
)

func GetLaunchpadByID(launchpadID string) (model.SpaceXLaunchpad, error) {
	return repository.GetLaunchpadByID(launchpadID)
}
