package service

import (
	"crypto/sha1"
	b64 "encoding/base64"
)

func SHAHash(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)

	base64S := b64.StdEncoding.EncodeToString(bs)
	return base64S
}