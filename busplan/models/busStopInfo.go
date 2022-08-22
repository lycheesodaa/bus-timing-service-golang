package models

// for the sake of convenience since we don't have a database
var AvailableStopIDs = []string{
	"378204", "383050", "378202", "383049", "382998", "378237", "378233", "378230",
	"378229", "378228", "378227", "382995", "378224", "378226", "383010", "383009",
	"383006", "383004", "378234", "383003", "378222", "383048", "378203", "382999",
	"378225", "383014", "383013", "383011", "377906", "383018", "383015", "378207",
}
var BusStopMap = make(map[string]BusStop)

// data from external API
type BusStop struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Forecast []Forecast `json:"forecast"`
	Geometry []Geometry `json:"geometry"`
}

type Forecast struct {
	VehicleID      int     `json:"vehicle_id"`
	Vehicle        string  `json:"vehicle"`
	Seconds        float32 `json:"forecast_seconds"`
	Route          Route   `json:"route"`
	RouteVariantID int     `json:"rv_id"`
	TotalPass      float32 `json:"total_pass"`
}

type Geometry struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

type Route struct {
	RouteID   int    `json:"id"`
	RouteName string `json:"short_name"`
}
