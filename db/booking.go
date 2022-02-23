package db

import "github.com/oismailov/launchpad-booking/model"

func SaveBooking(booking *model.Booking) error {
	if err := GetInstance().Save(booking).Error; err != nil {
		return err
	}

	return nil
}

func GetBookings(booking model.Booking) ([]model.Booking, error) {
	var result []model.Booking

	if err := GetInstance().
		Where(&model.Booking{LaunchpadID: booking.LaunchpadID, LaunchDate: booking.LaunchDate}).
		Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
