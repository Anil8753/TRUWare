package customer

import (
	"net/http"

	httputils "github.com/Anil8753/truware/app/server/utils/http"
)

func (h *Handler) ReadAllWarehouses(w http.ResponseWriter, r *http.Request) {

	result, err := h.ccOrder.EvaluateTransaction("ReadAllWarehouses")
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
