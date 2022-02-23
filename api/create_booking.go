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
	isValidLaunchpad, err := isValidLaunchpad(booking)
	if err != nil || isValidLaunchpad == false {
		return false
	}

	isLaunchAvailable, err := isLaunchAvailable(booking)
	if err != nil || isLaunchAvailable == false {
		return false
	}

	hasDuplicatedDestination, err := hasDuplicatedDestination(booking)
	if err != nil || hasDuplicatedDestination == true {
		return false
	}

	return true
}

func isValidLaunchpad(booking model.Booking) (bool, error) {
	launchpad, err := svc.GetLaunchpadByID(booking.LaunchpadID)
	if err != nil {
		return false, err
	}

	if launchpad.ID == 0 {
		return false, nil
	}

	return true, nil
}

func isLaunchAvailable(booking model.Booking) (bool, error) {
	launchDate, err := time.Parse("2006-01-02", booking.LaunchDate)

	if err != nil {
		return false, err
	}

	launch, err := svc.GetLaunchByLaunchpadIDAndDate(booking.LaunchpadID, launchDate)
	if err != nil {
		return false, err
	}

	if launch.ID != 0 {
		return false, nil
	}

	return true, nil
}

func hasDuplicatedDestination(booking model.Booking) (bool, error) {
	result, err := svc.GetBookings(booking)
	if err != nil {
		return true, err
	}

	if len(result) > 0 && result[0].DestinationID != booking.DestinationID {
		return true, nil
	}

	return false, nil
}
