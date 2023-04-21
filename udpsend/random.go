package main

import (
	"crypto/rand"
	"math/big"
	mathrandom "math/rand"
	"time"
)

var (
	character = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	chLen     = len(character)
)

// Random 随机数据
func Random(size int) []byte {
	buf := make([]byte, size)

	if _, err := rand.Read(buf); err != nil {
		mathrandom.Seed(time.Now().UnixNano())
		mathrandom.Read(buf)
	}
	return buf
}

// RandomString 随机字符串
func RandomString(size int) string {
	buf := make([]byte, size, size)
	max := big.NewInt(int64(chLen))
	for i := 0; i < size; i++ {
		random, err := rand.Int(rand.Reader, max)
		if err != nil {
			mathrandom.Seed(time.Now().UnixNano())
			buf[i] = character[mathrandom.Intn(chLen)]
			continue
		}
		buf[i] = character[random.Int64()]
	}
	return string(buf)
}

func RandomInt(min, max int) int {
	random, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		mathrandom.Seed(time.Now().UnixNano())
		return mathrandom.Intn(max-min) + min
	}
	return int(random.Int64()) + min
}
