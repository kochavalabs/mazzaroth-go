package mazzaroth

import (
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

// BareClient is the Mazzaroth's client behaviour. A simple interface to avoid dealing with xdr
type BareClient interface {
	TransactionSubmit(transaction xdr.Transaction) (*xdr.TransactionSubmitResponse, error)
	ReadOnly(function string, parameters ...xdr.Parameter) (*xdr.ReadonlyResponse, error)
	TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error)
	ReceiptLookup(receiptID xdr.ID) (*xdr.ReceiptLookupResponse, error)
	BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error)
	BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error)
	AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error)
	NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error)
	ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error)
}

// ClientWithPrivateKey : a Mazzaroth's client that also stores the user's private key, helping the user with signing operations 
type ClientWithPrivateKey interface {
	CallAction(action xdr.Action, authority *xdr.Authority) (*xdr.TransactionSubmitResponse, error)
	ReadOnly(function string, parameters ...xdr.Parameter) (*xdr.ReadonlyResponse, error)
	TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error)
	ReceiptLookup(receiptID xdr.ID) (*xdr.ReceiptLookupResponse, error)
	BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error)
	BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error)
	AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error)
	NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error)
	ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error)
}
