package warehouse

import (
	"io/ioutil"
	"net/http"

	httputils "github.com/Anil8753/truware/app/api/utils/http"
	"github.com/gorilla/mux"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	args, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.ccWarehouse.SubmitTransaction("UpdateAsset", id, string(args))
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
