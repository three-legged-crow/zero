// three-legged-crow.zero
//
// @(#)id.go  Monday, July 15, 2024
// Copyright(c) 2024, Leyton Goth. All rights reserved.

package strings0

import (
	"bytes"
	"math/rand"
	"time"
)

const (
	letterBytes    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytesLen = len(letterBytes)
	letterIdxBits  = 6                    // 6 bits to represent a letter index
	letterIdxMask  = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax   = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GenerateSessionID 生成 session id
func GenerateSessionID() (s string) {
	var bu bytes.Buffer
	bu.WriteString(RandStringBytesMask(3))
	bu.WriteByte('_')
	t := time.Now().Format("20060102_150405.000")
	bu.WriteString(t[:15])
	bu.WriteByte('_')
	bu.WriteString(t[16:])
	bu.WriteByte('_')
	bu.WriteString(RandStringBytesMask(8))
	s = String(bu.Bytes())
	return
}

// RandStringBytesMask 随机生成字符串
func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < letterBytesLen {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return String(b)
}
