package repository

import "github.com/oismailov/launchpad-booking/model"

func SaveBooking(booking *model.Booking) error {
	if err := GetConnection().Save(booking).Error; err != nil {
		return err
	}

	return nil
}

func GetBookingsByLaunchpadIdAndLaunchDate(booking model.Booking) ([]model.Booking, error) {
	var bookings []model.Booking

	if err := GetConnection().
		Where(&model.Booking{
			LaunchpadID: booking.LaunchpadID,
			LaunchDate:  booking.LaunchDate,
		}).
		Find(&bookings).Error; err != nil {
		return bookings, err
	}

	return bookings, nil
}

func GetAllBookings() ([]model.Booking, error) {
	var bookings []model.Booking

	if err := GetConnection().Model(&model.Booking{}).Find(&bookings).Error; err != nil {
		return bookings, err
	}

	return bookings, nil
}

func DeleteBookingByUUID(bookingUUID string) error {
	if err := GetConnection().
		Where(&model.Booking{
			UUID: bookingUUID,
		}).
		Delete(&model.Booking{}).
		Error; err != nil {
		return err
	}

	return nil
}
