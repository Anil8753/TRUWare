package main

type Order struct {
	Id                 string      `json:"id"`
	Type               string      `json:"type"`
	WarehouseId        string      `json:"warehouseId"`
	CustomerId         string      `json:"customerId"`
	Value              float64     `json:"value"`
	Rate               float64     `json:"rate"`
	Space              int         `json:"space"`
	Duration           int         `json:"duration"`
	PanalityAfterLimit float64     `json:"panalityAfterLimit"`
	PanalityPremature  float64     `json:"panalityPremature"`
	Status             OrderStatus `json:"status"`
	Conmments          string      `json:"conmments"`
}

type OrderStatus int

const (
	Placed OrderStatus = iota
	Approved
	Rejected
	Completed
	PrematureClosed
	AfterLimit
)
