package main

import (
	"fmt"

	"github.com/Anil8753/truware/app/api/warehouse"
	"github.com/gorilla/mux"
)

func RegisterWarehouseRoutes(r *mux.Router) error {

	h, err := warehouse.GetHandler()
	if err != nil {
		return fmt.Errorf("failed to get handler. %v", err)
	}

	r.HandleFunc("/api/warehouse/identity", h.Identity).Methods("GET")

	r.HandleFunc("/api/wallet", h.Wallet).Methods("GET")

	r.HandleFunc("/api/warehouse", h.ReadAll).Methods("GET")
	r.HandleFunc("/api/warehouse", h.Create).Methods("POST")
	r.HandleFunc("/api/warehouse/{id}", h.Update).Methods("PUT")

	return nil
}
