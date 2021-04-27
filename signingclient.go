package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

var _ SigningClient = &SigningClientImpl{}

// SigningClientImpl : is the actual implementation of ClientWithPrivateKey interface
type SigningClientImpl struct {
	client     *ClientImpl
	privateKey ed25519.PrivateKey
}

// NewMazzarothSigningClient : builds a client storing your privateKey, accpts a custom-configured http.Client
func NewMazzarothSigningClient(privateKey ed25519.PrivateKey, servers []string, options ...Options) (*SigningClientImpl, error) {
	client, err := NewMazzarothClient(servers, options...)
	if err != nil {
		return nil, err
	}
	return &SigningClientImpl{privateKey: privateKey, client: client}, nil
}

// CallAction : makes a transaction with the given action call
func (cl *SigningClientImpl) CallAction(action xdr.Action, authority *xdr.Authority) (*xdr.TransactionSubmitResponse, error) {
    txBuilder := NewTransactionBuilder().WithAction(action)
	if authority != nil {
		txBuilder = txBuilder.WithAuthority(*authority)
	}
    tx, err := txBuilder.Sign(cl.privateKey)
	if err != nil {
	    return nil, errors.Wrap(err, "in signing the transaction")

	}
	return cl.client.TransactionSubmit(*tx)
}

// ReadOnly calls the endpoint: /readonly.
func (cl *SigningClientImpl) ReadOnly(function string, parameters ...xdr.Parameter) (*xdr.ReadonlyResponse, error) {
	return cl.client.ReadOnly(function, parameters...)
}

// TransactionLookup calls the endpoint: /transaction/lookup.
func (cl *SigningClientImpl) TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error) {
	return cl.client.TransactionLookup(transactionID)
}

// ReceiptLookup calls the endpoint: /receipt/lookup.
func (cl *SigningClientImpl) ReceiptLookup(transactionID xdr.ID) (*xdr.ReceiptLookupResponse, error) {
	return cl.client.ReceiptLookup(transactionID)
}

// BlockLookup calls the endpoint: /block/lookup.
func (cl *SigningClientImpl) BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error) {
	return cl.client.BlockLookup(blockID)
}

// BlockHeaderLookup calls the endpoint: /block/header/lookup.
func (cl *SigningClientImpl) BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error) {
	return cl.client.BlockHeaderLookup(blockID)
}

// AccountInfoLookup calls the endpoint: /account/info/lookup.
func (cl *SigningClientImpl) AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error) {
	return cl.client.AccountInfoLookup(accountID)
}

// NonceLookup calls the endpoint: /account/nonce/lookup.
func (cl *SigningClientImpl) NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error) {
	return cl.client.NonceLookup(accountID)
}

// ChannelInfoLookup calls the endpoint: /channel/info/lookup.
func (cl *SigningClientImpl) ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error) {
	return cl.client.ChannelInfoLookup(channelInfoType)
}
