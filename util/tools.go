package util

import (
	"math/rand"
	"time"
)

const strRandTable = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GetRandomString len int 随机指定长度的字符串
func GetRandomString(num int) string {
	context := []byte(strRandTable)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, context[r.Intn(len(context))])
	}

	return string(result)
}

// GetRandomByte len int 随机指定长度的字符串
func GetRandomByte(num int) []byte {
	context := []byte(strRandTable)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, context[r.Intn(len(context))])
	}

	return result
}

//RandInt64 在两个数区间中随机
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}
