package customer

import (
	"io/ioutil"
	"net/http"

	httputils "github.com/Anil8753/truware/app/server/utils/http"
)

func (h *Handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	args, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.ccOrder.SubmitTransaction("PlaceOrder", string(args))
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
