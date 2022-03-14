package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/api/validation"
	"github.com/oismailov/launchpad-booking/application"
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/pkg/util"
)

func CreateBooking(c *gin.Context) {
	booking := model.Booking{UUID: util.GetUUID()}
	err := c.BindJSON(&booking)

	if err != nil {
		util.ResponseBadRequest(c, "invalid request body")
		return
	}

	if err = validation.ValidateCreateBookingRequest(booking); err != nil {
		util.ResponseBadRequest(c, err.Error())
		return
	}

	err = application.SaveBooking(&booking)
	if err != nil {
		util.ResponseBadRequest(c, "unable to save a booking")
		return
	}

	util.ResponseCreated(c, booking)
}
