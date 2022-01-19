package mazzaroth

import (
	"math/rand"
	"sync"
	"time"
)

func seedRand() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateNonce() uint64 {
	var once sync.Once
	once.Do(seedRand)
	return rand.Uint64()
}
