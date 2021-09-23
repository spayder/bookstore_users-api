package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
