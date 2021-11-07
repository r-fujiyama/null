package nulltype

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestByteNewByte(t *testing.T) {
	val := NewByte(byte(97), true)
	want := Byte{Byte: byte(97), Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestByteScanNull(t *testing.T) {
	val := Byte{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := Byte{Byte: byte(0), Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestByteScanByte(t *testing.T) {
	val := Byte{}
	if err := val.Scan(byte(97)); err != nil {
		t.Fatal(err)
	}

	want := Byte{Byte: byte(97), Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestByteScanError(t *testing.T) {
	val := Byte{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestByteValueByte(t *testing.T) {
	val := Byte{Byte: byte(97), Valid: true}
	got, err := val.Value()
	if got != byte(97) || err != nil {
		t.Fatalf("want %v, but %v:", byte(97), got)
	}
}

func TestByteValueNull(t *testing.T) {
	val := Byte{Byte: byte(0), Valid: false}
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", "", got)
	}
}

func TestByteMarshalJSONByte(t *testing.T) {
	val := Byte{Byte: byte(97), Valid: true}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := "97"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestByteMarshalJSONNull(t *testing.T) {
	val := Byte{Byte: byte(0), Valid: false}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestByteUnmarshalJSONByte(t *testing.T) {
	var val Byte
	err := json.NewDecoder(strings.NewReader("97")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Byte{Byte: byte(97), Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestByteUnmarshalJSONNull(t *testing.T) {
	var val Byte
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Byte{Byte: byte(0), Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestByteUnmarshalJSONError(t *testing.T) {
	val := Byte{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestByteString(t *testing.T) {
	val := Byte{Byte: byte(97), Valid: true}
	want := "a"
	got := val.String()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	val = Byte{Byte: byte(0), Valid: false}
	want = "<null>"
	got = val.String()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
