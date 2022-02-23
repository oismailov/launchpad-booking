package model

import "time"

type SpaseXLaunches struct {
	ID            uint      `json:"-" gorm:"primary_key"`
	LaunchpadUuid string    `valid:"required" gorm:"type:char(36);launchpad" json:"launchpad"`
	DateLocal     time.Time `valid:"required" gorm:"date_local" json:"date_local"`

	BaseModel
}
