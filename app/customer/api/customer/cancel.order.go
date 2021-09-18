package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	httputils "github.com/Anil8753/truware/app/server/utils/http"
	"github.com/gorilla/mux"
)

func (h *Handler) CancelOrder(w http.ResponseWriter, r *http.Request) {

	fmt.Println("cancel order req")
	vars := mux.Vars(r)
	id := vars["id"]

	postdata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	type Comment struct {
		Comments string `json:"comments"`
	}

	cmt := Comment{}
	if err := json.Unmarshal(postdata, &cmt); err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.ccWarehouse.SubmitTransaction("CancelOrder", id, cmt.Comments)
	if err != nil {
		httputils.HttpResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("cancel order req DONEEEE")
	httputils.HttpResponse(w, string(result), http.StatusOK)
}
