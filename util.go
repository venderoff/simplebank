package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(1))
	fmt.Println("New Value", rand.New(rand.NewSource(1)))
}

// Random Int Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min-1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i <= k; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Create Random Owner Name
func RandomOwner() string {
	return RandomString(6)
}

// Create Random AMount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Random Currency
func RandomCurrency() string {
	currencies := []string{"USD", "INR", "EUR", "CAD"}
	return currencies[RandomInt(0, 4)]

}
