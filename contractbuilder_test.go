package mazzaroth

import (
	"crypto/ed25519"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/kochavalabs/crypto"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

func TestContractBuilderDeploy(t *testing.T) {
	testChannel, _ := xdr.IDFromSlice([]byte("0000000000000000000000000000000000000000000000000000000000000000"))
	seedstr := "0000000000000000000000000000000000000000000000000000000000000000"
	seed, _ := hex.DecodeString(seedstr)
	privateKey := ed25519.NewKeyFromSeed(seed)

	testAddress, err := xdr.IDFromPublicKey(privateKey.Public())
	if err != nil {
		t.Fatal(err)
	}
	hasher := &crypto.Sha3_256Hasher{}
	hash := hasher.Hash([]byte("example"))
	xdrHash, err := xdr.HashFromSlice(hash)
	if err != nil {
		t.Fatal(err)
	}

	data := xdr.Data{
		ChannelID:             testChannel,
		Nonce:                 0,
		BlockExpirationNumber: 1,
		Category: xdr.Category{
			Type: xdr.CategoryTypeDEPLOY,
			Contract: &xdr.Contract{
				ContractBytes: []byte("example"),
				ContractHash:  xdrHash,
				Version:       "1",
				Abi:           xdr.Abi{Functions: []xdr.FunctionSignature{{FunctionType: xdr.FunctionTypeREAD, FunctionName: "Test"}}},
			},
		},
	}

	dataStream, err := data.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	signatureSlice := ed25519.Sign(privateKey, dataStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		t.Fatal(err)
	}

	wantTx := &xdr.Transaction{
		Sender:    testAddress,
		Signature: signature,
		Data:      data,
	}

	cb := new(ContractBuilder)
	tx, err := cb.Contract(&testAddress, &testChannel, 0, 1).
		Deploy("1", xdr.Abi{Functions: []xdr.FunctionSignature{{FunctionType: xdr.FunctionTypeREAD, FunctionName: "Test"}}}, []byte("example")).
		Sign(privateKey)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantTx, tx) {
		t.Fatalf("expected: %v, got: %v", wantTx, tx)
	}
}

func TestContractBuilderDelete(t *testing.T) {
	testChannel, _ := xdr.IDFromSlice([]byte("0000000000000000000000000000000000000000000000000000000000000000"))
	seedstr := "0000000000000000000000000000000000000000000000000000000000000000"
	seed, _ := hex.DecodeString(seedstr)
	privateKey := ed25519.NewKeyFromSeed(seed)

	testAddress, err := xdr.IDFromPublicKey(privateKey.Public())
	if err != nil {
		t.Fatal(err)
	}

	data := xdr.Data{
		ChannelID:             testChannel,
		Nonce:                 0,
		BlockExpirationNumber: 1,
		Category: xdr.Category{
			Type: xdr.CategoryTypeDELETE,
		},
	}

	dataStream, err := data.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	signatureSlice := ed25519.Sign(privateKey, dataStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		t.Fatal(err)
	}

	wantTx := &xdr.Transaction{
		Sender:    testAddress,
		Signature: signature,
		Data:      data,
	}

	cb := new(ContractBuilder)
	tx, err := cb.Contract(&testAddress, &testChannel, 0, 1).
		Delete().
		Sign(privateKey)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantTx, tx) {
		t.Fatalf("expected: %v, got: %v", wantTx, tx)
	}
}

func TestContractBuilderPause(t *testing.T) {
	testChannel, _ := xdr.IDFromSlice([]byte("0000000000000000000000000000000000000000000000000000000000000000"))
	seedstr := "0000000000000000000000000000000000000000000000000000000000000000"
	seed, _ := hex.DecodeString(seedstr)
	privateKey := ed25519.NewKeyFromSeed(seed)

	testAddress, err := xdr.IDFromPublicKey(privateKey.Public())
	if err != nil {
		t.Fatal(err)
	}
	pause := true

	data := xdr.Data{
		ChannelID:             testChannel,
		Nonce:                 0,
		BlockExpirationNumber: 1,
		Category: xdr.Category{
			Type:  xdr.CategoryTypePAUSE,
			Pause: &pause,
		},
	}

	dataStream, err := data.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	signatureSlice := ed25519.Sign(privateKey, dataStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		t.Fatal(err)
	}

	wantTx := &xdr.Transaction{
		Sender:    testAddress,
		Signature: signature,
		Data:      data,
	}

	cb := new(ContractBuilder)
	tx, err := cb.Contract(&testAddress, &testChannel, 0, 1).
		Pause(true).
		Sign(privateKey)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantTx, tx) {
		t.Fatalf("expected: %v, got: %v", wantTx, tx)
	}
}
