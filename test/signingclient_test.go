// +build integration

package mazzarothtest

import (
	"bytes"
	"crypto/ed25519"
	"encoding/json"
	"testing"

	"github.com/kochavalabs/mazzaroth-go"
	xdrtypes "github.com/kochavalabs/mazzaroth-go/test/xdr"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

func TestActionCall(t *testing.T) {
	privateKeyBuffer := bytes.Repeat([]byte{0}, ed25519.SeedSize)
	privateKey := ed25519.NewKeyFromSeed(privateKeyBuffer)

	client, err := mazzaroth.NewMazzarothSigningClient(privateKey, []string{server})
	require.NoError(t, err)

	foo := xdrtypes.Foo{
		Status: 100,
		One:    "one",
		Two:    "two",
		Three:  "three",
	}

	fooBytes, err := json.Marshal(&foo)
	require.NoError(t, err)
	callParameters := []xdr.Parameter{xdr.Parameter(fooBytes)}
	call := xdr.Call{
		Function:   "insert_foo",
		Parameters: callParameters,
	}
	address, err := xdr.IDFromPublicKey(privateKey.Public())
	require.NoError(t, err)
	var channel xdr.ID
	nonce := uint64(2000)

	cat, err := xdr.NewActionCategory(xdr.ActionCategoryTypeCALL, call)
	action := xdr.Action{
		Address:   address,
		ChannelID: channel,
		Nonce:     nonce,
		Category:  cat,
	}

	resp, err := client.CallAction(action, nil)
	require.NoError(t, err)
	require.Equal(t, xdr.TransactionStatusACCEPTED, resp.Status)
	require.Equal(t, xdr.StatusInfo("Transaction has been accepted and is being executed."), resp.StatusInfo)
}
