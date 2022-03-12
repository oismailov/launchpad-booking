package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/pkg/svc"
	"github.com/oismailov/launchpad-booking/pkg/util"
	"net/http"
)

func ListBookings(c *gin.Context) {
	bookings, err := svc.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, bookings)
}
