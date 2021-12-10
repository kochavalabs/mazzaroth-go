package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

func authorizationAccount(authorizationBuilder *mazzaroth.AuthorizationBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		id, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		authorizationBuilder.Account(&id)
		return map[string]interface{}{
			"Authorize": authorizationAuthorize(authorizationBuilder),
		}
	})
}

func authorizationAuthorize(authorizationBuilder *mazzaroth.AuthorizationBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		authorizationBuilder.Authorize(args[0].Bool())
		return map[string]interface{}{
			"Sign": sign(authorizationBuilder),
		}
	})
}
