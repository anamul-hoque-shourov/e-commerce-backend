package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/utils"
	"net/http"
	"strings"
)

func (middlewares *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "No Authorization header provided", http.StatusUnauthorized)
			return
		}
		headerArray := strings.Split(header, " ")
		if len(headerArray) != 2 {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		accessToken := headerArray[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArraySecret := []byte(middlewares.config.JwtSecret)
		byteArrayMessage := []byte(message)

		h := hmac.New(sha256.New, byteArraySecret)
		h.Write(byteArrayMessage)

		hash := h.Sum(nil)
		ExpectedSignature := utils.Base64UrlEncode(hash)

		if jwtSignature != ExpectedSignature {
			http.Error(w, "Invalid token signature. Tui hacker", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
