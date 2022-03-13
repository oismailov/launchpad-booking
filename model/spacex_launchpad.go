package model

type SpaceXLaunchpad struct {
	ID          uint   `json:"-" gorm:"primary_key"`
	LaunchpadID string `valid:"required" json:"launchpad_id"`

	BaseModel
}
