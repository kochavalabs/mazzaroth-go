package mazzaroth

import "testing"

func TestStringBool(t *testing.T) {
	p := Bool(false)
	if p != "false" {
		t.Fatalf("string %s does not match string %s", p, "false")
	}
}

func TestStringInt32(t *testing.T) {
	p := Int32(1)
	if p != "1" {
		t.Fatalf("string %s does not match string %s", p, "1")
	}
}

func TestStringInt64(t *testing.T) {
	p := Int64(1)
	if p != "1" {
		t.Fatalf("string %s does not match string %s", p, "1")
	}
}

func TestStringUint32(t *testing.T) {
	p := Uint32(1)
	if p != "1" {
		t.Fatalf("string %s does not match string %s", p, "1")
	}
}

func TestStringUint64(t *testing.T) {
	p := Uint64(1)
	if p != "1" {
		t.Fatalf("string %s does not match string %s", p, "1")
	}
}

func TestStringFloat64(t *testing.T) {
	p := Float64(3.14159265)
	if p != "3.14159265e+00" {
		t.Fatalf("string %s does not match string %s", p, "3.14159265e+00")
	}
}

func TestStringString(t *testing.T) {
	p := String("test")
	if p != "test" {
		t.Fatalf("string %s does not match string %s", p, "test")
	}
}

func TestStringByte(t *testing.T) {
	p := Bytes([]byte("test"))
	if p != "test" {
		t.Fatalf("string %s does not match string %s", p, "test")
	}
}

func TestStringJson(t *testing.T) {
	p := JsonBytes([]byte(`{"test": "hello"}`))
	if p != "\"{\\\"test\\\": \\\"hello\\\"}\"" {
		t.Fatalf("string %s does not match string %s", p, "\"{\\\"test\\\": \\\"hello\\\"}\"")
	}
}
