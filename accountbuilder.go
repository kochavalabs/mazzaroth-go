package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type AccountBuilder struct {
	sender                *xdr.ID
	channel               *xdr.ID
	alias                 string
	nonce                 uint64
	blockExpirationNumber uint64
}

func (acctb *AccountBuilder) Account(sender, channel *xdr.ID, nonce, blockExpirationNumber uint64) *AccountBuilder {
	acctb.sender = sender
	acctb.channel = channel
	acctb.nonce = nonce
	acctb.blockExpirationNumber = blockExpirationNumber
	return acctb
}

func (acctb *AccountBuilder) Alias(alias string) *AccountBuilder {
	acctb.alias = alias
	return acctb
}

func (acctb *AccountBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	if acctb.alias == "" {
		return nil, errors.New("missing require fields")
	}

	data := xdr.Data{
		ChannelID:             *acctb.channel,
		Nonce:                 acctb.nonce,
		BlockExpirationNumber: acctb.blockExpirationNumber,
		Category: xdr.Category{
			Type: xdr.CategoryTypeACCOUNT,
			Account: &xdr.Account{
				Alias: acctb.alias,
			},
		},
	}
	dataBytes, err := data.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in data.MarshalBinary")
	}

	signer, err := xdr.IDFromPublicKey(pk.Public())
	if err != nil {
		return nil, errors.Wrap(err, "in xdr.IDFromPublicKey")
	}

	signatureSlice := ed25519.Sign(pk, dataBytes)

	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}

	transaction := &xdr.Transaction{
		Sender:    *acctb.sender,
		Signer:    signer,
		Signature: signature,
		Data:      data,
	}
	return transaction, nil
}
