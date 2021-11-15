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
			return
		}
		log.Println(response)
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
	require.Equal(t, channelResponse.Type, xdr.ResponseTypeCHANNEL)

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
	require.Equal(t, txLookupResponse.Type, xdr.ResponseTypeTRANSACTION)

	// Blocks from height.
	blockListResponse, err := client.BlockListFromBlockHeight(channelStr, 2)
	require.NoError(t, err)
	require.Equal(t, blockListResponse.Type, xdr.ResponseTypeBLOCKLIST)
	require.True(t, len(*blockListResponse.Blocks) > 0)

	blockID := hex.EncodeToString((*blockListResponse.Blocks)[1].Header.PreviousHeader[:])

	// Blocks from id.
	blockListResponse, err = client.BlockListFromBlockID(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, blockListResponse.Type, xdr.ResponseTypeBLOCKLIST)
	require.True(t, len(*blockListResponse.Blocks) > 0)

	// Block lookup.
	blockResponse, err := client.BlockLookup(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, blockResponse.Type, xdr.ResponseTypeBLOCK)

	// Receipt lookup.
	receiptLookupResponse, err := client.ReceiptLookup(channelStr, transactionStr)
	require.NoError(t, err)
	require.Equal(t, receiptLookupResponse.Type, xdr.ResponseTypeRECEIPT)

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
	require.Equal(t, blockHeaderLookupResponse.Type, xdr.ResponseTypeBLOCKHEADER)

	// Block headers from block height.
	blockHeaderLookupResponse, err = client.BlockHeaderListFromBlockHeight(channelStr, 2)
	require.NoError(t, err)
	require.Equal(t, blockHeaderLookupResponse.Type, xdr.ResponseTypeBLOCKHEADERLIST)

	// Block headers from block id.
	blockHeaderLookupResponse, err = client.BlockHeaderListFromBlockID(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, blockHeaderLookupResponse.Type, xdr.ResponseTypeBLOCKHEADERLIST)
}
