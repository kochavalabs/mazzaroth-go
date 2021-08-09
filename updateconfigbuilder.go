package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type UpdateConfigBuilder struct {
	address, channel      *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	signer                *xdr.Authority
	ownerAddress          *xdr.ID
	adminAddresses        []xdr.ID
}

func (ucb *UpdateConfigBuilder) UpdateConfig(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *UpdateConfigBuilder {
	ucb.address = address
	ucb.channel = channel
	ucb.nonce = nonce
	ucb.blockExpirationNumber = blockExpirationNumber
	return ucb
}

func (ucb *UpdateConfigBuilder) Owner(address *xdr.ID) *UpdateConfigBuilder {
	ucb.ownerAddress = address
	return ucb
}

func (ucb *UpdateConfigBuilder) Admins(addresses ...*xdr.ID) *UpdateConfigBuilder {
	for _, address := range addresses {
		ucb.adminAddresses = append(ucb.adminAddresses, *address)
	}
	return ucb
}

func (ucb *UpdateConfigBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {

	action := xdr.Action{
		Address:               *ucb.address,
		ChannelID:             *ucb.channel,
		Nonce:                 ucb.nonce,
		BlockExpirationNumber: ucb.blockExpirationNumber,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeUPDATE,
			Update: &xdr.Update{
				Type: xdr.UpdateTypeCONFIG,
				ChannelConfig: &xdr.ChannelConfig{
					Owner:  *ucb.ownerAddress,
					Admins: ucb.adminAddresses,
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

	if ucb.signer != nil {
		transaction.Signer = *ucb.signer
	}
	return transaction, nil
}
