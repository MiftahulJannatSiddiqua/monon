// package middleware

// import (
// 	"crypto/hmac"
// 	"crypto/sha256"

// 	//"ecommerce/config"
// 	"encoding/base64"
// 	"net/http"
// 	"strings"
// )

// func (m *Middlewares) AuthenticateJwt(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//frontend signature
// 		header := r.Header.Get("Authorization")
// 		if header == "" {
// 			http.Error(w, "Unathorized", http.StatusUnauthorized)
// 			return
// 		}

// 		headerArr := strings.Split(header, " ")
// 		if len(headerArr) != 2 {
// 			http.Error(w, "Unathorized", http.StatusUnauthorized)
// 			return
// 		}

// 		accessToken := headerArr[1]

// 		tokenParts := strings.Split(accessToken, ".")
// 		if len(tokenParts) != 3 {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		jwtHeader := tokenParts[0]
// 		jwtClaim := tokenParts[1]
// 		signature := tokenParts[2]

// 		//backend signature
// 		message := jwtHeader + "." + jwtClaim
// 		//cnf := config.GetConfig()
// 		byteArrSecret := []byte(m.cnf.JwtSecretKey)
// 		byteArrMessage := []byte(message)

// 		h := hmac.New(sha256.New, byteArrSecret)
// 		h.Write(byteArrMessage)

// 		hash := h.Sum(nil)
// 		newSignature := base64UrlEncode(hash)

// 		if newSignature != signature {
// 			http.Error(w, "Unauthorize. halar po tui hacker", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	},
// 	)
// }

// func base64UrlEncode(data []byte) string {
// 	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
// }


package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type ctxKey string

const userIDKey ctxKey = "userID"

func UserID(r *http.Request) (int, bool) {
	id, ok := r.Context().Value(userIDKey).(int)
	return id, ok
}

func (m *Middlewares) AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(headerArr[1], ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		message := tokenParts[0] + "." + tokenParts[1]

		h := hmac.New(sha256.New, []byte(m.cnf.JwtSecretKey))
		h.Write([]byte(message))
		expected := base64UrlEncode(h.Sum(nil))

		if !hmac.Equal([]byte(expected), []byte(tokenParts[2])) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claimBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(tokenParts[1])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var claims util.Payload
		if err := json.Unmarshal(claimBytes, &claims); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims.Exp <= time.Now().Unix() {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.Sub)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}