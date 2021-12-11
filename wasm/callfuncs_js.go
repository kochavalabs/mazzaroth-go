package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
)

func callFunction(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callBuilder.Function(args[0].String())
		return map[string]interface{}{
			"Arguments": callArguments(callBuilder),
		}
	})
}

func callArguments(callBuilder *mazzaroth.CallBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return map[string]interface{}{
			"Sign": sign(callBuilder),
		}
	})
}
