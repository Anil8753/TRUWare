package warehouse

import (
	"net/http"

	"github.com/gorilla/mux"

	httputils "github.com/Anil8753/truware/app/api/utils/http"
)

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := h.contract.EvaluateTransaction("ReadAsset", id)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
