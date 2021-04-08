package mazzarothtest

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"testing"

	v1 "mazzaroth-go/v1"
	xdrtypes "mazzaroth-go/v1/test/xdr"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

/*
BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error)
AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error)
NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error)
ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error)
*/

func TestTransactionSubmit(t *testing.T) {
	var client v1.Mazzaroth = v1.NewProductionClient(http.Client{})

	foo := xdrtypes.Foo{
		Status: 100,
		One:    "one",
		Two:    "two",
		Three:  "three",
	}

	fooBytes, err := json.Marshal(&foo)
	require.NoError(t, err)

	privateKeyBuffer := bytes.Repeat([]byte{0}, 32)
	privateKey := ed25519.NewKeyFromSeed(privateKeyBuffer)
	pubKey, ok := privateKey.Public().(ed25519.PublicKey)
	require.True(t, ok)

	var (
		address   xdr.ID
		channel   xdr.ID
		signature xdr.Signature
	)

	copy(address[:], pubKey)
	nonce := uint64(2000)
	call := xdr.Call{
		Function:   "insert_foo",
		Parameters: []xdr.Parameter{xdr.Parameter(fooBytes)},
	}

	action := *v1.BuildActionForTransactionCall(address, channel, nonce, call)
	xdrAction, err := action.MarshalBinary()
	require.NoError(t, err)

	copy(signature[:], ed25519.Sign(privateKey, xdrAction))

	transaction := xdr.Transaction{
		Signature: signature,
		Signer:    xdr.Authority{Type: xdr.AuthorityTypeNONE},
		Action:    action,
	}

	resp, err := client.TransactionSubmit(transaction)
	require.NoError(t, err)
	require.Equal(t, xdr.TransactionStatusACCEPTED, resp.Status)
	require.Equal(t, xdr.StatusInfo("Transaction has been accepted and is being executed."), resp.StatusInfo)
}

func TestReadonly(t *testing.T) {
	var client v1.Mazzaroth = v1.NewProductionClient(http.Client{})

	function := "simple"
	parameters := []xdr.Parameter{}
	resp, err := client.ReadOnly(function, parameters...)
	require.NoError(t, err)
	require.Equal(t, xdr.ReadonlyStatusSUCCESS, resp.Status)
	require.Equal(t, xdr.StatusInfo("Readonly request executed successfully."), resp.StatusInfo)
}

func TestTransactionLookup(t *testing.T) {
	var client v1.Mazzaroth = v1.NewProductionClient(http.Client{})

	var id xdr.ID

	resp, err := client.TransactionLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.TransactionStatusNOT_FOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("The transaction you looked up was not found."), resp.StatusInfo)
}

func TestReceiptLookup(t *testing.T) {
	var client v1.Mazzaroth = v1.NewProductionClient(http.Client{})

	var id xdr.ID

	resp, err := client.TransactionLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.TransactionStatusNOT_FOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("The transaction you looked up was not found."), resp.StatusInfo)
}

func TestBlockLookup(t *testing.T) {
	var client v1.Mazzaroth = v1.NewProductionClient(http.Client{})

	var number uint64
	var hash xdr.Hash

	number = 5
	copy(hash[:], bytes.Repeat([]byte{10}, 32))

	id := xdr.Identifier{
		Type:   xdr.IdentifierTypeHASH,
		Number: &number,
		Hash:   &hash,
	}

	resp, err := client.BlockLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.BlockStatusNOT_FOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("key not found in kv store"), resp.StatusInfo)
}

func TestBlockHeaderLookup(t *testing.T) {
	var client v1.Mazzaroth = v1.NewProductionClient(http.Client{})

	var number uint64
	var hash xdr.Hash

	number = 5
	copy(hash[:], bytes.Repeat([]byte{10}, 32))

	id := xdr.Identifier{
		Type:   xdr.IdentifierTypeHASH,
		Number: &number,
		Hash:   &hash,
	}

	resp, err := client.BlockHeaderLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.BlockStatusNOT_FOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("key not found in kv store"), resp.StatusInfo)
}

func TestAccountInfoLookup(t *testing.T) {
	var client v1.Mazzaroth = v1.NewProductionClient(http.Client{})

	var id xdr.ID

	ex, err := hex.DecodeString("dddd")
	require.NoError(t, err)

	copy(id[:], ex)

	resp, err := client.AccountInfoLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.InfoLookupStatusFOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("Found info for account."), resp.StatusInfo)
}