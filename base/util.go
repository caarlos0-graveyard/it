package base

import (
	"math/rand"
	"time"
)

// RandomStr creates a random string. Used to create the database name.
func RandomStr() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 10)
	for i := 0; i < 10; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
