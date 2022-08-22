package models

// our own data
type Location struct {
	ID    int               `json:"bus_line_ID"`
	Name  string            `json:"bus_line_name"`
	Buses []BusLocationInfo `json:"operating_buses"`
}

type BusLocationInfo struct {
	BusID       int         `json:"bus_id"`
	Speed       float32     `json:"bus_speed"`
	Coordinates Coordinates `json:"coordinates"`
	LastUpdated string      `json:"last_updated"`
}

type Coordinates struct {
	Bearing   int  `json:"bearing"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
