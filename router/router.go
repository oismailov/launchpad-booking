package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/api"
	"github.com/oismailov/launchpad-booking/config"
	"github.com/oismailov/launchpad-booking/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.JSONMiddleware())
	r.POST("/api/bookings", api.CreateBooking)
	r.GET("/api/bookings", api.ListBookings)

	return r
}

func GetPort() string {
	return fmt.Sprintf(":%d", config.Cfg.ServerConfigurations.Port)
}
