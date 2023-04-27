package main

import (
	"fmt"
	"safat99/task/repository"
	"safat99/task/routes"
	"safat99/task/service"
)

func main() {
	fmt.Println("hello")

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository)

	r := routes.SetupBookingRouter(bookingService)
	r.Run(":8080")

}
