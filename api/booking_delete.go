package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/application"
	"github.com/oismailov/launchpad-booking/pkg/util"
)

func DeleteBooking(c *gin.Context) {
	bookingUUID := c.Param("uuid")
	if !util.IsValidUUID(bookingUUID) {
		util.ResponseBadRequest(c, "Invalid booking UUID")
		return
	}

	err := application.DeleteBookingByUUID(bookingUUID)
	if err != nil {
		util.ResponseBadRequest(c, "Unable to delete a booking")
		return
	}

	util.ResponseOK(c, util.Message{Message: "Booking has been deleted"})
}
