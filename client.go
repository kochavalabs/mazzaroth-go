package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// Client defines a mazzaroth client that exposes common
// endpoints exposed by a mazzaroth readonly node.
type Client interface {
	TransactionSubmitCall(channelID string, seed string, functionName string, parameters []string, nonce uint64, blockExpirationNumber uint64) (*xdr.Response, error)
	TransactionSubmitContract(channelID string, seed string, contractBytes []byte, abiDef []byte, nonce uint64, blockExpirationNumber uint64) (*xdr.Response, error)
	TransactionSubmitConfig(channelID string, seed string, owner string, nonce uint64, blockExpirationNumber uint64) (*xdr.Response, error)
	TransactionUpdateAuthorization(channelID string, seed string, nonce uint64, blockExpirationNumber uint64, authorizedAddressStr string, alias string, authorizedAlias string, authorize bool) (*xdr.Response, error)
	TransactionLookup(channelID string, transactionID string) (*xdr.Response, error)
	TransactionLookupByBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	TransactionLookupByBlockID(channelID string, blockID string) (*xdr.Response, error)
	ReceiptLookup(channelID string, transactionID string) (*xdr.Response, error)
	BlockLookup(channelID string, blockID string) (*xdr.Response, error)
	BlockList(channelID string, blockHeight int, number int) (*xdr.Response, error)
	BlockHeaderLookup(channelID string, blockID string) (*xdr.Response, error)
	BlockHeaderList(channelID string, blockHeight int, number int) (*xdr.Response, error)
	ChannelLookup(channelID string) (*xdr.Response, error)
	BlockHeight(channelID string) (*xdr.Response, error)
	AccountLookup(channelID string, accountID string) (*xdr.Response, error)
}
