package utility

import (
	"math/rand"
	"time"
)

func RandomId() string {
	rand.Seed(time.Now().UTC().UnixNano())

	random := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

	r := make([]rune, 11)
	for i := range r {
		r[i] = random[rand.Intn(len(random))]
	}

	return string(r)
}
