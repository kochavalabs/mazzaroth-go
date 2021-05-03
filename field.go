package mazzaroth

import (
	"encoding/json"
	"strconv"
)

// Field
type Field string

func Bool(b bool) Field {
	s := strconv.FormatBool(b)
	return Field(s)
}

func Int32(i int32) Field {
	s := strconv.FormatInt(int64(i), 10)
	return Field(s)
}

func Int64(i int64) Field {
	s := strconv.FormatInt(i, 10)
	return Field(s)
}

func Uint32(u uint32) Field {
	s := strconv.FormatUint(uint64(u), 10)
	return Field(s)
}

func Uint64(u uint64) Field {
	s := strconv.FormatUint(u, 10)
	return Field(s)
}

func Float32(f float32) Field {
	s := strconv.FormatFloat(float64(f), 'E', -1, 32)
	return Field(s)
}

func Float64(f float64) Field {
	s := strconv.FormatFloat(f, 'E', -1, 64)
	return Field(s)
}

func String(s string) Field {
	return Field(s)
}

func Bytes(b []byte) Field {
	s := string(b)
	return Field(s)
}

func JsonBytes(b []byte) Field {
	s := string(b)
	final, _ := json.Marshal(s)
	return Field(final)
}
