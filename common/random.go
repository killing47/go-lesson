package common

import (
	"math/rand"
	"time"
)

func Random() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6)
	return s
}