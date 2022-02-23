package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/pkg/util"
	validator "gopkg.in/asaskevich/govalidator.v9"
	"net/http"
)

func CreateBooking(c *gin.Context) {
	var booking model.Booking
	err := c.BindJSON(&booking)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: err.Error(),
			//Message: "Invalid request body",
		})

		return
	}

	if _, err := validator.ValidateStruct(booking); err != nil {
		c.JSON(http.StatusBadRequest, util.Message{
			Message: err.Error(),
		})

		return
	}
}
