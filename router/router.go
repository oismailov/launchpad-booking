package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oismailov/launchpad-booking/api"
	"github.com/oismailov/launchpad-booking/config"
	"github.com/oismailov/launchpad-booking/middleware"
	"github.com/oismailov/launchpad-booking/pkg/util"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.JSONMiddleware(), middleware.CORSMiddleware())
	r.POST("/api/bookings", api.CreateBooking)
	r.GET("/api/bookings", api.ListAllBookings)
	r.DELETE("/api/bookings/:uuid", api.DeleteBooking)
	r.GET("/api/swagger", api.Swagger)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, util.Message{Message: "Page not found"})
	})

	return r
}

func GetPort() string {
	return fmt.Sprintf(":%d", config.Cfg.ServerConfigurations.Port)
}
