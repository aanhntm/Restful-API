package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func randomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func randomString(n int) string {
	var sub strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sub.WriteByte(c)
	}

	return sub.String()
}

func RandomName() string {
	return randomString(10)
}

func RandomProductName() string {
	products := []string{"Cake", "Milk", "Banana", "Shrimp", "Pomelo", "Egg"}

	n := len(products)
	return products[rand.Intn(n)]
}

func RandomId() int32 {
	return randomInt(1, 100)
}

func RandomAmount() int32 {
	return randomInt(3, 450)
}
