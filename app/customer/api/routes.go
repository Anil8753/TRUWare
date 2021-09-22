package main

import (
	"fmt"

	"github.com/Anil8753/truware/app/server/customer"
	"github.com/gorilla/mux"
)

func RegisterWarehouseRoutes(r *mux.Router) error {

	h, err := customer.GetHandler()
	if err != nil {
		return fmt.Errorf("failed to get handler. %v", err)
	}

	r.HandleFunc("/api/identity", h.Identity).Methods("GET")
	r.HandleFunc("/api/registration", h.CreateRegistration).Methods("POST")
	r.HandleFunc("/api/registration", h.ReadRegistration).Methods("GET")
	r.HandleFunc("/api/registration", h.UpdateRegistration).Methods("PUT")

	r.HandleFunc("/api/wallet", h.Wallet).Methods("GET")
	r.HandleFunc("/api/wallet/buy", h.BuyTokens).Methods("POST")

	r.HandleFunc("/api/warehouse", h.ReadAllWarehouses).Methods("GET")

	r.HandleFunc("/api/orders", h.ReadAllOrders).Methods("GET")
	r.HandleFunc("/api/order", h.PlaceOrder).Methods("POST")
	r.HandleFunc("/api/order/cancel/{id}", h.CancelOrder).Methods("PUT")

	return nil
}
