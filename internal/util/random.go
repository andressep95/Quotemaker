package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnñopqrstuvwxyz"

// Generate a ramdon number beetwen min and max values
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
