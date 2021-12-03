package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type UpdateAuthorizationBuilder struct {
	address, channel      *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	signer                *xdr.Authority
	key                   xdr.ID
	authorizedAccount     xdr.ID
	authorizedAlias       string
	authorize             bool
}

func (upb *UpdateAuthorizationBuilder) UpdatePermission(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *UpdateAuthorizationBuilder {
	upb.address = address
	upb.channel = channel
	upb.nonce = nonce
	upb.blockExpirationNumber = blockExpirationNumber
	return upb
}

func (upb *UpdateAuthorizationBuilder) Address(address xdr.ID) *UpdateAuthorizationBuilder {
	upb.key = address
	return upb
}

func (upb *UpdateAuthorizationBuilder) Authorize(account xdr.ID, alias string, authorize bool) *UpdateAuthorizationBuilder {
	upb.authorizedAccount = account
	upb.authorizedAlias = alias
	upb.authorize = authorize
	return upb
}

func (upb *UpdateAuthorizationBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {

	action := xdr.Action{
		Address:               *upb.address,
		ChannelID:             *upb.channel,
		Nonce:                 upb.nonce,
		BlockExpirationNumber: upb.blockExpirationNumber,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeUPDATE,
			Update: &xdr.Update{
				Type: xdr.UpdateTypeACCOUNT,
				Account: &xdr.AccountUpdate{
					Type: xdr.AccountUpdateTypeAUTHORIZATION,
					Authorization: &xdr.Authorization{
						Account: xdr.AuthorizedAccount{
							Key:   upb.authorizedAccount,
							Alias: upb.authorizedAlias,
						},
						Authorize: upb.authorize,
					},
				},
			},
		},
	}

	actionStream, err := action.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in action.MarshalBinary")
	}

	signatureSlice := ed25519.Sign(pk, actionStream)

	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}

	transaction := &xdr.Transaction{
		Signature: signature,
		Action:    action,
	}

	if upb.signer != nil {
		transaction.Signer = *upb.signer
	}
	return transaction, nil
}
