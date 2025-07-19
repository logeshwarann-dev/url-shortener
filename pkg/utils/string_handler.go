package utils

import (
	"crypto/rand"
	"log"

	"github.com/deatil/go-encoding/base62"
)

func GetShortCode(limit int) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, limit)
	rand.Read(bytes)
	result := make([]byte, limit)
	for i, b := range bytes {
		result[i] = chars[int(b)%62]
	}
	return string(result)
}

func EncodeString(src string) string {
	return base62.StdEncoding.EncodeToString([]byte(src))
}

func DecodeString(src string) (string, error) {
	orgUrl, err := base62.StdEncoding.DecodeString(src)
	if err != nil {
		log.Fatal("Error while decoding string [Source string: ", src, "] err: ", err.Error())
		return "", err
	}
	return string(orgUrl), nil
}
