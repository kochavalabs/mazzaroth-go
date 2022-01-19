package mazzaroth

import "testing"

func TestGenNonce(t *testing.T) {
	v1 := GenerateNonce()
	v2 := GenerateNonce()

	if v1 == v2 {
		t.Fatalf("values v1: %d, v2: %d should not match", v1, v2)
	}
}
