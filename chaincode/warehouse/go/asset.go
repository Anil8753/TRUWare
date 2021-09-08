package main

type WarehouseStatus int

const (
	Operational WarehouseStatus = iota
	NonOperational
)

type Asset struct {
	OwnerID    string
	ID         string
	Org        string
	Name       string
	Address    string
	TotalArea  int
	BookedArea int
	Rate       int
	Status     WarehouseStatus
	Bookings   []Booking
}

type Booking struct {
	ID       string
	BookerID string
	Area     int
	Duration int
}
