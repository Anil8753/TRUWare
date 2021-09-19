package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	httputils "github.com/Anil8753/truware/app/server/utils/http"
)

func (h *Handler) Wallet(w http.ResponseWriter, r *http.Request) {

	result, err := h.ccWarehouse.EvaluateTransaction("ReadOwnerWallet")
	if err != nil {
		fmt.Println(err)
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}

func (h *Handler) BuyTokens(w http.ResponseWriter, r *http.Request) {

	type BuyTokenParam struct {
		Amount string `json:"amount"`
		RefNo  string `json:"refNo"`
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	p := &BuyTokenParam{}
	if err := json.Unmarshal(bytes, &p); err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.ccWarehouse.SubmitTransaction("BuyTokens", p.Amount, p.RefNo)
	if err != nil {
		fmt.Println(err)
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
