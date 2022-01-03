package mazzaroth

import (
	"encoding/json"
	"strconv"

	"github.com/kochavalabs/mazzaroth-xdr/go-xdr/xdr"
)

func Bool(b bool) xdr.Argument {
	s := strconv.FormatBool(b)
	return xdr.Argument(s)
}

func Int32(i int32) xdr.Argument {
	s := strconv.FormatInt(int64(i), 10)
	return xdr.Argument(s)
}

func Int64(i int64) xdr.Argument {
	s := strconv.FormatInt(i, 10)
	return xdr.Argument(s)
}

func Uint32(u uint32) xdr.Argument {
	s := strconv.FormatUint(uint64(u), 10)
	return xdr.Argument(s)
}

func Uint64(u uint64) xdr.Argument {
	s := strconv.FormatUint(u, 10)
	return xdr.Argument(s)
}

func Float64(f float64) xdr.Argument {
	s := strconv.FormatFloat(f, 'e', -1, 64)
	return xdr.Argument(s)
}

func String(s string) xdr.Argument {
	return xdr.Argument(s)
}

func Bytes(b []byte) xdr.Argument {
	s := string(b)
	return xdr.Argument(s)
}

func JsonBytes(b []byte) xdr.Argument {
	s := string(b)
	final, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return xdr.Argument(final)
}
