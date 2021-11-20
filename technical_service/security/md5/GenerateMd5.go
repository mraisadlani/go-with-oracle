package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptToMd5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	data := hex.EncodeToString(hasher.Sum(nil))

	return data
}
