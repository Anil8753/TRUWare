package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ledger = Ledger{}

type Handler struct {
}

func (h Handler) Query(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	party := values.Get("party")
	if party == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ccnnot accept empty party name. either 'a' or 'b' is expected"))
		return
	}

	amount, err := ledger.Query(party)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"amount": %d }`, amount)))
}

func (h Handler) SubmitTransaction(w http.ResponseWriter, r *http.Request) {

	// define custom type
	type TxData struct {
		Sender   string `json:"sender"`
		Receiver string `json:"receiver"`
		Amount   int    `json:"amount"`
	}

	// define a var
	var postData TxData

	// decode input or return error
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("Sender: ", postData.Sender)
	fmt.Println("Receiver: ", postData.Receiver)
	fmt.Println("Amount: ", postData.Amount)

	err = ledger.Commit(postData.Sender, postData.Receiver, postData.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ "status" : "ok"}`))
}
