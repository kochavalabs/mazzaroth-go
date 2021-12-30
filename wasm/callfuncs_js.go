package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
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
	// dev note:: we are expecting args to be a string and that the caller
	// will pass multiple string args i.e func("x","y","z")
	// Alternatively we could construct a json array and pass a
	// json string to this function and unmarshal it into a []xdr.Agruments
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		arguments := make([]xdr.Argument, 0, 0)
		for _, arg := range args {
			a := xdr.Argument(arg.String())
			arguments = append(arguments, a)
		}
		callBuilder.Arguments(arguments...)
		return map[string]interface{}{
			"Sign": sign(callBuilder),
		}
	})
}
