package mazzaroth

import (
	"crypto/ed25519"
	"fmt"
	"reflect"
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

func TestCallBuilder(t *testing.T) {
	type test struct {
		testName     string
		address      xdr.ID
		channel      xdr.ID
		nonce        uint64
		functionName string
		fields       []Field
		pk           ed25519.PrivateKey
		want         *xdr.Transaction
	}

	tests := []test{}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, tc.testName), func(t *testing.T) {
			cb := new(CallBuilder)
			tx, err := cb.Call(tc.address, tc.channel, tc.nonce).
				Function(tc.functionName).
				Parameters(tc.fields...).Sign(tc.pk)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tc.want, tx) {
				t.Fatalf("expected: %v, got: %v", tc.want, tx)
			}
		})
	}
}
