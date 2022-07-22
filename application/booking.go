package application

import (
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/repository"
)

func SaveBooking(booking *model.Booking) error {
	return repository.SaveBooking(booking)
}

func GetBookingsByLaunchpadIdAndLaunchDate(booking model.Booking) ([]model.Booking, error) {
	return repository.GetBookingsByLaunchpadIdAndLaunchDate(booking)
}

func GetAllBookings() ([]model.Booking, error) {
	return repository.GetAllBookings()
}

func DeleteBookingByUUID(bookingUUID string) error {
	return repository.DeleteBookingByUUID(bookingUUID)
}
