package customer

import (
	"net/http"

	httputils "github.com/Anil8753/truware/app/server/utils/http"
)

func (h *Handler) ReadAllOrders(w http.ResponseWriter, r *http.Request) {

	result, err := h.ccWarehouse.EvaluateTransaction("ReadAllOrders")
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
