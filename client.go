package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// Client defines a mazzaroth client that exposes common
// endpoints exposed by a mazzaroth readonly node.
type Client interface {
	TransactionSubmit(transaction xdr.Transaction) (*xdr.Response, error)
	TransactionLookup(channelID string, transactionID string) (*xdr.Response, error)
	TransactionLookupByBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	TransactionLookupByBlockID(channelID string, blockID string) (*xdr.Response, error)
	ReceiptLookup(channelID string, transactionID string) (*xdr.Response, error)
	// ReceiptListFromBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	// ReceiptListFromBlockID(channelID string, blockID string) (*xdr.Response, error)
	BlockLookup(channelID string, transactionID string) (*xdr.Response, error)
	BlockListFromBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	BlockListFromBlockID(channelID string, blockID string) (*xdr.Response, error)
	BlockHeaderLookup(channelID string, transactionID string) (*xdr.Response, error)
	BlockHeaderListFromBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	BlockHeaderListFromBlockID(channelID string, blockID string) (*xdr.Response, error)
	ChannelLookup(channelID string) (*xdr.Response, error)
	ChannelHeight(channelID string) (*xdr.Response, error)
}
