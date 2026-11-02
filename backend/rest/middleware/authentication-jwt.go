package middleware

import (
	"crypto/hmac"
	"crypto/sha256"

	//"ecommerce/config"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//frontend signature
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtClaim := tokenParts[1]
		signature := tokenParts[2]

		//backend signature
		message := jwtHeader + "." + jwtClaim
		//cnf := config.GetConfig()
		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)

		if newSignature != signature {
			http.Error(w, "Unauthorize. halar po tui hacker", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	},
	)
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}