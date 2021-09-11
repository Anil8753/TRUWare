package main

type Asset struct {
	Id      string           `json:"id"`
	OwnerId string           `json:"ownerId"`
	General GeneralInfo      `json:"generalInfo"`
	Postion WarehousePostion `json:"postion"`
	Status  WarehouseStatus  `json:"status"`
	// Allocations []Allocation     `json:"allocations"`
}

type WarehouseStatus int

const (
	Operational WarehouseStatus = iota
	NonOperational
)

type GeneralInfo struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	Details       string `json:"details"`
	Rate          int    `json:"rate"`
	TotalArea     int    `json:"totalArea"`
	AllocatedArea int    `json:"allocatedArea"`
}

type WarehousePostion struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// type Allocation struct {
// 	Id       string `json:"name"`
// 	ClientId string `json:"clientId"`
// 	Area     int    `json:"area"`
// 	Duration int    `json:"duration"`
// }
