package db

import "github.com/oismailov/launchpad-booking/model"

func SaveBooking(booking *model.Booking) error {
	if err := GetInstance().Save(booking).Error; err != nil {
		return err
	}

	return nil
}

func GetBookingsByLaunchpadIdAndLaunchDate(booking model.Booking) ([]model.Booking, error) {
	var bookings []model.Booking

	if err := GetInstance().
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

	if err := GetInstance().Model(&model.Booking{}).Find(&bookings).Error; err != nil {
		return bookings, err
	}

	return bookings, nil
}
