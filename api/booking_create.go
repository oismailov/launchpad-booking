package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/api/validation"
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/pkg/util"
)

func CreateBooking(c *gin.Context) {
	booking := model.Booking{UUID: util.GetUUID()}
	err := c.BindJSON(&booking)

	if err != nil {
		util.ResponseBadRequest(c, "Invalid request body")
		return
	}

	if err = validation.ValidateCreateBookingRequest(booking); err != nil {
		util.ResponseBadRequest(c, err.Error())
		return
	}

	util.ResponseCreated(c, booking)
}
