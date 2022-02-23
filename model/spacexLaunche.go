package model

import "time"

type SpaseXLaunches struct {
	ID            uint      `json:"-" gorm:"primary_key"`
	LaunchpadUuid string    `valid:"required"  json:"launchpad"`
	DateLocal     time.Time `valid:"required" json:"date_local"`

	BaseModel
}
