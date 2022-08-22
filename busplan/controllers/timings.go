package controllers

import (
	"busplan/busplan/helpers"
	"busplan/busplan/models"

	"github.com/gin-gonic/gin"
)

// sorts incoming buses of a stop to their routes
func GetTimingByStopID(c *gin.Context) {
	// defer libs.RecoverError(context)
	var (
		status       = 200
		msg          string
		busStopID    = c.Param("id")
		responseData interface{}
	)

	if data, ok := models.BusStopMap[busStopID]; ok {
		msg = "Success"
		busToRouteMap := make(map[string][]models.BusTimingInfo)

		if buses := len(data.Forecast); buses > 0 {

			for i := 0; i < buses; i++ {
				index := data.Forecast[i].Route.RouteName

				if _, ok := busToRouteMap[index]; ok {
					busToRouteMap[index] = append(busToRouteMap[index], helpers.PopulateBusTimingInfo(data, i))
				} else {
					busToRouteMap[index] = []models.BusTimingInfo{helpers.PopulateBusTimingInfo(data, i)}
				}
			}
		}

		responseData = models.Timing{
			ID:    data.ID,
			Name:  data.Name,
			Buses: busToRouteMap,
		}

	} else {
		msg = "Error - Bus stop not found"
		status = 404
		responseData = gin.H{
			"status": status,
			"msg":    msg,
		}
	}
	helpers.APIResponseData(c, status, responseData)
}
