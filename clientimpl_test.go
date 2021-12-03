// +build integration

package mazzaroth

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
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
var authorizedAddress xdr.ID
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

	authorizedAddress, err = xdr.IDFromHexString(authorizedAddressStr)
	if err != nil {
		log.Fatal(err)
	}

	client, err = NewMazzarothClient(servers)
	if err != nil {
		log.Fatal(err)
	}

	blockHeightResp, err := client.BlockHeight(hex.EncodeToString(channel[:]))
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
	nonce := mathrand.Intn(100000000000000)

	ucb := ConfigBuilder{}
	ucb.Config(&address, &channel, uint64(nonce), blockHeight+10)
	transaction, err := ucb.
		Owner(&address).
		Sign(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the transaction.
	txResponse, err := client.TransactionSubmit(transaction)
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

	// Load the ABI.
	var abi xdr.Abi
	err = json.NewDecoder(bytes.NewReader(abiDef)).Decode(&abi)
	if err != nil {
		log.Fatal(err)
	}

	// Contract.
	contract, err := os.ReadFile(dataDir + "/contract.wasm")
	if err != nil {
		log.Fatal(err)
	}

	nonce := mathrand.Intn(100000000000000)

	// Create the transaction.
	ucb := ContractBuilder{}

	txTransacion, err := ucb.Contract(&address, &channel, uint64(nonce), blockHeight).
		Version(version).
		Abi(&abi).
		ContractBytes(contract).
		Sign(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the transaction.
	txResponse, err := client.TransactionSubmit(txTransacion)
	if err != nil {
		log.Fatal(err)
	}

	transactionStr := hex.EncodeToString((*txResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)
}

func TestIntegrationTest(t *testing.T) {
	blockHeightResp, err := client.BlockHeight(hex.EncodeToString(channel[:]))
	require.NoError(t, err)
	blockExpirationNumber := blockHeightResp.Height.Height

	channelResponse, err := client.ChannelLookup(channelStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeCONFIG, channelResponse.Type)

	var transactionStr string

	for i := 0; i < 5; i++ {
		// Submit.
		nonce := uint64(mathrand.Intn(100))
		blockExpirationNumber++

		transaction, err := Transaction(&address, &channel).
			Call(nonce, blockExpirationNumber).
			Function("args").
			Arguments([]xdr.Argument{String("a"), String("b"), String("c")}...).
			Sign(privateKey)
		require.NoError(t, err)

		txResponse, err := client.TransactionSubmit(transaction)

		require.NoError(t, err)
		require.Equal(t, txResponse.Type, xdr.ResponseTypeTRANSACTIONID)

		transactionStr = hex.EncodeToString((*txResponse.TransactionID)[:])

		waitForReceipt(channelStr, transactionStr)
	}

	// Transaction lookup.
	txLookupResponse, err := client.TransactionLookup(channelStr, transactionStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeTRANSACTION, txLookupResponse.Type)

	// Blocks list.
	blockListResponse, err := client.BlockList(channelStr, 2, 1)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKLIST, blockListResponse.Type)
	require.True(t, len(*blockListResponse.Blocks) > 0)

	blockID := hex.EncodeToString((*blockListResponse.Blocks)[0].Header.PreviousHeader[:])

	// Block lookup.
	blockResponse, err := client.BlockLookup(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCK, blockResponse.Type)

	// Receipt lookup.
	receiptLookupResponse, err := client.ReceiptLookup(channelStr, transactionStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeRECEIPT, receiptLookupResponse.Type)

	// Block header lookup.
	blockHeaderLookupResponse, err := client.BlockHeaderLookup(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKHEADER, blockHeaderLookupResponse.Type)

	// Block header list.
	blockHeaderLookupResponse, err = client.BlockHeaderList(channelStr, 2, 1)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeBLOCKHEADERLIST, blockHeaderLookupResponse.Type)

	// Channel lookup.
	channelLookupResponse, err := client.ChannelLookup(channelStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeCONFIG, channelLookupResponse.Type)

	// Authorize a key.
	nonce := uint64(mathrand.Intn(100))
	blockExpirationNumber++

	authBuilder := AuthorizationBuilder{}
	transaction, err := authBuilder.Authorization(&address, &channel, nonce, blockExpirationNumber).
		Account(&authorizedAddress).
		Authorize(true).
		Sign(privateKey)
	require.NoError(t, err)

	permissionResponse, err := client.TransactionSubmit(transaction)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeTRANSACTIONID, permissionResponse.Type)

	transactionStr = hex.EncodeToString((*permissionResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)

	// Account lookup, the permission is there.
	accountLookupResponse, err := client.AccountLookup(channelStr, seedStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeACCOUNT, accountLookupResponse.Type)

	// Unauthorize a key.
	nonce = uint64(mathrand.Intn(100))
	blockExpirationNumber++

	transaction, err = authBuilder.Authorization(&address, &channel, nonce, blockExpirationNumber).
		Account(&authorizedAddress).
		Authorize(false).
		Sign(privateKey)
	require.NoError(t, err)

	permissionResponse, err = client.TransactionSubmit(transaction)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeTRANSACTIONID, permissionResponse.Type)

	transactionStr = hex.EncodeToString((*permissionResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)

	// Account lookup, the permission is gone.
	accountLookupResponse, err = client.AccountLookup(channelStr, seedStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeACCOUNT, accountLookupResponse.Type)

	// Check channel abi.
	abiResponse, err := client.ChannelAbi(channelStr)
	require.NoError(t, err)
	require.Equal(t, xdr.ResponseTypeABI, abiResponse.Type)
}
