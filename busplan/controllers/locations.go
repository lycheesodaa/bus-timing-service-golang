package controllers

import (
	"busplan/busplan/helpers"
	"busplan/busplan/models"

	"github.com/gin-gonic/gin"
)

// gets bus locations by line id, reformatted from external location api into more readable format
func GetLocationByLineID(c *gin.Context) {
	// defer libs.RecoverError(context)
	var (
		status       = 200
		msg          string
		busLineID    = c.Param("id")
		responseData interface{}
	)

	if busLine, ok := models.BusLineMap[busLineID]; ok {
		list := make([]models.BusLocationInfo, len(busLine.Vehicles))

		for i, vehicle := range busLine.Vehicles {
			list[i] = helpers.PopulateBusLocationInfo(busLine, vehicle)
		}

		responseData = models.Location{
			ID:    busLine.ID,
			Name:  busLine.Name,
			Buses: list,
		}

	} else {
		msg = "Error - Bus line not found"
		status = 404
		responseData = gin.H{
			"status": status,
			"msg":    msg,
		}
	}

	helpers.APIResponseData(c, status, responseData)
}

// gets all bus locations, in a list
func GetAllBusLocations(c *gin.Context) {
	var (
		status       = 200
		msg          string
		responseData = gin.H{}
	)

	if len(models.BusLineMap) > 0 {
		lines := make([]models.Location, len(models.BusLineMap))
		index := 0
		for _, busLine := range models.BusLineMap {
			lines[index] = models.Location{
				ID:    busLine.ID,
				Name:  busLine.Name,
				Buses: make([]models.BusLocationInfo, len(busLine.Vehicles)),
			}
			for i, vehicle := range busLine.Vehicles {
				lines[index].Buses[i] = helpers.PopulateBusLocationInfo(busLine, vehicle)
			}
			index++
		}

		responseData = gin.H{
			"lines": lines,
		}

	} else {
		msg = "No bus lines found"
		status = 404
		responseData = gin.H{
			"status": status,
			"msg":    msg,
		}
	}

	helpers.APIResponseData(c, status, responseData)
}