package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/pkg/svc"
	"github.com/oismailov/launchpad-booking/pkg/util"
	"net/http"
)

func DeleteBooking(c *gin.Context) {
	bookingUuid := c.Param("uuid")
	if !util.IsValidUUID(bookingUuid) {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: "Invalid booking UUID",
		})

		return
	}

	err := svc.DeleteBookingByUUID(bookingUuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: "Unable to delete a booking",
		})

		return
	}

	c.JSON(http.StatusOK, util.Message{
		Message: "Booking has been deleted",
	})
}
