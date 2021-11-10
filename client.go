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
	ReceiptLookupByBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	ReceiptLookupByBlockID(channelID string, blockID string) (*xdr.Response, error)
	BlockLookup(channelID string, transactionID string) (*xdr.Response, error)
	BlockLookupByBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	BlockLookupByBlockID(channelID string, blockID string) (*xdr.Response, error)
	BlockHeaderLookup(channelID string, transactionID string) (*xdr.Response, error)
	BlockHeaderLookupByBlockHeight(channelID string, blockHeight int) (*xdr.Response, error)
	BlockHeaderLookupByBlockID(channelID string, blockID string) (*xdr.Response, error)
	ChannelLookup(channelID string) (*xdr.Response, error)
	ChannelHeight(channelID string) (*xdr.Response, error)
}
