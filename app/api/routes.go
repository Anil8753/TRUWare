package main

import (
	"github.com/Anil8753/truware/app/api/warehouse"
	"github.com/gorilla/mux"
)

func RegisterWarehouseRoutes(r *mux.Router) {

	h := warehouse.GetHandler()

	r.HandleFunc("/api/warehouse", h.Create).Methods("POST")
	r.HandleFunc("/api/warehouse/{id}", h.Read).Methods("GET")
}
