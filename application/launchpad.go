package application

import (
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/persistance"
)

func GetLaunchpadByID(launchpadID string) (model.SpaseXLaunchpad, error) {
	return persistance.GetLaunchpadByID(launchpadID)
}
