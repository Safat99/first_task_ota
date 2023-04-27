package routes

import (
	"safat99/task/service"

	"github.com/gin-gonic/gin"
)

func SetupBookingRouter(bookingService *service.BookingService) *gin.Engine {

	r := gin.Default()

	bookingRoutes := r.Group("/booking")
	{
		bookingRoutes.POST("/search", bookingService.Search)
		bookingRoutes.GET("/flightDetails/:id", bookingService.FlightDetails)
	}

	return r
}
