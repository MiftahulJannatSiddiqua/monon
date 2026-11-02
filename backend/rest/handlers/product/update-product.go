package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}

func (h Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "please give me valid json")
		return
	}

	_, err = h.svc.Update(domain.ProductType{
		ID:       pId,
		Title:    req.Title,
		Price:    req.Price,
		ImageUrl: req.ImageUrl,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, http.StatusOK, "successfully updated product")
}