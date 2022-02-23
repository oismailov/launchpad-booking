package svc

import (
	"github.com/oismailov/launchpad-booking/db"
	"github.com/oismailov/launchpad-booking/model"
)

func GetLaunchpadByID(launchpadID string) (model.SpaseXLaunchpad, error) {
	return db.GetLaunchpadByID(launchpadID)
}
