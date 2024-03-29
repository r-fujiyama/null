package null

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestIntScanNull(t *testing.T) {
	val := Int{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := NewInt(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanString(t *testing.T) {
	val := Int{}
	if err := val.Scan("1"); err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanByte(t *testing.T) {
	val := Int{}
	if err := val.Scan([]byte("1")); err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanInt(t *testing.T) {
	val := Int{}
	var i int = 1
	if err := val.Scan(i); err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanInt8(t *testing.T) {
	val := Int{}
	var i8 int8 = 1
	if err := val.Scan(i8); err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanInt16(t *testing.T) {
	val := Int{}
	var i16 int16 = 1
	if err := val.Scan(i16); err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanInt32(t *testing.T) {
	val := Int{}
	var i32 int32 = 1
	if err := val.Scan(i32); err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanInt64(t *testing.T) {
	val := Int{}
	var i64 int64 = 1
	if err := val.Scan(i64); err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntScanStringParseError(t *testing.T) {
	val := Int{}
	err := val.Scan("foo")
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestIntScanStringByteError(t *testing.T) {
	val := Int{}
	err := val.Scan([]byte("foo"))
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestIntScanTypeError(t *testing.T) {
	val := Int{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestIntValueInt(t *testing.T) {
	val := NewInt(1, true)
	got, err := val.Value()
	if got != int64(1) || err != nil {
		t.Fatalf("want %v, but %v:", "1", got)
	}
}

func TestIntValueZero(t *testing.T) {
	val := NewInt(0, true)
	got, err := val.Value()
	if got != int64(0) || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestIntValueNull(t *testing.T) {
	val := NewInt(0, false)
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestIntMarshalJSONInt(t *testing.T) {
	val := NewInt(1, true)
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

func TestIntMarshalJSONZero(t *testing.T) {
	val := NewInt(0, true)
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

func TestIntMarshalJSONNull(t *testing.T) {
	val := NewInt(0, false)
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

func TestIntUnmarshalJSONInt(t *testing.T) {
	var val Int
	err := json.NewDecoder(strings.NewReader("1")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntUnmarshalJSONZero(t *testing.T) {
	var val Int
	err := json.NewDecoder(strings.NewReader("0")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt(0, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntUnmarshalJSONNull(t *testing.T) {
	var val Int
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestIntUnmarshalJSONError(t *testing.T) {
	val := Int{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestIntIsNull(t *testing.T) {
	val := NewInt(0, true)
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = NewInt(0, false)
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
