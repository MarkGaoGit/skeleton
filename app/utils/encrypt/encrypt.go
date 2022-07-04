package encrypt

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 md5字符串
// @des Sha256 使用的是另一种加密方式 也可以用当前的方式去加密
func MD5(iv string) string {
	hasSlice := md5.Sum([]byte(iv))
	return hex.EncodeToString(hasSlice[:])
}

// Sha256 Sha256加密
func Sha256(iv string) string {
	m := sha256.New()
	m.Write([]byte(iv))
	return hex.EncodeToString(m.Sum(nil))
}
