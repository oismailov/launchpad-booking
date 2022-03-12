package svc

import (
	"github.com/oismailov/launchpad-booking/db"
	"github.com/oismailov/launchpad-booking/model"
)

func SaveBooking(booking *model.Booking) error {
	return db.SaveBooking(booking)
}

func GetBookingsByLaunchpadIdAndLaunchDate(booking model.Booking) ([]model.Booking, error) {
	return db.GetBookingsByLaunchpadIdAndLaunchDate(booking)
}

func GetAllBookings() ([]model.Booking, error) {
	return db.GetAllBookings()
}

func DeleteBookingByUUID(bookingUuid string) error {
	return db.DeleteBookingByUuid(bookingUuid)
}
