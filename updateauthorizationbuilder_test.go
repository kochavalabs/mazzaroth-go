package mazzaroth

import (
	"crypto/ed25519"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

func TestUpdateAuthorizationBuilder(t *testing.T) {
	authorizedAlias := "the authorized alias"
	authorized := true
	authorizedAddress, _ := xdr.IDFromSlice([]byte("00000000000000000000000000000001"))
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
				Type: xdr.UpdateTypeACCOUNT,
				Account: &xdr.AccountUpdate{
					Type: xdr.AccountUpdateTypeAUTHORIZATION,
					Authorization: &xdr.Authorization{
						Account: xdr.AuthorizedAccount{
							Key:   authorizedAddress,
							Alias: authorizedAlias,
						},
						Authorize: authorized,
					},
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
	ub := new(UpdateAuthorizationBuilder)
	tx, err := ub.UpdatePermission(&testAddress, &testChannel, 0, 1).
		Address(testAddress).
		Authorize(xdr.AccountUpdateTypeAUTHORIZATION, authorizedAddress, authorizedAlias, authorized).
		Sign(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, wantTx, tx)
	if !reflect.DeepEqual(wantTx, tx) {
		t.Fatalf("expected: %v, got: %v", wantTx, tx)
	}
}
