package mazzaroth

import (
	"crypto/ed25519"

	"github.com/kochavalabs/crypto"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
	"github.com/pkg/errors"
)

type ContractBuilder struct {
	categoryType          xdr.CategoryType
	sender                *xdr.ID
	channel               *xdr.ID
	nonce                 uint64
	blockExpirationNumber uint64
	pause                 bool
	contractBytes         []byte
	abi                   xdr.Abi
	version               string
}

func (cb *ContractBuilder) Contract(sender, channel *xdr.ID, nonce, blockExpirationNumber uint64) *ContractBuilder {
	cb.sender = sender
	cb.channel = channel
	cb.nonce = nonce
	cb.blockExpirationNumber = blockExpirationNumber
	return cb
}

func (cb *ContractBuilder) Delete() *ContractBuilder {
	cb.categoryType = xdr.CategoryTypeDELETE
	return cb
}

func (cb *ContractBuilder) Deploy(version string, abi xdr.Abi, b []byte) *ContractBuilder {
	cb.categoryType = xdr.CategoryTypeDEPLOY
	cb.contractBytes = b
	cb.version = version
	cb.abi = abi
	return cb
}

func (cb *ContractBuilder) Pause(pause bool) *ContractBuilder {
	cb.categoryType = xdr.CategoryTypePAUSE
	cb.pause = pause
	return cb
}

func (cb *ContractBuilder) Sign(pk ed25519.PrivateKey) (*xdr.Transaction, error) {
	hasher := &crypto.Sha3_256Hasher{}
	hash := hasher.Hash(cb.contractBytes)

	xdrHash, err := xdr.HashFromSlice(hash)
	if err != nil {
		return nil, errors.New("unable to create contract hash")
	}

	var data xdr.Data
	switch cb.categoryType {
	case xdr.CategoryTypeDELETE:
		data = xdr.Data{
			ChannelID:             *cb.channel,
			Nonce:                 cb.nonce,
			BlockExpirationNumber: cb.blockExpirationNumber,
			Category: xdr.Category{
				Type: xdr.CategoryTypeDELETE,
			},
		}
	case xdr.CategoryTypeDEPLOY:
		if len(cb.contractBytes) == 0 || version == "" || len(cb.abi.Functions) == 0 {
			return nil, errors.New("missing required fields for deploy transaction")
		}
		data = xdr.Data{
			ChannelID:             *cb.channel,
			Nonce:                 cb.nonce,
			BlockExpirationNumber: cb.blockExpirationNumber,
			Category: xdr.Category{
				Type: xdr.CategoryTypeDEPLOY,
				Contract: &xdr.Contract{
					ContractBytes: cb.contractBytes,
					ContractHash:  xdrHash,
					Version:       cb.version,
					Abi:           cb.abi,
				},
			},
		}
	case xdr.CategoryTypePAUSE:
		data = xdr.Data{
			ChannelID:             *cb.channel,
			Nonce:                 cb.nonce,
			BlockExpirationNumber: cb.blockExpirationNumber,
			Category: xdr.Category{
				Type:  xdr.CategoryTypePAUSE,
				Pause: &cb.pause,
			},
		}
	default:
		return nil, errors.New("unknown contract category type")
	}

	dataBytes, err := data.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "in data.MarshalBinary")
	}

	signatureSlice := ed25519.Sign(pk, dataBytes)
	signature, err := xdr.SignatureFromSlice(signatureSlice)
	if err != nil {
		return nil, errors.Wrap(err, "in signing the transaction")
	}

	transaction := &xdr.Transaction{
		Sender:    *cb.sender,
		Signature: signature,
		Data:      data,
	}

	return transaction, nil
}
