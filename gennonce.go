package mazzaroth

import "math/rand"

func GenerateNonce() uint64 {
	return rand.Uint64()
}
