package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
)

func contractBytes(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func contractVersion(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func contractAbi(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}
