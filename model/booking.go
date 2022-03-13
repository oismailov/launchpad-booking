package model

type Booking struct {
	ID            uint   `json:"-" gorm:"primary_key"`
	UUID          string `json:"uuid" gorm:"unique"`
	FirstName     string `valid:"required,alpha"  json:"first_name"`
	LastName      string `valid:"required,alpha"  json:"last_name"`
	Gender        string `valid:"required,alpha"  json:"gender"`
	Birthday      string `valid:"required"  json:"birthday"`
	LaunchpadID   string `valid:"required"  json:"launchpad_id"`
	DestinationID string `valid:"required"  json:"destination_id"`
	LaunchDate    string `valid:"required" json:"launch_date" gorm:"date"`

	BaseModel
}
