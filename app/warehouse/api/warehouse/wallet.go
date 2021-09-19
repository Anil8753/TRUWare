package warehouse

import (
	"net/http"

	httputils "github.com/Anil8753/truware/app/api/utils/http"
)

func (h *Handler) Wallet(w http.ResponseWriter, r *http.Request) {

	result, err := h.ccWarehouse.EvaluateTransaction("ReadOwnerWallet")
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
