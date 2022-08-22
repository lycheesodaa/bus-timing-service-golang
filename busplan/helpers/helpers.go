package helpers

import (
	"busplan/busplan/models"

	"github.com/gin-gonic/gin"
)

func PopulateBusTimingInfo(data models.BusStop, index int) models.BusTimingInfo {
	return models.BusTimingInfo{
		VehicleID:  data.Forecast[index].VehicleID,
		Vehicle:    data.Forecast[index].Vehicle,
		ETASeconds: int(data.Forecast[index].Seconds),
		ETAMinutes: int(data.Forecast[index].Seconds) / 60,
		RouteInfo:  data.Forecast[index].Route,
	}
}

func PopulateBusLocationInfo(data models.BusLine, vehicle models.Vehicle) models.BusLocationInfo {
	return models.BusLocationInfo{
		BusID: vehicle.ID,
		Speed: vehicle.Speed,
		Coordinates: models.Coordinates{
			Bearing: vehicle.Statistics.Bearing,
			Latitude:  vehicle.Projection.Latitude,
			Longitude: vehicle.Projection.Longitude,
		},
		LastUpdated: vehicle.TimeStamp,
	}
}

func APIResponseData(c *gin.Context, status int, responseData interface{}) {
	responseType := c.Request.Header.Get("ResponseType")
	if responseType == "application/xml" {
		c.XML(status, responseData)
	} else {
		c.JSON(status, responseData)
	}
}
