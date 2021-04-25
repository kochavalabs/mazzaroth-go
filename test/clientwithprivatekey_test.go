// +build integration

package mazzarothtest

import (
	"bytes"
	"crypto/ed25519"
	"encoding/json"
	"net/http"
	"testing"

	v1 "github.com/kochavalabs/mazzaroth-go/v1"
	xdrtypes "github.com/kochavalabs/mazzaroth-go/v1/test/xdr"

	xdrlib "github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

func TestActionCall(t *testing.T) {
	privateKeyBuffer := bytes.Repeat([]byte{0}, ed25519.SeedSize)
	privateKey := ed25519.NewKeyFromSeed(privateKeyBuffer)

	client, err := v1.NewClientWithPrivateKey(privateKey, &http.Client{}, server)
	require.NoError(t, err)

	foo := xdrtypes.Foo{
		Status: 100,
		One:    "one",
		Two:    "two",
		Three:  "three",
	}

	fooBytes, err := json.Marshal(&foo)
	require.NoError(t, err)
	callParameters := []xdrlib.Parameter{xdrlib.Parameter(fooBytes)}
	call := xdrlib.Call{
		Function:   "insert_foo",
		Parameters: callParameters,
	}
	address, err := xdrlib.IDFromPublicKey(privateKey.Public())
	require.NoError(t, err)
	var channel xdrlib.ID
	nonce := uint64(2000)


	action, err := v1.BuildActionForTransactionCall(address, channel, nonce, call)
	require.NoError(t, err)
	require.NotNil(t, action)

	resp, err := client.CallAction(*action, nil)
	require.NoError(t, err)
	require.Equal(t, xdrlib.TransactionStatusACCEPTED, resp.Status)
	require.Equal(t, xdrlib.StatusInfo("Transaction has been accepted and is being executed."), resp.StatusInfo)
}
