package main

import (
	"fmt"
	"net/http"

	"github.com/Anil8753/truware/app/api/warehouse"
	"github.com/gorilla/mux"
)

func main() {

	h := warehouse.GetHandler()

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./www")))
	router.HandleFunc("/api/warehouse", h.Create).Methods("POST")
	router.HandleFunc("/api/warehouse/{id}", h.Read).Methods("GET")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println(err)
	}

}
