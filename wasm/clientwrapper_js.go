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
