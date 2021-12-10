package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
)

func callFunction(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func callArguments(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}

func callSign(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	})
}
