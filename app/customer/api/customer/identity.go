package customer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	httputils "github.com/Anil8753/truware/app/server/utils/http"
)

func (h *Handler) Identity(w http.ResponseWriter, r *http.Request) {

	type Identity struct {
		Name string
	}

	identity := &Identity{Name: "E-CommorceORG"}

	bytes, err := json.Marshal(identity)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.HttpResponse(w, string(bytes), http.StatusOK)
}

func (h *Handler) CreateRegistration(w http.ResponseWriter, r *http.Request) {

	args, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.ccWarehouse.SubmitTransaction("CreateRegistration", string(args))
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}

func (h *Handler) ReadRegistration(w http.ResponseWriter, r *http.Request) {

	result, err := h.ccWarehouse.EvaluateTransaction("ReadRegistration")
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}

func (h *Handler) UpdateRegistration(w http.ResponseWriter, r *http.Request) {

	args, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.ccWarehouse.SubmitTransaction("UpdateRegistration", string(args))
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
