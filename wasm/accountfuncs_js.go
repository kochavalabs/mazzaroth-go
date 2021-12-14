package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
)

func accountAlias(accountBuilder *mazzaroth.AccountBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		accountBuilder.Alias(args[0].String())
		return map[string]interface{}{
			"Sign": sign(accountBuilder),
		}
	})
}
