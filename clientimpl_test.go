// +build integration

package mazzaroth

import (
	"crypto/ed25519"
	"encoding/hex"
	"log"
	mathrand "math/rand"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/stretchr/testify/require"
)

var servers = []string{"http://localhost:8081"}

var channelStr string = "0000000000000000000000000000000000000000000000000000000000000000"
var seedStr string = "0000000000000000000000000000000000000000000000000000000000000000"
var addressStr string = "0000000000000000000000000000000000000000000000000000000000000000"
var authorizedAddressStr string = "0000000000000000000000000000000000000000000000000000000000000001"

var channel xdr.ID
var address xdr.ID
var privateKey ed25519.PrivateKey
var client Client

func init() {
	var err error

	channel, err = xdr.IDFromHexString(channelStr)
	if err != nil {
		log.Fatal(err)
	}

	seed, err := hex.DecodeString(seedStr)
	if err != nil {
		log.Fatal(err)
	}

	privateKey = ed25519.NewKeyFromSeed(seed)
	address, err = xdr.IDFromPublicKey(privateKey.Public())
	if err != nil {
		log.Fatal(err)
	}

	client, err = NewMazzarothClient(servers)
	if err != nil {
		log.Fatal(err)
	}

	blockHeightResp, err := client.ChannelHeight(hex.EncodeToString(channel[:]))
	if err != nil {
		log.Fatal(err)
	}
	blockHeight := blockHeightResp.Height.Height

	// Owner.
	blockHeight++
	setOwner(blockHeight)

	// Contract.
	blockHeight++
	uploadContract(blockHeight)
}

func getTestDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func waitForReceipt(channelStr string, transactionIDstr string) {
	for i := 0; i < 10; i++ {
		response, err := client.ReceiptLookup(channelStr, transactionIDstr)
		if err == nil {
			log.Println("receipt -> ", response.Receipt)
			return
		}
		// log.Println(response)
		time.Sleep(1 * time.Second)
	}

	log.Fatal("Transaction not found")
}

func setOwner(blockHeight uint64) {
	nonce := uint64(mathrand.Intn(100000000000000))

	// Execute the transaction.
	txResponse, err := client.TransactionSubmitConfig(channelStr, seedStr, addressStr, nonce, blockHeight)
	if err != nil {
		log.Fatal(err)
	}

	transactionStr := hex.EncodeToString((*txResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)
}

func uploadContract(blockHeight uint64) {
	dataDir := getTestDir() + "/testdata"

	abiDef, err := os.ReadFile(dataDir + "/ExampleContract.json")
	if err != nil {
		log.Fatal(err)
	}

	// Contract.
	contract, err := os.ReadFile(dataDir + "/contract.wasm")
	if err != nil {
		log.Fatal(err)
	}

	nonce := uint64(mathrand.Intn(100000000000000))

	// Execute the transaction.
	txResponse, err := client.TransactionSubmitContract(channelStr, seedStr, contract, abiDef, nonce, blockHeight)
	if err != nil {
		log.Fatal(err)
	}

	transactionStr := hex.EncodeToString((*txResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)
}

func TestIntegrationTest(t *testing.T) {
	blockHeightResp, err := client.ChannelHeight(hex.EncodeToString(channel[:]))
	require.NoError(t, err)
	blockExpirationNumber := blockHeightResp.Height.Height

	channelResponse, err := client.ChannelLookup(channelStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeCHANNEL, channelResponse.Type)

	var transactionStr string

	for i := 0; i < 5; i++ {
		// Submit.
		nonce := uint64(mathrand.Intn(100))
		blockExpirationNumber++

		txResponse, err := client.TransactionSubmitCall(channelStr, seedStr, "args", []string{"a", "b", "c"}, nonce, blockExpirationNumber)
		require.NoError(t, err)
		require.Equal(t, txResponse.Type, xdr.ResponseTypeTRANSACTIONID)

		transactionStr = hex.EncodeToString((*txResponse.TransactionID)[:])

		waitForReceipt(channelStr, transactionStr)
	}

	// Transaction lookup.
	txLookupResponse, err := client.TransactionLookup(channelStr, transactionStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeTRANSACTION, txLookupResponse.Type)

	// Blocks from height.
	blockListResponse, err := client.BlockListFromBlockHeight(channelStr, 2)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKLIST, blockListResponse.Type)
	require.True(t, len(*blockListResponse.Blocks) > 0)

	blockID := hex.EncodeToString((*blockListResponse.Blocks)[1].Header.PreviousHeader[:])

	// Blocks from id.
	blockListResponse, err = client.BlockListFromBlockID(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKLIST, blockListResponse.Type)
	require.True(t, len(*blockListResponse.Blocks) > 0)

	// Block lookup.
	blockResponse, err := client.BlockLookup(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCK, blockResponse.Type)

	// Receipt lookup.
	receiptLookupResponse, err := client.ReceiptLookup(channelStr, transactionStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeRECEIPT, receiptLookupResponse.Type)

	// // Receipt lookup by height.
	// receiptListLookupResponse, err := client.ReceiptListFromBlockHeight(channelStr, 2)
	// require.NoError(t, err)
	// require.Equal(t, receiptListLookupResponse.Type, xdr.ResponseTypeRECEIPTLIST)

	// // Receipt lookup by block id.
	// receiptListLookupResponse, err := client.ReceiptListFromBlockID(channelStr, blockID)
	// require.NoError(t, err)
	// require.Equal(t, receiptListLookupResponse.Type, xdr.ResponseTypeRECEIPTLIST)

	// Block header lookup.
	blockHeaderLookupResponse, err := client.BlockHeaderLookup(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKHEADER, blockHeaderLookupResponse.Type)

	// Block headers from block height.
	blockHeaderLookupResponse, err = client.BlockHeaderListFromBlockHeight(channelStr, 2)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKHEADERLIST, blockHeaderLookupResponse.Type)

	// Block headers from block id.
	blockHeaderLookupResponse, err = client.BlockHeaderListFromBlockID(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKHEADERLIST, blockHeaderLookupResponse.Type)

	// Channel lookup.
	channelLookupResponse, err := client.ChannelLookup(channelStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeCHANNEL, channelLookupResponse.Type)

	// Authorize a key.
	nonce := uint64(mathrand.Intn(100))
	blockExpirationNumber++
	alias := "the alias"
	authorizedAlias := "the authorized alias"
	permissionResponse, err := client.TransactionUpdatePermission(channelStr, seedStr, nonce, blockExpirationNumber, authorizedAddressStr, alias, authorizedAlias, true)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeTRANSACTIONID, permissionResponse.Type)

	transactionStr = hex.EncodeToString((*permissionResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)

	// Account lookup, the permission is there.
	accountLookupResponse, err := client.AccountLookup(channelStr, seedStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeACCOUNT, accountLookupResponse.Type)
	require.Equal(t, "0000000000000000000000000000000000000000000000000000000000000001", hex.EncodeToString(accountLookupResponse.Account.AuthorizedAccounts[0].Key[:]))
	require.Equal(t, "the authorized alias", accountLookupResponse.Account.AuthorizedAccounts[0].Alias)

	// Unauthorize a key.
	nonce = uint64(mathrand.Intn(100))
	blockExpirationNumber++
	alias = "the alias"
	authorizedAlias = "the authorized alias"
	permissionResponse, err = client.TransactionUpdatePermission(channelStr, seedStr, nonce, blockExpirationNumber, authorizedAddressStr, alias, authorizedAlias, false)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeTRANSACTIONID, permissionResponse.Type)

	transactionStr = hex.EncodeToString((*permissionResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)

	// Account lookup, the permission is gone.
	accountLookupResponse, err = client.AccountLookup(channelStr, seedStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeACCOUNT, accountLookupResponse.Type)
	require.Equal(t, 0, len(accountLookupResponse.Account.AuthorizedAccounts))
}
