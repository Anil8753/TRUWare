package main

type Asset struct {
	ID         string
	OwnerID    string
	Org        string
	Name       string
	Address    string
	TotalArea  int
	BookedArea int
	Rate       int

	Postion  WarehousePostion
	Status   WarehouseStatus
	Bookings []Booking
}

type WarehouseStatus int

const (
	Operational WarehouseStatus = iota
	NonOperational
)

type WarehousePostion struct {
	Latitude  float64
	Longitude float64
}

type Booking struct {
	ID       string
	BookerID string
	Area     int
	Duration int
}
