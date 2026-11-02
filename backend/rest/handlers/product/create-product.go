package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"imageUrl"`
}

func (h Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid json", 400)
		return
	}
	createdProduct, err := h.svc.Create(domain.ProductType{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImageUrl:    req.ImageUrl,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusCreated, createdProduct)
}