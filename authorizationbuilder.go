package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type AuthorizationBuilder struct {
	sender                *xdr.ID
	channel               *xdr.ID
	account               *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	key                   xdr.ID
	authorizedAccount     xdr.ID
	authorize             bool
}

func (ab *AuthorizationBuilder) Authorization(sender, channel *xdr.ID, nonce, blockExpirationNumber uint64) *AuthorizationBuilder {
	ab.sender = sender
	ab.channel = channel
	ab.nonce = nonce
	ab.blockExpirationNumber = blockExpirationNumber
	return ab
}

func (ab *AuthorizationBuilder) Account(account *xdr.ID) *AuthorizationBuilder {
	ab.account = account
	return ab
}

// Authorization - default to false
func (ab *AuthorizationBuilder) Authorize(authorize bool) *AuthorizationBuilder {
	ab.authorize = authorize
	return ab
}

func (ab *AuthorizationBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	if ab.account == nil {
		return nil, errors.New("missing account")
	}

	data := xdr.Data{
		ChannelID:             *ab.channel,
		Nonce:                 ab.nonce,
		BlockExpirationNumber: ab.blockExpirationNumber,
		Category: xdr.Category{
			Type: xdr.CategoryTypeAUTHORIZATION,
			Authorization: &xdr.Authorization{
				Account:   *ab.account,
				Authorize: ab.authorize,
			},
		},
	}

	dataBytes, err := data.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in action.MarshalBinary")
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
		Sender:    *ab.sender,
		Signer:    signer,
		Signature: signature,
		Data:      data,
	}

	return transaction, nil
}
