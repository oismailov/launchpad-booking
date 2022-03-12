package db

import "github.com/oismailov/launchpad-booking/model"

func SaveBooking(booking *model.Booking) error {
	if err := GetInstance().Save(booking).Error; err != nil {
		return err
	}

	return nil
}

func GetDuplicatedBookingsCount(booking model.Booking) (int64, error) {
	var result int64

	if err := GetInstance().
		Where(&model.Booking{
			LaunchpadID:   booking.LaunchpadID,
			LaunchDate:    booking.LaunchDate,
			DestinationID: booking.DestinationID,
		}).
		Count(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func GetAllBookings() ([]model.Booking, error) {
	var bookings []model.Booking

	if err := GetInstance().Find(&bookings).Error; err != nil {
		return bookings, err
	}

	return bookings, nil
}
