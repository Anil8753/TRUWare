package main

import (
	"log"
	"net/http"
)

func main() {

	h := Handler{}

	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/api/query", h.Query)
	http.HandleFunc("/api/submit_txn", h.SubmitTransaction)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
