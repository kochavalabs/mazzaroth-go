package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// Client defines a mazzaroth client that exposes common
// endpoints exposed by a mazzaroth readonly node.
type Client interface {
	TransactionSubmit(transaction *xdr.Transaction) (*xdr.Response, error)
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
