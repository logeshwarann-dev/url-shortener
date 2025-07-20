package utils

import (
	"crypto/rand"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/deatil/go-encoding/base62"
)

func RemoveCharFromString(src string, target string, newChar string) string {
	return strings.ReplaceAll(src, target, newChar)
}

func CheckIfStringType(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.String
}

func IsStringEmpty(s string) bool {
	return len(s) == 0
}

func IntToStr(n int) string {
	return strconv.Itoa(n)
}

func StrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

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
