package random

import (
	"math/rand"
	"time"
)

func GenerateSixDigit() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + rng.Intn(900000)
	return otp
}