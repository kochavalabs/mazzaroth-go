package mazzaroth

import (
	"crypto/ed25519"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

func TestCallBuilder(t *testing.T) {
	testChannel, _ := xdr.IDFromSlice([]byte("0000000000000000000000000000000000000000000000000000000000000000"))
	seedstr := "0000000000000000000000000000000000000000000000000000000000000000"

	seed, _ := hex.DecodeString(seedstr)
	privateKey := ed25519.NewKeyFromSeed(seed)
	testAddress, err := xdr.IDFromPublicKey(privateKey.Public())
	if err != nil {
		t.Fatal(err)
	}
	data := &xdr.Data{
		ChannelID:             testChannel,
		Nonce:                 0,
		BlockExpirationNumber: 1,
		Category: xdr.Category{
			Type: xdr.CategoryTypeCALL,
			Call: &xdr.Call{
				Function:  "test",
				Arguments: []xdr.Argument{"1"},
			},
		},
	}

	dataBytes, err := data.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	signatureSlice := ed25519.Sign(privateKey, dataBytes)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		t.Fatal(err)
	}

	wantTx := &xdr.Transaction{
		Sender:    testAddress,
		Signer:    testAddress,
		Signature: signature,
		Data:      data,
	}

	cb := new(CallBuilder)
	tx, err := cb.Call(&testAddress, &testChannel, 0, 1).
		Function("test").
		Arguments([]xdr.Argument{Int32(1)}...).Sign(privateKey)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantTx, tx) {
		t.Fatalf("expected: %v, got: %v", wantTx, tx)
	}
}
