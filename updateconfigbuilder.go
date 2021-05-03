package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type UpdateConfigBuilder struct {
	address, channel [32]byte
	nonce            uint64
	signer           xdr.Authority
}

func (ucb *UpdateConfigBuilder) UpdateConfig(address, channel [32]byte, nonce uint64) *UpdateConfigBuilder {
	ucb.address = address
	ucb.channel = channel
	ucb.nonce = nonce
	return ucb
}

func (ucb *UpdateConfigBuilder) Owner(ID [32]byte) *UpdateConfigBuilder {
	return nil
}

func (ucb *UpdateConfigBuilder) ChannelName(name string) *UpdateConfigBuilder {
	return nil
}

func (ucb *UpdateConfigBuilder) Admins(IDs [][32]byte) *UpdateConfigBuilder {
	return nil
}

func (ucb *UpdateConfigBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {

	action := xdr.Action{
		Address:   xdr.ID(ucb.address),
		ChannelID: xdr.ID(ucb.channel),
		Nonce:     ucb.nonce,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeUPDATE,
			Update: &xdr.Update{
				Type:          xdr.UpdateTypeCONFIG,
				ChannelConfig: &xdr.ChannelConfig{},
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
	if &transaction.Signer == nil {
		transaction.Signer, err = xdr.NewAuthority(xdr.AuthorityTypeNONE, nil)
		if err != nil {
			return nil, errors.Wrap(err, "in xdr.NewAuthority(xdr.AuthorityTypeNONE)")
		}
	}
	return nil, nil
}
