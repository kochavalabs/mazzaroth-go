package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/crypto"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type UpdateContractBuilder struct {
	address, channel      *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	signer                *xdr.Authority
	contract              []byte
	abi                   xdr.Abi
	version               string
}

func (ub *UpdateContractBuilder) UpdateContract(address, channel *xdr.ID, nonce, blockExpirationNumber uint64) *UpdateContractBuilder {
	ub.address = address
	ub.channel = channel
	ub.nonce = nonce
	ub.blockExpirationNumber = blockExpirationNumber
	return ub
}

func (ub *UpdateContractBuilder) Contract(b []byte) *UpdateContractBuilder {
	ub.contract = b
	return ub
}

func (ub *UpdateContractBuilder) Version(version string) *UpdateContractBuilder {
	ub.version = version
	return ub
}

func (ub *UpdateContractBuilder) Abi(abi xdr.Abi) *UpdateContractBuilder {
	ub.abi = abi
	return ub
}

func (ub *UpdateContractBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	if (len(ub.contract) < 0) || ub.version == "" {
		return nil, errors.New("missing require fields")
	}

	hasher := &crypto.Sha3_256Hasher{}
	hash := hasher.Hash(ub.contract)
	xdrHash, err := xdr.HashFromSlice(hash)
	if err != nil {
		return nil, errors.New("unable to create contract hash")
	}
	action := xdr.Action{
		Address:               *ub.address,
		ChannelID:             *ub.channel,
		Nonce:                 ub.nonce,
		BlockExpirationNumber: ub.blockExpirationNumber,
		Category: xdr.ActionCategory{
			Type: xdr.ActionCategoryTypeUPDATE,
			Update: &xdr.Update{
				Type: xdr.UpdateTypeCONTRACT,
				Contract: &xdr.Contract{
					ContractBytes: ub.contract,
					ContractHash:  xdrHash,
					Version:       ub.version,
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

	if ub.signer != nil {
		transaction.Signer = *ub.signer
	}
	return transaction, nil
}
