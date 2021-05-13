package mazzaroth

import (
	"encoding/json"
	"strconv"

	"github.com/kochavalabs/mazzaroth-xdr/xdr"
)

func Bool(b bool) xdr.Parameter {
	s := strconv.FormatBool(b)
	return xdr.Parameter(s)
}

func Int32(i int32) xdr.Parameter {
	s := strconv.FormatInt(int64(i), 10)
	return xdr.Parameter(s)
}

func Int64(i int64) xdr.Parameter {
	s := strconv.FormatInt(i, 10)
	return xdr.Parameter(s)
}

func Uint32(u uint32) xdr.Parameter {
	s := strconv.FormatUint(uint64(u), 10)
	return xdr.Parameter(s)
}

func Uint64(u uint64) xdr.Parameter {
	s := strconv.FormatUint(u, 10)
	return xdr.Parameter(s)
}

func Float32(f float32) xdr.Parameter {
	s := strconv.FormatFloat(float64(f), 'E', -1, 32)
	return xdr.Parameter(s)
}

func Float64(f float64) xdr.Parameter {
	s := strconv.FormatFloat(f, 'E', -1, 64)
	return xdr.Parameter(s)
}

func String(s string) xdr.Parameter {
	return xdr.Parameter(s)
}

func Bytes(b []byte) xdr.Parameter {
	s := string(b)
	return xdr.Parameter(s)
}

func JsonBytes(b []byte) xdr.Parameter {
	s := string(b)
	final, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return xdr.Parameter(final)
}
