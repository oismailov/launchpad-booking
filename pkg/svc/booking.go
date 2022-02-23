package svc

import (
	"github.com/oismailov/launchpad-booking/db"
	"github.com/oismailov/launchpad-booking/model"
)

func SaveBooking(booking *model.Booking) error {
	return db.SaveBooking(booking)
}

func GetBookings(booking model.Booking) ([]model.Booking, error) {
	return db.GetBookings(booking)
}
