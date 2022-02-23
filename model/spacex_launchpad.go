package model

type SpaseXLaunchpad struct {
	ID          uint   `json:"-" gorm:"primary_key"`
	LaunchpadID string `valid:"required" json:"launchpad_id"`

	BaseModel
}
