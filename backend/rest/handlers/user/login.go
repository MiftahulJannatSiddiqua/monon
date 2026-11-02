package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Request Data")
		return
	}

	usr, err := h.svc.Find(req.Email, req.Password) //database.Find(reqLogin.Email, reqLogin.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr == nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.SendData(w, http.StatusCreated, accessToken)
}