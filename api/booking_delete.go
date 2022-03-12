package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/application"
	"github.com/oismailov/launchpad-booking/pkg/util"
)

func DeleteBooking(c *gin.Context) {
	bookingUUID := c.Param("uuid")
	if !util.IsValidUUID(bookingUUID) {
		util.ResponseBadRequest(c, "invalid booking UUID")
		return
	}

	err := application.DeleteBookingByUUID(bookingUUID)
	if err != nil {
		util.ResponseBadRequest(c, "unable to delete a booking")
		return
	}

	util.ResponseOK(c, util.Message{Message: "booking has been deleted"})
}
