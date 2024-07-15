// three-legged-crow.zero
//
// @(#)strings0.go  Friday, December 02, 2022
// Copyright(c) 2022, Leyton Goth. All rights reserved.

package strings0

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
	"unsafe"
)

// String 用于将 byte 数组转换为 string
//
//	 type stringHeader struct {
//		 Data unsafe.Pointer
//		 Len  int
//	 }
//	 type SliceHeader struct {
//		 Data uintptr
//		 Len  int
//		 Cap  int
//	 }
//
// 字节数组转为string （性能高于string(b)),编译器会完成内联处理，不会发生逃逸行为
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// MD5 字符串MD5编码
func MD5(s string) string {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 字符串SHA256编码
func SHA256(s string) string {
	hash := sha256.New()
	_, _ = hash.Write([]byte(s))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

// MySQLEscape escape sql for MySQL
func MySQLEscape(str string) string {
	str0 := strings.Replace(str, `\`, `\\`, -1)
	return strings.Replace(str0, "'", `\'`, -1)
}
