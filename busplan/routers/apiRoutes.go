package routers

import (
	"busplan/busplan/controllers"

	"github.com/gin-gonic/gin"
)

func APIRoutes(router *gin.RouterGroup) {
	router.GET("timing/:id", controllers.GetTimingByStopID)
	router.GET("location/:id", controllers.GetLocationByLineID)
	router.GET("location/all", controllers.GetAllBusLocations)
}