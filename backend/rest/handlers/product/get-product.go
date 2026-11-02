package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Invalid req body")
		return
	}

	product, err := h.svc.Get(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if product == nil {
		util.SendError(w, http.StatusNotFound, "product not found")
		return
	}

	util.SendData(w, http.StatusOK, product)
}