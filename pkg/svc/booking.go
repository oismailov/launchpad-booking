package svc

import (
	"github.com/oismailov/launchpad-booking/db"
	"github.com/oismailov/launchpad-booking/model"
)

func SaveBooking(booking *model.Booking) error {
	return db.SaveBooking(booking)
}

func GetDuplicatedBookingsCount(booking model.Booking) (int64, error) {
	return db.GetDuplicatedBookingsCount(booking)
}

func GetAllBookings() ([]model.Booking, error) {
	return db.GetAllBookings()
}
