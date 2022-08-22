package services

import (
	"busplan/busplan/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
}

func CallTimings() {
	log.Info("Calling external timing api...")

	url := os.Getenv("TIMING_URL")

	// calling the external api for all bus stop IDs given
	for _, stopId := range models.AvailableStopIDs {
		currentURL := url + stopId + "/?format=json"

		// get response from url
		response, err := http.Get(currentURL)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Failed to get external api data")
			// os.Exit(1)
		}

		// read url data
		responseData, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("Failed to read external api data")
		}

		// parsing json response data into BusStop struct
		var busStopInfo models.BusStop
		json.Unmarshal(responseData, &busStopInfo)
		log.WithFields(log.Fields{
			"stopID":   busStopInfo.ID,
			"stopName": busStopInfo.Name,
		}).Debug("Loaded Bus Stop")

		// // creating a new struct to insert into the map, to be returned by the endpoints
		// // ? could possibly be optimised to only modify the related field
		// timing := models.Timing{
		// 	ID:          busStopInfo.ID,
		// 	Name:        busStopInfo.Name,
		// 	Buses:       []models.BusTimingInfo{},
		// 	LastUpdated: time.Now(),
		// }
		// for i := 0; i < len(busStopInfo.Forecast); i++ {
		// 	timing.Buses = append(timing.Buses, models.BusTimingInfo{
		// 		VehicleID:  busStopInfo.Forecast[i].VehicleID,
		// 		Vehicle:    busStopInfo.Forecast[i].Vehicle,
		// 		ETASeconds: int(busStopInfo.Forecast[i].Seconds),
		// 		ETAMinutes: int(busStopInfo.Forecast[i].Seconds) / 60,
		// 		RouteInfo:  busStopInfo.Forecast[i].Route,
		// 	})
		// }

		models.BusStopMap[stopId] = busStopInfo
	}

	log.Info("Successfully imported timing data")
}

func CallLocations() {
	log.Info("Calling external location api...")

	url := os.Getenv(("LOCATION_URL"))

	for _, lineId := range models.AvailableLineIDs {
		currentURL := url + lineId

		response, err := http.Get(currentURL)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Failed to get external api data")
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Fatal("Failed to read external api data")
		}

		var busLineInfo models.BusLine
		json.Unmarshal(responseData, &busLineInfo)
		log.WithFields(log.Fields{
			"ID":   busLineInfo.ID,
			"Name": busLineInfo.Name,
		}).Debug("Loaded Bus Line")

		// location := models.Location{
		// 	ID:          busLineInfo.ID,
		// 	Name:        busLineInfo.Name,
		// 	Buses:       []models.BusLocationInfo{},
		// 	LastUpdated: time.Now(),
		// }
		// for i := 0; i < len(busLineInfo.Vehicles); i++ {
		// 	location.Buses = append(location.Buses, models.BusLocationInfo{
		// 		BusID: busLineInfo.Vehicles[i].ID,
		// 		Speed: busLineInfo.Vehicles[i].Speed,
		// 		Coordinates: models.Coordinates{
		// 			Latitude:  busLineInfo.Vehicles[i].Projection.Latitude,
		// 			Longitude: busLineInfo.Vehicles[i].Projection.Longitude,
		// 		},
		// 	})
		// }

		models.BusLineMap[lineId] = busLineInfo
	}

	log.Info("Successfully imported location data")
}
