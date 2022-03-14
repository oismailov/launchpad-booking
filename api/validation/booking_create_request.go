package validation

import (
	"errors"
	"github.com/oismailov/launchpad-booking/application"
	"github.com/oismailov/launchpad-booking/model"
	validator "gopkg.in/asaskevich/govalidator.v9"
	"time"
)

func ValidateCreateBookingRequest(booking model.Booking) error {
	if _, err := validator.ValidateStruct(booking); err != nil {
		return err
	}

	if !isValidLaunchDate(booking) {
		return errors.New("invalid launch_date")
	}

	if !isValidBirthDay(booking) {
		return errors.New("invalid birthday")
	}

	isValidBooking := isValidBooking(booking)
	if isValidBooking == false {
		return errors.New("unable to create a booking")
	}

	return nil
}

func isValidLaunchDate(booking model.Booking) bool {
	return validator.IsTime(booking.LaunchDate, "2006-01-02")
}

func isValidBirthDay(booking model.Booking) bool {
	return validator.IsTime(booking.Birthday, "2006-01-02")
}

func isValidBooking(booking model.Booking) bool {
	if !isValidLaunchpad(booking) || !isLaunchAvailable(booking) || hasDuplicatedBooking(booking) {
		return false
	}

	return true
}

func isValidLaunchpad(booking model.Booking) bool {
	launchpad, err := application.GetLaunchpadByID(booking.LaunchpadID)
	if err != nil {
		return false
	}

	if launchpad.ID == 0 {
		return false
	}

	return true
}

func isLaunchAvailable(booking model.Booking) bool {
	launchDate, err := time.Parse("2006-01-02", booking.LaunchDate)

	if err != nil {
		return false
	}

	launch, err := application.GetLaunchByLaunchpadIdAndDate(booking.LaunchpadID, launchDate)
	if err != nil {
		return false
	}

	if launch.ID != 0 {
		return false
	}

	return true
}

func hasDuplicatedBooking(booking model.Booking) bool {
	bookings, err := application.GetBookingsByLaunchpadIdAndLaunchDate(booking)
	if err != nil {
		return true
	}

	if len(bookings) > 0 && bookings[0].DestinationID != booking.DestinationID {
		return true
	}

	return false
}
