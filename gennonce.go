package mazzaroth

import (
	"math/rand"
	"sync"
	"time"
)

var randNonce *rand.Rand

func seedRand() {
	randNonce = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateNonce() uint64 {
	var once sync.Once
	once.Do(seedRand)
	return randNonce.Uint64()
}
