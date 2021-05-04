package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type UpdatePermissionBuilder struct {
	address, channel [32]byte
	nonce            uint64
	signer           *xdr.Authority
}

func (upb *UpdatePermissionBuilder) UpdatePermission(address, channel [32]byte, nonce uint64) *UpdatePermissionBuilder {
	return nil
}

func (upb *UpdatePermissionBuilder) Action(permissionAction int32) *UpdatePermissionBuilder {
	return nil
}

func (upb *UpdatePermissionBuilder) Address(address [32]byte) *UpdatePermissionBuilder {
	return nil
}

func (upb *UpdatePermissionBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	action := xdr.Action{
		Address:   xdr.ID(upb.address),
		ChannelID: xdr.ID(upb.channel),
		Nonce:     upb.nonce,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeUPDATE,
			Update: &xdr.Update{
				Type:       xdr.UpdateTypePERMISSION,
				Permission: &xdr.Permission{},
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
