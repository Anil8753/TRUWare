package main

type Asset struct {
	Type    string           `json:"type"`
	Id      string           `json:"id"`
	OwnerId string           `json:"ownerId"`
	General GeneralInfo      `json:"generalInfo"`
	Postion WarehousePostion `json:"postion"`
	Status  WarehouseStatus  `json:"status"`
}

type WarehouseStatus int

const (
	Operational WarehouseStatus = iota
	NonOperational
)

type GeneralInfo struct {
	Name               string  `json:"name"`
	Phone              string  `json:"phone"`
	Email              string  `json:"email"`
	Address            string  `json:"address"`
	Details            string  `json:"details"`
	Rate               float64 `json:"rate"`
	PanalityAfterLimit float64 `json:"panalityAfterLimit"`
	PanalityPremature  float64 `json:"panalityPremature"`
	TotalArea          int     `json:"totalArea"`
	AllocatedArea      int     `json:"allocatedArea"`
}

type WarehousePostion struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
