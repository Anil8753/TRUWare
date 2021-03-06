package warehouse

import (
	"encoding/json"
	"net/http"

	httputils "github.com/Anil8753/truware/app/api/utils/http"
)

func (h *Handler) Identity(w http.ResponseWriter, r *http.Request) {

	type Identity struct {
		Name string
	}

	identity := &Identity{Name: "Warehouse Admin"}

	bytes, err := json.Marshal(identity)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusInternalServerError)
	}
	httputils.HttpResponse(w, string(bytes), http.StatusOK)
}
