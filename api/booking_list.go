package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/application"
	"github.com/oismailov/launchpad-booking/pkg/util"
)

func ListAllBookings(c *gin.Context) {
	bookings, err := application.GetAllBookings()
	if err != nil {
		util.ResponseBadRequest(c, err.Error())

		return
	}

	util.ResponseOK(c, bookings)
}
