package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random int between max and min
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates random string with given length
func RandomString(n int) string {
	k := len(alphabet)
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner create random name for owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance create random balance between 0 to 10000
func RandomBalance() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency create random with given list of currency
func RandomCurrency() string {
	currency := []string{"USD", "EUR", "RP"}
	length := len(currency)
	return currency[rand.Intn(length)]
}
