// +build OS_JS_TARGET_WASM

package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
)

type mazzarothJsWrapperClient struct {
	client mazzaroth.Client
}

func (m *mazzarothJsWrapperClient) accountLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) authorizationLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) blockHeaderLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) blockHeaderList() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) blockHeight() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) blockLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) blockList() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) channelAbi() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) channelConfig() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) receiptLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) transactionLookup() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (m *mazzarothJsWrapperClient) transactionSubmit() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func mzzarothClient(this js.Value, p []js.Value) interface{} {
	var client mazzaroth.Client
	var err error

	if len(p) == 0 {
		// allow client to be created with default values
		client, err = mazzaroth.NewMazzarothClient()
		if err != nil {
			return err.Error()
		}
	}
	wrapperClient := &mazzarothJsWrapperClient{
		client: client,
	}

	// list of client functions
	return js.ValueOf(map[string]interface{}{
		"AccountLookup":       wrapperClient.accountLookup(),
		"AuthorizationLookup": wrapperClient.authorizationLookup(),
		"BlockHeaderLookup":   wrapperClient.blockHeaderLookup(),
		"BlockHeaderList":     wrapperClient.blockHeaderList(),
		"BlockHeight":         wrapperClient.blockHeight(),
		"BlockLookup":         wrapperClient.blockLookup(),
		"BlockList":           wrapperClient.blockList(),
		"ChannelAbi":          wrapperClient.channelAbi(),
		"ChannelConfig":       wrapperClient.channelConfig(),
		"ReceiptLookup":       wrapperClient.receiptLookup(),
		"TransactionLookup":   wrapperClient.transactionLookup(),
		"TransactionSubmit":   wrapperClient.transactionSubmit(),
	})
}

type transacttionBuilderJsWrapper struct {
}

func (tb *transacttionBuilderJsWrapper) Account() js.Func {
	accountFunctions := &accountBuilderFuncs{}
	accountFuncRegister := map[string]interface{}{
		"Account": accountFunctions.Account(),
		"Alias":   accountFunctions.Alias(),
		"Sign":    accountFunctions.Sign(),
	}
	accountBuilderObj := &accountBuilderJsWrapper{
		functions: accountFuncRegister,
	}
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return map[string]interface{}{
			"Account": accountBuilderObj,
		}
	})
}

func transactionBuilder(this js.Value, p []js.Value) interface{} {
	// todo check for required values for tx

	txBuilder := &transacttionBuilderJsWrapper{}
	return js.ValueOf(map[string]interface{}{
		"Account": txBuilder.Account(),
	})
}

type accountBuilderJsWrapper struct {
	functions map[string]interface{}
}

type accountBuilderFuncs struct {
}

func (ab *accountBuilderFuncs) Account() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (ab *accountBuilderFuncs) Alias() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func (ab *accountBuilderFuncs) Sign() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func main() {
	js.Global().Set("NewMazzarothClient", js.FuncOf(mzzarothClient))
	js.Global().Set("TransactionBuilder", js.FuncOf(transactionBuilder))
	// Must keep go program alive when instantiated to allow access to functions
}
