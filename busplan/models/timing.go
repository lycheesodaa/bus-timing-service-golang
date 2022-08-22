package models

// our own data
type Timing struct {
	ID    int                        `json:"bus_stop_id"`
	Name  string                     `json:"bus_stop_name"`
	Buses map[string][]BusTimingInfo `json:"buses_en_route"`
}

type BusTimingInfo struct {
	VehicleID  int    `json:"vehicle_id"`
	Vehicle    string `json:"vehicle"`
	ETASeconds int    `json:"eta_in_seconds"`
	ETAMinutes int    `json:"eta_in_minutes"`
	RouteInfo  Route  `json:"route_info"`
}
