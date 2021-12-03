package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/crypto"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
	"github.com/pkg/errors"
)

type ContractBuilder struct {
	sender                *xdr.ID
	channel               *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	contractBytes         []byte
	abi                   *xdr.Abi
	version               string
}

func (cb *ContractBuilder) Contract(sender, channel *xdr.ID, nonce, blockExpirationNumber uint64) *ContractBuilder {
	cb.sender = sender
	cb.channel = channel
	cb.nonce = nonce
	cb.blockExpirationNumber = blockExpirationNumber
	return cb
}

func (cb *ContractBuilder) ContractBytes(b []byte) *ContractBuilder {
	cb.contractBytes = b
	return cb
}

func (cb *ContractBuilder) Version(version string) *ContractBuilder {
	cb.version = version
	return cb
}

func (cb *ContractBuilder) Abi(abi *xdr.Abi) *ContractBuilder {
	cb.abi = abi
	return cb
}

func (cb *ContractBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	if (len(cb.contractBytes) < 0) || cb.version == "" || cb.abi == nil {
		return nil, errors.New("missing require fields")
	}

	hasher := &crypto.Sha3_256Hasher{}
	hash := hasher.Hash(cb.contractBytes)

	xdrHash, err := xdr.HashFromSlice(hash)
	if err != nil {
		return nil, errors.New("unable to create contract hash")
	}

	data := xdr.Data{
		ChannelID:             *cb.channel,
		Nonce:                 cb.nonce,
		BlockExpirationNumber: cb.blockExpirationNumber,
		Category: xdr.Category{
			Type: xdr.CategoryTypeCONTRACT,
			Contract: &xdr.Contract{
				ContractBytes: cb.contractBytes,
				ContractHash:  xdrHash,
				Version:       cb.version,
				Abi:           *cb.abi,
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
		Sender:    *cb.sender,
		Signer:    signer,
		Signature: signature,
		Data:      data,
	}

	return transaction, nil
}
