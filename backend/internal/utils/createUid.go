package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func CreateUid() (string, error) {
	max := big.NewInt(9000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	uid_int := n.Int64() + 1000
	uid := fmt.Sprintf("%d", uid_int)
	return uid, nil
}
