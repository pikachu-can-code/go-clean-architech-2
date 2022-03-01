package encoder

import (
	"math/rand"
	"time"
)

func GenerateQR(size int) string {
	rand.Seed(time.Now().UnixNano())
	str := "IJKLMNOXYZ0123456789abcdefghijklPQRSTUVWmnopqrstuvwxyzABCDEFGH"

	min := 0
	max := len(str)

	code := ""
	for i := 0; i < size; i++ {
		idx := rand.Intn(max-min) + min
		code += string(str[idx])
	}

	return code
}
