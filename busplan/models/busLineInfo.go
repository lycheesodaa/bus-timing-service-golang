package models

// for the sake of convenience since we don't have a database
var AvailableLineIDs = []string{"44478", "44479", "44480", "44481"}
var BusLineMap = make(map[string]BusLine)

// data from external API
type BusLine struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	RouteName string    `json:"routename"`
	Vehicles  []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	ID         int        `json:"vehicle_id"`
	Projection Projection `json:"projection"`
	Speed      float32    `json:"speed,string"`
	Statistics Statistics `json:"stats"`
	TimeStamp  string     `json:"ts"`
}

type Projection struct {
	EdgeID            int     `json:"edge_id"`
	EdgeDistance      float32 `json:"edge_distance,string"`
	EdgeProjection    float32 `json:"edge_projection,string"`
	StartNode         int     `json:"edge_start_node_id"`
	StopNode          int     `json:"edge_stop_node_id"`
	Latitude          float32 `json:"lat,string"`
	Longitude         float32 `json:"lon,string"`
	OriginalLatitude  float32 `json:"orig_lat,string"`
	OriginalLongitude float32 `json:"orig_lon,string"`
}

type Statistics struct {
	AvgSpeed float32 `json:"avg_speed,string"`
	Bearing  int     `json:"bearing"`
	Speed10  float32 `json:"cumm_speed_10,string"`
	Speed    int     `json:"speed"`
}
