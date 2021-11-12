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

var channelStr string
var channel xdr.ID
var address xdr.ID
var privateKey ed25519.PrivateKey
var client Client

func init() {
	var err error

	channelStr = "0000000000000000000000000000000000000000000000000000000000000000"
	channel, err = xdr.IDFromHexString(channelStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := "0000000000000000000000000000000000000000000000000000000000000000"
	seed, err := hex.DecodeString(publicKey)
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
	nonce := mathrand.Intn(100000000000000)

	ucb := UpdateConfigBuilder{}
	ucb.UpdateConfig(&address, &channel, uint64(nonce), blockHeight+10)
	transaction, err := ucb.
		Owner(&address).
		Sign(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the transaction.
	txResponse, err := client.TransactionSubmit(*transaction)
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
	ucb := UpdateContractBuilder{}

	txTransacion, err := ucb.UpdateContract(&address, &channel, uint64(nonce), blockHeight).
		Version(version).
		Abi(abi).
		Contract(contract).
		Sign(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the transaction.
	txResponse, err := client.TransactionSubmit(*txTransacion)
	if err != nil {
		log.Fatal(err)
	}

	transactionStr := hex.EncodeToString((*txResponse.TransactionID)[:])
	waitForReceipt(channelStr, transactionStr)
}

func TestTransactionSubmit(t *testing.T) {
	nonce := uint64(mathrand.Intn(100))
	blockHeightResp, err := client.ChannelHeight(hex.EncodeToString(channel[:]))
	require.NoError(t, err)
	blockExpirationNumber := blockHeightResp.Height.Height + 1

	// Submit.
	builder := CallBuilder{}
	transaction, err := builder.
		Call(&address, &channel, nonce, blockExpirationNumber).
		Function("args").
		Arguments([]xdr.Argument{String("a"), String("b"), String("c")}...).
		Sign(privateKey)
	require.NoError(t, err)

	txResponse, err := client.TransactionSubmit(*transaction)
	require.NoError(t, err)
	require.Equal(t, txResponse.Type, xdr.ResponseTypeTRANSACTIONID)

	transactionStr := hex.EncodeToString((*txResponse.TransactionID)[:])

	waitForReceipt(channelStr, transactionStr)

	// Transaction lookup.
	txLookupResponse, err := client.TransactionLookup(channelStr, transactionStr)
	require.NoError(t, err)
	require.Equal(t, txLookupResponse.Type, xdr.ResponseTypeTRANSACTION)

	// Block lookup.
	receiptLookupResponse, err := client.ReceiptLookup(channelStr, transactionStr)
	require.NoError(t, err)
	require.Equal(t, receiptLookupResponse.Type, xdr.ResponseTypeRECEIPT)

	// Block lookup by height.
	blockListResponse, err := client.BlockLookupByBlockHeight(channelStr, int(blockExpirationNumber-1))
	require.NoError(t, err)
	require.Equal(t, blockListResponse.Type, xdr.ResponseTypeBLOCKLIST)
	require.True(t, len(*blockListResponse.Blocks) > 0)

	blockID := hex.EncodeToString((*blockListResponse.Blocks)[0].Header.StateRoot[:])

	// Receipt lookup.
	blockResponse, err := client.BlockLookup(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, blockResponse.Type, xdr.ResponseTypeBLOCK)

	// Block lookup by id.
	blockListResponse, err = client.BlockLookupByBlockID(channelStr, blockID)
	require.NoError(t, err)
	require.Equal(t, blockListResponse.Type, xdr.ResponseTypeBLOCKLIST)
	require.True(t, len(*blockListResponse.Blocks) > 0)

	// // Receipt lookup by block height.
	// receiptLookupResponse, err = client.ReceiptLookupByBlockHeight(channelStr, 1)
	// require.NoError(t, err)
	// require.Equal(t, txLookupResponse.Type, xdr.ResponseTypeTRANSACTION)

	// // Receipt lookup by block id.
	// receiptLookupResponse, err = client.ReceiptLookupByBlockID(channelStr, blockID)
	// require.NoError(t, err)
	// require.Equal(t, txLookupResponse.Type, xdr.ResponseTypeTRANSACTION)

}

// func TestReceiptLookup(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000000"
// 	transactionID := "00000000000000000000000000000000"

// 	response, err := client.ReceiptLookup(channelID, transactionID)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeRECEIPT)
// }

// func TestReceiptLookupByBlockHeight(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000001"
// 	blockHeight := 10

// 	response, err := client.ReceiptLookupByBlockHeight(channelID, blockHeight)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeRECEIPTLIST)
// }

// func TestReceiptLookupByBlockID(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000001"
// 	blockID := "00000000000000000000000000000002"

// 	response, err := client.ReceiptLookupByBlockID(channelID, blockID)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeRECEIPTLIST)
// }

// func TestBlockLookup(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000000"
// 	transactionID := "00000000000000000000000000000000"

// 	response, err := client.BlockLookup(channelID, transactionID)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeBLOCK)
// }

// func TestBlockLookupByBlockHeight(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000001"
// 	blockHeight := 10

// 	response, err := client.BlockLookupByBlockHeight(channelID, blockHeight)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKLIST)
// }

// func TestBlockLookupByBlockID(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000001"
// 	blockID := "00000000000000000000000000000002"

// 	response, err := client.BlockLookupByBlockID(channelID, blockID)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKLIST)
// }

// func TestBlockHeaderLookup(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000000"
// 	transactionID := "00000000000000000000000000000000"

// 	response, err := client.BlockHeaderLookup(channelID, transactionID)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKHEADER)
// }

// func TestBlockHeaderLookupByBlockHeight(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000001"
// 	blockHeight := 10

// 	response, err := client.BlockHeaderLookupByBlockHeight(channelID, blockHeight)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKHEADERLIST)
// }

// func TestBlockHeaderLookupByBlockID(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000001"
// 	blockID := "00000000000000000000000000000002"

// 	response, err := client.BlockHeaderLookupByBlockID(channelID, blockID)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeBLOCKHEADERLIST)
// }

// func TestChannelLookup(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000000"

// 	response, err := client.ChannelLookup(channelID)
// 	require.NoError(t, err)
// 	require.Equal(t, response.Type, xdr.ResponseTypeCHANNEL)
// }

// func TestChannelHeight(t *testing.T) {
// 	client, err := NewMazzarothClient(servers)
// 	require.NoError(t, err)

// 	channelID := "00000000000000000000000000000000"

// 	response, err := client.ChannelHeight(channelID)
// 	require.NoError(t, err)
// 	t.Log(response)
// 	require.Equal(t, response.Type, xdr.ResponseTypeHEIGHT)
// }
