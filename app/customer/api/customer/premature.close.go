package customer

import (
	"net/http"

	httputils "github.com/Anil8753/truware/app/server/utils/http"
	"github.com/gorilla/mux"
)

func (h *Handler) PrematureClose(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := h.ccWarehouse.SubmitTransaction("PrematureClose", id)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
