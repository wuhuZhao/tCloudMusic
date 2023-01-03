package common

import (
	"crypto/md5"
	"encoding/hex"
)

// HashPassword: use md5 to encrypt password
func HashPassword(source string) string {
	d := []byte(source)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
