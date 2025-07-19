package utils

import (
	"log"

	"github.com/deatil/go-encoding/base62"
)

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
