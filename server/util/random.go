package util

import (
	"math/rand"
	"time"
)

func GenarateRandomIntForJSON(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}
