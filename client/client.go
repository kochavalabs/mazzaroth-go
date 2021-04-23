package client

import "github.com/kochavalabs/mazzaroth-xdr/xdr"

type MazzarothClient interface {
	TransactionSubmit(transaction xdr.Transaction) (*xdr.TransactionSubmitResponse, error)
	ReadOnly(function string, parameters ...xdr.Parameter) (*xdr.ReadonlyResponse, error)
	TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error)
	ReceiptLookup(transactionID xdr.ID) (*xdr.ReceiptLookupResponse, error)
	BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error)
	BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error)
	AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error)
	NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error)
	ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error)
}
