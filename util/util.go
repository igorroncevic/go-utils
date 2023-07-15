package util

import (
	"crypto/rand"
	"math/big"
)

func Ptr[T any](val T) *T {
	return &val
}

func RandomInt64(limit int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(limit))
	if err != nil {
		return -1, err
	}

	return n.Int64(), nil
}
