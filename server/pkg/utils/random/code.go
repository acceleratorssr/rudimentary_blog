package random

import (
	"math/rand"
	"time"
)

func Code(length int) string {
	//seed := time.Now().Unix() // 使用当前时间作为种子
	//src := rand.NewSource(seed)
	//r := rand.New(src) // 创建一个新的随机数生成器
	//
	//// 现在你可以使用r来生成随机数
	//num := r.Intn(100) // 生成一个[0,100)的随机数
	//println(num)
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成六位随机数
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}

	return string(result)
}
