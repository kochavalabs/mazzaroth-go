package mazzaroth

import (
	"crypto/ed25519"
	"net/http"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

var _ ClientWithPrivateKey = &ClientWithPrivateKeyImpl{}

// ClientWithPrivateKeyImpl : is the actual implementation of ClientWithPrivateKey interface
type ClientWithPrivateKeyImpl struct {
	client     *BareClientImpl
	privateKey ed25519.PrivateKey
}

// NewClientWithPrivateKey : builds a client storing your privateKey can be used to mock a real client
func NewClientWithPrivateKey(privateKey ed25519.PrivateKey, httpClient *http.Client, servers ...string) (*ClientWithPrivateKeyImpl, error) {
	client, err := NewBareClient(httpClient, servers...)
	if err != nil {
		return nil, err
	}
	return &ClientWithPrivateKeyImpl{privateKey: privateKey, client: client}, nil
}

// NewClientWithPrivateKeyWithDefaultHTTPClient : client for production with default http.Client
func NewClientWithPrivateKeyWithDefaultHTTPClient(privateKey ed25519.PrivateKey, servers ...string) (*ClientWithPrivateKeyImpl, error) {
	client, err := NewBareClientWithDefaultHTTPClient(servers...)
	if err != nil {
		return nil, err
	}
	return &ClientWithPrivateKeyImpl{privateKey: privateKey, client: client}, nil
}

// CallAction : makes a transaction with the given action call
func (cl *ClientWithPrivateKeyImpl) CallAction(action xdr.Action, authority *xdr.Authority) (*xdr.TransactionSubmitResponse, error) {
	actionStream, err := action.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in action.MarshalBinary")
	}
	signatureSlice := ed25519.Sign(cl.privateKey, actionStream)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}
	tx := xdr.Transaction{
		Action:    action,
		Signature: signature,
	}
	if authority == nil {
		tx.Signer = *authority
	} else {
		tx.Signer, err = xdr.NewAuthority(xdr.AuthorityTypeNONE, nil)
		if err != nil {
			return nil, errors.Wrap(err, "in xdr.NewAuthority(xdr.AuthorityTypeNONE)")
		}
	}
	return cl.client.TransactionSubmit(tx)
}

// ReadOnly calls the endpoint: /readonly.
func (cl *ClientWithPrivateKeyImpl) ReadOnly(function string, parameters ...xdr.Parameter) (*xdr.ReadonlyResponse, error) {
	return cl.client.ReadOnly(function, parameters...)
}

// TransactionLookup calls the endpoint: /transaction/lookup.
func (cl *ClientWithPrivateKeyImpl) TransactionLookup(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error) {
	return cl.client.TransactionLookup(transactionID)
}

// ReceiptLookup calls the endpoint: /receipt/lookup.
func (cl *ClientWithPrivateKeyImpl) ReceiptLookup(transactionID xdr.ID) (*xdr.ReceiptLookupResponse, error) {
	return cl.client.ReceiptLookup(transactionID)
}

// BlockLookup calls the endpoint: /block/lookup.
func (cl *ClientWithPrivateKeyImpl) BlockLookup(blockID xdr.Identifier) (*xdr.BlockLookupResponse, error) {
	return cl.client.BlockLookup(blockID)
}

// BlockHeaderLookup calls the endpoint: /block/header/lookup.
func (cl *ClientWithPrivateKeyImpl) BlockHeaderLookup(blockID xdr.Identifier) (*xdr.BlockHeaderLookupResponse, error) {
	return cl.client.BlockHeaderLookup(blockID)
}

// AccountInfoLookup calls the endpoint: /account/info/lookup.
func (cl *ClientWithPrivateKeyImpl) AccountInfoLookup(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error) {
	return cl.client.AccountInfoLookup(accountID)
}

// NonceLookup calls the endpoint: /account/nonce/lookup.
func (cl *ClientWithPrivateKeyImpl) NonceLookup(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error) {
	return cl.client.NonceLookup(accountID)
}

// ChannelInfoLookup calls the endpoint: /channel/info/lookup.
func (cl *ClientWithPrivateKeyImpl) ChannelInfoLookup(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error) {
	return cl.client.ChannelInfoLookup(channelInfoType)
}
