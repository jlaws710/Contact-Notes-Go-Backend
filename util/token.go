package util

import (
	"crypto/rand"
	"math/big"
	"sync"
)

var (
	tokens  = make(map[string]bool)
	mutex   sync.Mutex
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func GenerateToken() string {
	b := make([]rune, 32)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[num.Int64()]
	}

	return string(b)
}

func StoreToken(token string) {
	mutex.Lock()
	defer mutex.Unlock()
	tokens[token] = true
}

func IsValidToken(token string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	return tokens[token]
}
