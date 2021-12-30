package mazzaroth

import (
	"context"

	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

// Client defines a mazzaroth client that exposes common
// endpoints exposed by a mazzaroth readonly node.
type Client interface {
	BlockHeaderLookup(ctx context.Context, channelID string, blockID string) (*xdr.BlockHeader, error)
	BlockHeaderList(ctx context.Context, channelID string, blockHeight int, number int) (*[]xdr.BlockHeader, error)
	BlockHeight(ctx context.Context, channelID string) (*xdr.BlockHeight, error)
	BlockLookup(ctx context.Context, channelID string, blockID string) (*xdr.Block, error)
	BlockList(ctx context.Context, channelID string, blockHeight int, number int) (*[]xdr.Block, error)
	ChannelAbi(ctx context.Context, channelID string) (*xdr.Abi, error)
	ReceiptLookup(ctx context.Context, channelID string, transactionID string) (*xdr.Receipt, error)
	TransactionLookup(ctx context.Context, channelID string, transactionID string) (*xdr.Transaction, error)
	TransactionSubmit(ctx context.Context, transaction *xdr.Transaction) (*xdr.ID, *xdr.Receipt, error)
}
