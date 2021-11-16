package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type UpdatePermissionBuilder struct {
	address, channel      *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	signer                *xdr.Authority
	accountUpdateType     xdr.AccountUpdateType
	key                   xdr.ID
	alias                 string
	authorizedAccount     xdr.ID
	authorizedAlias       string
	authorize             bool
}

func (upb *UpdatePermissionBuilder) UpdatePermission(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *UpdatePermissionBuilder {
	upb.address = address
	upb.channel = channel
	upb.nonce = nonce
	upb.blockExpirationNumber = blockExpirationNumber
	return upb
}

func (upb *UpdatePermissionBuilder) Address(address xdr.ID) *UpdatePermissionBuilder {
	upb.key = address
	return upb
}

func (upb *UpdatePermissionBuilder) Alias(alias string) *UpdatePermissionBuilder {
	upb.alias = alias
	return upb
}

func (upb *UpdatePermissionBuilder) Authorize(accountUpdateType xdr.AccountUpdateType, account xdr.ID, alias string, authorize bool) *UpdatePermissionBuilder {
	upb.accountUpdateType = accountUpdateType
	upb.authorizedAccount = account
	upb.authorizedAlias = alias
	upb.authorize = authorize
	return upb
}

func (upb *UpdatePermissionBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {

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
					Type:  upb.accountUpdateType,
					Alias: &upb.alias,
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
