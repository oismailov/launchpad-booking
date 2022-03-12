package application

import (
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/persistance"
)

func SaveBooking(booking *model.Booking) error {
	return persistance.SaveBooking(booking)
}

func GetBookingsByLaunchpadIdAndLaunchDate(booking model.Booking) ([]model.Booking, error) {
	return persistance.GetBookingsByLaunchpadIdAndLaunchDate(booking)
}

func GetAllBookings() ([]model.Booking, error) {
	return persistance.GetAllBookings()
}

func DeleteBookingByUUID(bookingUUID string) error {
	return persistance.DeleteBookingByUUID(bookingUUID)
}
