package mazzaroth

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

var servers = []string{"http://localhost:8081"}

func TestTransactionSubmit(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

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
			Type: 1,
			Call: &xdr.Call{
				Function:  "test",
				Arguments: []xdr.Argument{"1"},
			},
		},
	}
	actionStream, err := action.MarshalBinary()
	require.NoError(t, err)

	signatureSlice := ed25519.Sign(privateKey, actionStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	require.NoError(t, err)
	transaction := &xdr.Transaction{
		Signature: signature,
		Action:    action,
	}

	response, err := client.TransactionSubmit(*transaction)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeTRANSACTIONID)
}

func TestTransactionLookup(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000000"
	transactionID := "00000000000000000000000000000000"

	response, err := client.TransactionLookup(channelID, transactionID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeTRANSACTION)
}

func TestTransactionLookupByBlockHeight(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockHeight := 10

	response, err := client.BlockLookupByBlockHeight(channelID, blockHeight)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKLIST)
}

func TestTransactionLookupByBlockID(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockID := "00000000000000000000000000000002"

	response, err := client.BlockLookupByBlockID(channelID, blockID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKLIST)
}

func TestReceiptLookup(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000000"
	transactionID := "00000000000000000000000000000000"

	response, err := client.ReceiptLookup(channelID, transactionID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeRECEIPT)
}

func TestReceiptLookupByBlockHeight(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockHeight := 10

	response, err := client.ReceiptLookupByBlockHeight(channelID, blockHeight)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeRECEIPTLIST)
}

func TestReceiptLookupByBlockID(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockID := "00000000000000000000000000000002"

	response, err := client.ReceiptLookupByBlockID(channelID, blockID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeRECEIPTLIST)
}

func TestBlockLookup(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000000"
	transactionID := "00000000000000000000000000000000"

	response, err := client.BlockLookup(channelID, transactionID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCK)
}

func TestBlockLookupByBlockHeight(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockHeight := 10

	response, err := client.BlockLookupByBlockHeight(channelID, blockHeight)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKLIST)
}

func TestBlockLookupByBlockID(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockID := "00000000000000000000000000000002"

	response, err := client.BlockLookupByBlockID(channelID, blockID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKLIST)
}

func TestBlockHeaderLookup(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000000"
	transactionID := "00000000000000000000000000000000"

	response, err := client.BlockHeaderLookup(channelID, transactionID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKHEADER)
}

func TestBlockHeaderLookupByBlockHeight(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockHeight := 10

	response, err := client.BlockHeaderLookupByBlockHeight(channelID, blockHeight)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKHEADERLIST)
}

func TestBlockHeaderLookupByBlockID(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000001"
	blockID := "00000000000000000000000000000002"

	response, err := client.BlockHeaderLookupByBlockID(channelID, blockID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKHEADERLIST)
}

func TestChannelLookup(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000000"

	response, err := client.ChannelLookup(channelID)
	require.NoError(t, err)
	require.Equal(t, response.Type, xdr.ResponseTypeCHANNEL)
}

func TestChannelHeight(t *testing.T) {
	client, err := NewMazzarothClient(servers)
	require.NoError(t, err)

	channelID := "00000000000000000000000000000000"

	response, err := client.ChannelHeight(channelID)
	require.NoError(t, err)
	t.Log(response)
	require.Equal(t, response.Type, xdr.ResponseTypeHEIGHT)
}
