package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()

	// First register '/api' routess
	err := RegisterWarehouseRoutes(r)
	if err != nil {
		panic(err)
	}

	// Second rest other paths serve from 'www' dir
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www")))

	port := ":8081"
	fmt.Printf("\nStarting API server at port %s \n", port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200",
			"http://127.0.0.1:4200",
			"http://localhost:8081",
			"http://127.0.0.1:8081",
		},
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

	srv := &http.Server{
		Handler: h,
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
