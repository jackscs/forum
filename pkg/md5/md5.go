package md5

import (
	"crypto/md5"
	"encoding/hex"
)

var secret = "scs"

// EncryptPassword 用于对传进来的密码进行加密
func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	opassword := h.Sum([]byte(password))
	return hex.EncodeToString(opassword)
}
