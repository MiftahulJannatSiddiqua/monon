package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "please give me a valid product Id", 400)
		return
	}

	h.svc.Delete(pId)
	util.SendData(w, http.StatusOK, "successfully deleted")

}