package nulltype

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestInt32NewInt32(t *testing.T) {
	val := NewInt32(1, true)
	want := Int32{Int32: 1, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32ScanNull(t *testing.T) {
	val := Int32{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 0, Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32ScanString(t *testing.T) {
	val := Int32{}
	if err := val.Scan("1"); err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 1, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32ScanByte(t *testing.T) {
	val := Int32{}
	if err := val.Scan([]byte("1")); err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 1, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32ScanInt(t *testing.T) {
	val := Int32{}
	var i int = 1
	if err := val.Scan(i); err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 1, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32ScanInt16(t *testing.T) {
	val := Int32{}
	var i16 int16 = 1
	if err := val.Scan(i16); err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 1, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32ScanInt32(t *testing.T) {
	val := Int32{}
	var i32 int32 = 1
	if err := val.Scan(i32); err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 1, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32ScanStringParseError(t *testing.T) {
	val := Int32{}
	err := val.Scan("foo")
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestInt32ScanStringByteError(t *testing.T) {
	val := Int32{}
	err := val.Scan([]byte("foo"))
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestInt32ScanTypeError(t *testing.T) {
	val := Int32{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "got data of type struct {}" {
		t.Fatalf("want %v, but %v:", "got data of type struct {}", err)
	}
}

func TestInt32ValueInt(t *testing.T) {
	val := Int32{Int32: 1, Valid: true}
	got, err := val.Value()
	if got != int32(1) || err != nil {
		t.Fatalf("want %v, but %v:", "1", got)
	}
}

func TestInt32ValueZero(t *testing.T) {
	val := Int32{Int32: 0, Valid: true}
	got, err := val.Value()
	if got != int32(0) || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestInt32ValueNull(t *testing.T) {
	val := Int32{Int32: 0, Valid: false}
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestInt32MarshalJSONInt(t *testing.T) {
	val := Int32{Int32: 1, Valid: true}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := "1"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestInt32MarshalJSONZero(t *testing.T) {
	val := Int32{Int32: 0, Valid: true}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := "0"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestInt32MarshalJSONNull(t *testing.T) {
	val := Int32{Int32: 0, Valid: false}
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

func TestInt32UnmarshalJSONInt(t *testing.T) {
	var val Int32
	err := json.NewDecoder(strings.NewReader("1")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 1, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32UnmarshalJSONZero(t *testing.T) {
	var val Int32
	err := json.NewDecoder(strings.NewReader("0")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 0, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32UnmarshalJSONNull(t *testing.T) {
	var val Int32
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Int32{Int32: 0, Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt32UnmarshalJSONError(t *testing.T) {
	val := Int32{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestInt32isZeroOrNullInt(t *testing.T) {
	val := Int32{Int32: 1, Valid: true}
	if val.IsZeroOrNull() {
		t.Fatal("should not be zero or null")
	}
}

func TestInt32isZeroOrNullZero(t *testing.T) {
	val := Int32{Int32: 0, Valid: true}
	if !val.IsZeroOrNull() {
		t.Fatal("it has to be zero or null")
	}
}

func TestInt32isZeroOrNullNull(t *testing.T) {
	val := Int32{Int32: 0, Valid: false}
	if !val.IsZeroOrNull() {
		t.Fatal("it has to be zero or null")
	}
}
