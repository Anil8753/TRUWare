package warehouse

import (
	"io/ioutil"
	"net/http"

	httputils "github.com/Anil8753/truware/app/api/utils/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	args, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.contract.SubmitTransaction("CreateAsset", string(args))
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.HttpResponse(w, string(result), http.StatusOK)
}
