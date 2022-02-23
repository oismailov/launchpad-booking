package model

import "time"

type SpaseXLaunch struct {
	ID          uint      `json:"-" gorm:"primary_key"`
	LaunchID    string    `valid:"required"  json:"launch_id"`
	LaunchpadID string    `valid:"required"  json:"launchpad_id"`
	DateLocal   time.Time `valid:"required" json:"date_local"`

	BaseModel
}
