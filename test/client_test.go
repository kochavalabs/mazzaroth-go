// +build integration

package mazzarothtest

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/kochavalabs/mazzaroth-go"
	xdrtypes "github.com/kochavalabs/mazzaroth-go/test/xdr"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

const server = "http://localhost:8081"

/*
	How to run these tests:

		1. Run a Mazzaroth node.

			docker run -p 8081:8081 kochavalabs/mazzaroth start standalone

		2. Deploy the contract.

			cd $HOME/.go/src/github.com/Kochava/full-contract-example && mazzaroth-cli deploy deploy.json

		3. Go to the project's root and run the tests as usual.

			cd $HOME/.go/src/github.com/Kochava/mazzaroth-go && make integration
*/

// TestTransactionSubmit tests the happy path of the TransactionSubmit method.
func TestTransactionSubmit(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

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
		address xdr.ID
		channel xdr.ID
	)

	copy(address[:], pubKey)
	nonce := uint64(1)
	blockExpirationNumber := uint64(5)

	tx, err := mazzaroth.Transaction().Call(&address, &channel, nonce, blockExpirationNumber).
		Function("insert_foo").Arguments([]xdr.Argument{xdr.Argument(fooBytes)}...).Sign(privateKey)
	require.NoError(t, err)
	resp, err := client.TransactionSubmit(*tx)
	require.NoError(t, err)
	require.Equal(t, xdr.TransactionStatusACCEPTED, resp.Status)
	require.Equal(t, xdr.StatusInfo("Transaction has been accepted and is being executed."), resp.StatusInfo)
}

// TestTransactionLookup tests the happy path of the TestTransactionLookup method.
func TestTransactionLookup(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

	var id xdr.ID

	resp, err := client.TransactionLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.TransactionStatusNOT_FOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("The transaction you looked up was not found."), resp.StatusInfo)
}

// TestReceiptLookup tests the happy path of the TestReceiptLookup method.
func TestReceiptLookup(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

	var id xdr.ID

	resp, err := client.TransactionLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.TransactionStatusNOT_FOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("The transaction you looked up was not found."), resp.StatusInfo)
}

// TestBlockLookup tests the happy path of the TestBlockLookup method.
func TestBlockLookup(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

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

// TestBlockHeaderLookup tests the happy path of the TestBlockHeaderLookup method.
func TestBlockHeaderLookup(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

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

// TestAccountInfoLookup tests the happy path of the TestAccountInfoLookup method.
func TestAccountInfoLookup(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

	var id xdr.ID

	ex, err := hex.DecodeString("dddd")
	require.NoError(t, err)

	copy(id[:], ex)

	resp, err := client.AccountInfoLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.InfoLookupStatusFOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("Found info for account."), resp.StatusInfo)
}

// TestChannelInfoLookup tests the happy path of the TestChannelInfoLookup method.
func TestChannelInfoLookup(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

	id := xdr.ChannelInfoTypeCONTRACT

	resp, err := client.ChannelInfoLookup(id)
	require.NoError(t, err)
	require.Equal(t, xdr.InfoLookupStatusFOUND, resp.Status)
	require.Equal(t, xdr.StatusInfo("Found info for channel."), resp.StatusInfo)
	require.Equal(t, xdr.ChannelInfoTypeCONTRACT, resp.ChannelInfo.Type)
	require.NotNil(t, resp.ChannelInfo.Contract.ContractBytes)
	require.Equal(t, "0.1", resp.ChannelInfo.Contract.Version)
}

// TestBlockHeightLookup
func TestBlockHeightLookup(t *testing.T) {
	options := []mazzaroth.Options{}
	client, err := mazzaroth.NewMazzarothClient([]string{server}, options...)
	require.NoError(t, err)

	height, err := client.BlockHeightLookup()
	require.NoError(t, err)
	require.Equal(t, uint64(0), height)
}
