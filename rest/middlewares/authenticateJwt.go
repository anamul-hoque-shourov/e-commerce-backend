package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/utils"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("Authorization")
		if header == "" {
			http.Error(res, "No Authorization header provided", http.StatusUnauthorized)
			return
		}
		headerArray := strings.Split(header, " ")
		if len(headerArray) != 2 {
			http.Error(res, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		accessToken := headerArray[1]
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(res, "Invalid token format", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArraySecret := []byte(m.config.JwtSecret)
		byteArrayMessage := []byte(message)

		h := hmac.New(sha256.New, byteArraySecret)
		h.Write(byteArrayMessage)

		hash := h.Sum(nil)
		ExpectedSignature := utils.Base64UrlEncode(hash)

		if jwtSignature != ExpectedSignature {
			http.Error(res, "Invalid token signature. Tui hacker", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(res, req)
	})
}
