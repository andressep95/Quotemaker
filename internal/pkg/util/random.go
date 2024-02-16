package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Generate a ramdon number beetwen min and max values
func RandomInt(min, max int) int {
	return min + rand.Intn(50)
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

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
