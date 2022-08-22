package main

import (
	"busplan/busplan/routers"
	"busplan/busplan/services"
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func getPort() string {
	p := os.Getenv("HOST_PORT")
	if p != "" {
		return ":" + p
	}
	return ":8888"
}

func main() {
	// initialising data for the first time
	services.CallTimings()
	services.CallLocations()

	// create cron job for external api scheduling
	c := cron.New()
	c.AddFunc("@every 30s", services.CallTimings)
	c.AddFunc("@every 30s", services.CallLocations)
	c.Start()

	port := getPort()
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(location.Default())

	routerGroup := router.Group("api")
	{
		routers.APIRoutes(routerGroup) // routers != router
	}
	router.Run(port)
}
