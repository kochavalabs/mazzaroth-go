package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

type transactionBuilderJsWrapper struct{}

func (tb *transactionBuilderJsWrapper) call() js.Func {
	callBuilder := &mazzaroth.CallBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		signer, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		channel, err := xdr.IDFromHexString(args[1].String())
		if err != nil {
			return err.Error()
		}
		callBuilder.Call(&signer, &channel, uint64(args[2].Int()), uint64(args[3].Int()))
		return map[string]interface{}{
			"Function": callFunction(callBuilder),
		}
	})
}

func (tb *transactionBuilderJsWrapper) contract() js.Func {
	contractBuilder := &mazzaroth.ContractBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		signer, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		channel, err := xdr.IDFromHexString(args[1].String())
		if err != nil {
			return err.Error()
		}
		contractBuilder.Contract(&signer, &channel, uint64(args[2].Int()), uint64(args[3].Int()))
		return map[string]interface{}{
			"ContractBytes": contractBytes(contractBuilder),
		}
	})
}
