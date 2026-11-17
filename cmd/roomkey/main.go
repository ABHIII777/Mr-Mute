package roomkey

import (
	"crypto/rand"
	"math/big"
)

const alphabet = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

func GenerateRoomKey() (string, error) {
	const length = 8

	key := make([]byte, length)

	for i := range key {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}

		key[i] = alphabet[n.Int64()]
	}

	return string(key), nil
}
