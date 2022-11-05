package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	letters := []byte("agasngsjsgehgsesfsgsegesasgaeasglghjhns")
	result := make([]byte, n)

	//给随机函数一个种子
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
