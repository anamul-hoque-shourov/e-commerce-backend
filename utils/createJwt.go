package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Id          int    `json:"id"`
	FistName    string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"isShopOwner"`
}

func CreateJwt(secret string, payload Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	encodedHeader := Base64UrlEncode(byteArrayHeader)

	byteArrayPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	encodedPayload := Base64UrlEncode(byteArrayPayload)

	message := encodedHeader + "." + encodedPayload

	byteArraySecret := []byte(secret)
	byteArrayMessage := []byte(message)

	h := hmac.New(sha256.New, byteArraySecret)
	h.Write(byteArrayMessage)
	signature := h.Sum(nil)
	encodedSignature := Base64UrlEncode(signature)

	jwt := encodedHeader + "." + encodedPayload + "." + encodedSignature
	return jwt, nil
}

func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
