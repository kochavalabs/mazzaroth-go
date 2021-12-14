package main

import (
	"syscall/js"

	"github.com/kochavalabs/mazzaroth-go"
	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

func configOwner(configBuilder *mazzaroth.ConfigBuilder) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		id, err := xdr.IDFromHexString(args[0].String())
		if err != nil {
			return err.Error()
		}
		configBuilder.Owner(&id)
		return map[string]interface{}{
			"Admins": configAdmins(configBuilder),
		}
	})
}

func configAdmins(configBuilder *mazzaroth.ConfigBuilder) js.Func {
	// dev note:: we are expecting args to be a string and that the caller
	// will pass multiple string args i.e func("x","y","z")
	// Alternatively we could construct a json array and pass a
	// json string to this function and unmarshal it into a []xdr.ID
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ids := make([]xdr.ID, 0, 0)
		for _, admin := range args {
			id, err := xdr.IDFromHexString(admin.String())
			if err != nil {
				return err.Error()
			}
			ids = append(ids, id)
		}
		configBuilder.Admins(ids...)
		return map[string]interface{}{
			"Sign": sign(configBuilder),
		}
	})
}
