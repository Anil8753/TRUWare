package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./www")))

	err := RegisterWarehouseRoutes(r)
	if err != nil {
		panic(err)
	}

	port := ":8081"
	fmt.Printf("\nStarting API server at port %s \n", port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200", "http://127.0.0.1:4200", "http://localhost:8081"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowCredentials: true,
		Debug:            false,
	})

	h := c.Handler(r)

	err = http.ListenAndServe(port, h)
	if err != nil {
		fmt.Println(err)
	}
}
