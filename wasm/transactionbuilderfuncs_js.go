package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

type transactionBuilderJsWrapper struct{}

func (tb *transactionBuilderJsWrapper) account() js.Func {
	accountBuilder := &mazzaroth.AccountBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		signer, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		channel, err := xdr.IDFromHexString(args[1].String())
		if err != nil {
			return err.Error()
		}
		accountBuilder.Account(&signer, &channel, uint64(args[2].Int()), uint64(args[3].Int()))
		return map[string]interface{}{
			"Alias": accountAlias(accountBuilder),
		}
	})
}

func (tb *transactionBuilderJsWrapper) authorization() js.Func {
	authorizationBuilder := &mazzaroth.AuthorizationBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		signer, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		channel, err := xdr.IDFromHexString(args[1].String())
		if err != nil {
			return err.Error()
		}
		authorizationBuilder.Authorization(&signer, &channel, uint64(args[2].Int()), uint64(args[3].Int()))
		return map[string]interface{}{
			"Account": authorizationAccount(authorizationBuilder),
		}
	})
}

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

func (tb *transactionBuilderJsWrapper) config() js.Func {
	configBuilder := &mazzaroth.ConfigBuilder{}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		signer, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		channel, err := xdr.IDFromHexString(args[1].String())
		if err != nil {
			return err.Error()
		}
		configBuilder.Config(&signer, &channel, uint64(args[2].Int()), uint64(args[3].Int()))
		return map[string]interface{}{
			"Owner": configOwner(configBuilder),
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
