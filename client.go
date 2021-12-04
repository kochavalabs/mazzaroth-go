package mazzaroth

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

// Client defines a mazzaroth client that exposes common
// endpoints exposed by a mazzaroth readonly node.
type Client interface {
	AccountLookup(channelID string, accountID string) (*xdr.Account, error)
	AuthorizationLookup(channelID string, accountID string) (*xdr.Authorization, error)
	BlockHeaderLookup(channelID string, blockID string) (*xdr.BlockHeader, error)
	BlockHeaderList(channelID string, blockHeight int, number int) (*[]xdr.BlockHeader, error)
	BlockHeight(channelID string) (*xdr.BlockHeight, error)
	BlockLookup(channelID string, blockID string) (*xdr.Block, error)
	BlockList(channelID string, blockHeight int, number int) (*[]xdr.Block, error)
	ChannelAbi(channelID string) (*xdr.Abi, error)
	ChannelLookup(channelID string) (*xdr.Config, error)
	ReceiptLookup(channelID string, transactionID string) (*xdr.Receipt, error)
	TransactionLookup(channelID string, transactionID string) (*xdr.Transaction, error)
	TransactionSubmit(transaction *xdr.Transaction) (*xdr.ID, *xdr.Receipt, error)
}
