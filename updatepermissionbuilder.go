package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type UpdatePermissionBuilder struct {
	address, channel xdr.ID
	nonce            uint64
	signer           *xdr.Authority
	permissionAction xdr.PermissionAction
	key              xdr.ID
}

func (upb *UpdatePermissionBuilder) UpdatePermission(address, channel xdr.ID, nonce uint64) *UpdatePermissionBuilder {
	upb.address = address
	upb.channel = channel
	upb.nonce = nonce
	return upb
}

func (upb *UpdatePermissionBuilder) Action(permissionAction int32) *UpdatePermissionBuilder {
	upb.permissionAction = xdr.PermissionAction(permissionAction)
	return upb
}

func (upb *UpdatePermissionBuilder) Address(address xdr.ID) *UpdatePermissionBuilder {
	upb.key = address
	return upb
}

func (upb *UpdatePermissionBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {

	action := xdr.Action{
		Address:   upb.address,
		ChannelID: upb.channel,
		Nonce:     upb.nonce,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeUPDATE,
			Update: &xdr.Update{
				Type: xdr.UpdateTypePERMISSION,
				Permission: &xdr.Permission{
					Action: upb.permissionAction,
					Key:    upb.key,
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
