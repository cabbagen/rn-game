package utils

import (
	"fmt"
	"time"
	"math/rand"
	"crypto/md5"
)


// 获取随机整数
func GetRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

// 是否存在字符串
func IncludeInStringList(list []string, s string) bool {
	var result bool = false
	for _, v := range list {
		if v == s {
			result = true
		}
	}
	return result
}

// 是否存在整型数字
func IncludeInIntList(list []int, t int) bool {
	var result bool = false
	for _, v := range list {
		if v == t {
			result = true
		}
	}
	return result
}

// MD5 加密
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}