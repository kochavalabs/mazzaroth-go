package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type ConfigBuilder struct {
	sender                *xdr.ID
	channel               *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	ownerAddress          *xdr.ID
	adminAddresses        []xdr.ID
}

func (cfgb *ConfigBuilder) Config(sender, channel *xdr.ID, nonce, blockExpirationNumber uint64) *ConfigBuilder {
	cfgb.sender = sender
	cfgb.channel = channel
	cfgb.nonce = nonce
	cfgb.blockExpirationNumber = blockExpirationNumber
	return cfgb
}

func (cfgb *ConfigBuilder) Owner(address *xdr.ID) *ConfigBuilder {
	cfgb.ownerAddress = address
	return cfgb
}

func (cfgb *ConfigBuilder) Admins(addresses ...*xdr.ID) *ConfigBuilder {
	for _, address := range addresses {
		cfgb.adminAddresses = append(cfgb.adminAddresses, *address)
	}
	return cfgb
}

func (cfgb *ConfigBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {

	data := xdr.Data{
		ChannelID:             *cfgb.channel,
		Nonce:                 cfgb.nonce,
		BlockExpirationNumber: cfgb.blockExpirationNumber,
		Category: xdr.Category{
			Type: xdr.CategoryTypeCONFIG,
			Config: &xdr.Config{
				Owner:  *cfgb.ownerAddress,
				Admins: cfgb.adminAddresses,
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
		Sender:    *cfgb.sender,
		Signer:    signer,
		Signature: signature,
		Data:      data,
	}
	return transaction, nil
}
