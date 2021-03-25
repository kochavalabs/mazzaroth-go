package mazzaroth

import (
	"fmt"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

type ErrMazzaroth struct {
	Code        int
	Msg         string
	OriginalErr error
}

func (e ErrMazzaroth) Error() string {
	return fmt.Sprintf("mazzaroth error: %d %s %s", e.Code, e.Msg, e.OriginalErr.Error())
}

type Mazzaroth interface {
	TransactionSubmit(signature []byte, signer xdr.Authority, action []byte) (*xdr.TransactionSubmitResponse, error)
	ReadOnly(call xdr.Call) (*xdr.ReadonlyResponse, error)
	TransactionLookUp(transactionID xdr.ID) (*xdr.TransactionLookupResponse, error)
	ReceiptLookUp(receiptID xdr.ID) (*xdr.ReceiptLookupResponse, error)
	BlockLookUp(blockID xdr.ID) (*xdr.BlockLookupResponse, error)
	BlockHeaderLookUp(blockID xdr.ID) (*xdr.BlockHeaderLookupResponse, error)
	AccountInfoLookUp(accountID xdr.ID) (*xdr.AccountInfoLookupResponse, error)
	NonceLookUp(accountID xdr.ID) (*xdr.AccountNonceLookupResponse, error)
	ChannelInfoLookUp(channelInfoType xdr.ChannelInfoType) (*xdr.ChannelInfoLookupResponse, error)
}
