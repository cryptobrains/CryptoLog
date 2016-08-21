package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
	"math/big"
)

func CW(n int) string {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	result := base64.URLEncoding.EncodeToString(b)
	return result
}

func Makedir(path string) error {
	_, err := os.Open(path)
	if (err != nil) {
		err = os.MkdirAll(path, 0777)
	}
	return err
}

func StrToBigInt(str string) *big.Int {
	result := new(big.Int)
	result.SetString(str, 10)
	return result
}
