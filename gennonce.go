package mazzaroth

import (
	"math/rand"
	"sync"
	"time"
)

var randNonce *rand.Rand
var once sync.Once

func seedRand() {
	randNonce = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateNonce() uint64 {
	once.Do(seedRand)
	return randNonce.Uint64()
}
