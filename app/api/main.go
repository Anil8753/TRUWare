package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./www")))
	RegisterWarehouseRoutes(r)

	port := ":8081"
	fmt.Printf("\nStarting API server at port %s \n", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println(err)
	}
}
