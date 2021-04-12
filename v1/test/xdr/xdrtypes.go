// Package xdr is automatically generated
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"bytes"
	"encoding"
	"fmt"
	"io"

	"github.com/stellar/go-xdr/xdr3"
)

// Unmarshal reads an xdr element from `r` into `v`.
func Unmarshal(r io.Reader, v interface{}) (int, error) {
	// delegate to xdr package's Unmarshal
	return xdr.Unmarshal(r, v)
}

// Marshal writes an xdr element `v` into `w`.
func Marshal(w io.Writer, v interface{}) (int, error) {
	// delegate to xdr package's Marshal
	return xdr.Marshal(w, v)
}

// Namspace start example

// Start typedef section

// ID generated typedef
type ID [32]byte

// XDRMaxSize implements the Sized interface for ID
func (s ID) XDRMaxSize() int {
	return 32
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s ID) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *ID) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*ID)(nil)
	_ encoding.BinaryUnmarshaler = (*ID)(nil)
)

// End typedef section

// Start struct section

// Bar generated struct
type Bar struct {
	Id ID
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Bar) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Bar) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Bar)(nil)
	_ encoding.BinaryUnmarshaler = (*Bar)(nil)
)

// End struct section

// Start enum section

// End enum section

// Start union section

// End union section

// Namspace end example
// Namspace start example

// Start typedef section

// End typedef section

// Start struct section

// Foo generated struct
type Foo struct {
	Status FooStatus

	One string `xdrmaxsize:"256"`

	Two string `xdrmaxsize:"256"`

	Three string `xdrmaxsize:"256"`
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s Foo) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *Foo) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*Foo)(nil)
	_ encoding.BinaryUnmarshaler = (*Foo)(nil)
)

// End struct section

// Start enum section

// FooStatus generated enum
type FooStatus int32

const (

	// FooStatusZero enum value 0
	FooStatusZero FooStatus = 0

	// FooStatusOne enum value 1
	FooStatusOne FooStatus = 1

	// FooStatusTwo enum value 2
	FooStatusTwo FooStatus = 2

	// FooStatusThree enum value 3
	FooStatusThree FooStatus = 3
)

// FooStatusMap generated enum map
var FooStatusMap = map[int32]string{

	0: "FooStatusZero",

	1: "FooStatusOne",

	2: "FooStatusTwo",

	3: "FooStatusThree",
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for FooStatus
func (s FooStatus) ValidEnum(v int32) bool {
	_, ok := FooStatusMap[v]
	return ok
}

// String returns the name of `e`
func (s FooStatus) String() string {
	name, _ := FooStatusMap[int32(s)]
	return name
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (s FooStatus) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	_, err := Marshal(b, s)
	return b.Bytes(), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (s *FooStatus) UnmarshalBinary(inp []byte) error {
	_, err := Unmarshal(bytes.NewReader(inp), s)
	return err
}

var (
	_ encoding.BinaryMarshaler   = (*FooStatus)(nil)
	_ encoding.BinaryUnmarshaler = (*FooStatus)(nil)
)

// End enum section

// Start union section

// End union section

// Namspace end example
var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
