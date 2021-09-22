package main

type Order struct {
	Id     string      `json:"id"`
	Type   string      `json:"type"`
	Status OrderStatus `json:"status"`

	WarehouseId      string `json:"warehouseId"`
	WarehouseOwnerId string `json:"warehouseOwnerId"`
	WarehouseName    string `json:"warehouseName"`
	WarehouseContact string `json:"warehouseContact"`
	WarehouseGST     string `json:"warehouseGST"`

	CustomerId      string `json:"customerId"`
	CustomerName    string `json:"customerName"`
	CustomerContact string `json:"customerContact"`
	CustomerGST     string `json:"customerGST"`

	Value              float64 `json:"value"`
	Rate               float64 `json:"rate"`
	Space              int     `json:"space"`
	Date               string  `json:"date"`
	Duration           int     `json:"duration"`
	PanalityAfterLimit float64 `json:"panalityAfterLimit"`
	PanalityPremature  float64 `json:"panalityPremature"`
	Conmments          string  `json:"conmments"`
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
