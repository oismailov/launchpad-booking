package model

import (
	"github.com/jackc/pgtype"
	"time"
)

type Booking struct {
	ID            uint        `json:"-" gorm:"primary_key"`
	FirstName     string      `valid:"required,alpha"  json:"first_name"`
	LastName      string      `valid:"required,alpha"  json:"last_name"`
	Gender        string      `valid:"required"  json:"gender"`
	Birthday      pgtype.Date `valid:"required"  json:"birthday"`
	LaunchpadId   string      `valid:"required"  json:"launchpad_id"`
	DestinationId string      `valid:"required"  json:"destination_id"`
	LaunchDate    time.Time   `valid:"required" json:"launch_date"`

	BaseModel
}
