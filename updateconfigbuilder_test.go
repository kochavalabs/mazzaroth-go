package mazzaroth

import (
	"crypto/ed25519"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

func TestUpdateConfigBuilder(t *testing.T) {
	testAddress, _ := xdr.IDFromSlice([]byte("00000000000000000000000000000000"))
	testChannel, _ := xdr.IDFromSlice([]byte("00000000000000000000000000000000"))
	publicKey := "0000000000000000000000000000000000000000000000000000000000000000"
	seed, _ := hex.DecodeString(publicKey)
	privateKey := ed25519.NewKeyFromSeed(seed)
	action := xdr.Action{
		Address:               testAddress,
		ChannelID:             testChannel,
		Nonce:                 0,
		BlockExpirationNumber: 1,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeUPDATE,
			Update: &xdr.Update{
				Type: xdr.UpdateTypeCONFIG,
				ChannelConfig: &xdr.ChannelConfig{
					Owner:  testAddress,
					Admins: []xdr.ID{testAddress},
				},
			},
		},
	}
	actionStream, err := action.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	signatureSlice := ed25519.Sign(privateKey, actionStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		t.Fatal(err)
	}
	wantTx := &xdr.Transaction{
		Signature: signature,
		Action:    action,
	}
	ucb := new(UpdateConfigBuilder)
	tx, err := ucb.UpdateConfig(&testAddress, &testChannel, 0, 1).
		Owner(&testAddress).
		Admins(&testAddress).Sign(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(wantTx, tx) {
		t.Fatalf("expected: %v, got: %v", wantTx, tx)
	}
}
