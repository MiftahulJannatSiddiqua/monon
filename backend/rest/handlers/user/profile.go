package user

import (
	"ecommerce/rest/middleware"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"unicode/utf8"
)

type ReqUpdateProfile struct {
	Bio  *string `json:"bio"`
	City *string `json:"city"`
}

func (h Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	id, ok := middleware.UserID(r)
	if !ok {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	profile, err := h.svc.GetProfile(id)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if profile == nil {
		util.SendError(w, http.StatusNotFound, "Not found")
		return
	}

	util.SendData(w, http.StatusOK, profile)
}

func (h Handler) UpdateMe(w http.ResponseWriter, r *http.Request) {
	id, ok := middleware.UserID(r)
	if !ok {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req ReqUpdateProfile
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Bio == nil && req.City == nil {
		util.SendError(w, http.StatusBadRequest, "Nothing to update")
		return
	}
	if req.Bio != nil && utf8.RuneCountInString(*req.Bio) > 500 {
		util.SendError(w, http.StatusBadRequest, "Bio is too long")
		return
	}
	if req.City != nil && utf8.RuneCountInString(*req.City) > 100 {
		util.SendError(w, http.StatusBadRequest, "City is too long")
		return
	}

	if err := h.svc.UpdateProfile(id, req.Bio, req.City); err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	profile, err := h.svc.GetProfile(id)
	if err != nil || profile == nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusOK, profile)
}