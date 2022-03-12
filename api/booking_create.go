package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/pkg/svc"
	"github.com/oismailov/launchpad-booking/pkg/util"
	validator "gopkg.in/asaskevich/govalidator.v9"
	"net/http"
	"time"
)

func CreateBooking(c *gin.Context) {
	var booking model.Booking
	err := c.BindJSON(&booking)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: "Invalid request body",
		})

		return
	}

	if _, err := validator.ValidateStruct(booking); err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: err.Error(),
		})

		return
	}

	isValidBooking := isValidBooking(booking)

	if isValidBooking == false {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: "Unable to create a booking.",
		})

		return
	}

	err = svc.SaveBooking(&booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, booking)
}

func isValidBooking(booking model.Booking) bool {
	if !isValidLaunchpad(booking) || !isLaunchAvailable(booking) || hasDuplicatedBooking(booking) {
		return false
	}

	return true
}

func isValidLaunchpad(booking model.Booking) bool {
	launchpad, err := svc.GetLaunchpadByID(booking.LaunchpadID)
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

	launch, err := svc.GetLaunchByLaunchpadIdAndDate(booking.LaunchpadID, launchDate)
	if err != nil {
		return false
	}

	if launch.ID != 0 {
		return false
	}

	return true
}

func hasDuplicatedBooking(booking model.Booking) bool {
	bookings, err := svc.GetBookingsByLaunchpadIdAndLaunchDate(booking)
	if err != nil {
		return true
	}

	if len(bookings) > 0 && bookings[0].DestinationID != booking.DestinationID {
		return true
	}

	return false
}
