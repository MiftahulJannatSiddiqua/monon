package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	prdCh := make(chan []*domain.ProductType)
	go func() {
		productList, err := h.svc.List(page, limit)
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		prdCh <- productList
	}()

	ch := make(chan int64)
	go func() {
		cnt, err := h.svc.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		ch <- cnt
	}()

	productList := <-prdCh
	totalCnt := <-ch

	util.SendPage(w, productList, page, limit, totalCnt)
}