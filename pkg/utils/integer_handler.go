package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

func GenerateRandomNum(size int) int {
	var randStr string
	for i := 0; i < size; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(62))
		randStr = fmt.Sprintf("%v%v", randStr, num.Int64())
	}
	randNum, _ := strconv.Atoi(randStr)
	return randNum
}
