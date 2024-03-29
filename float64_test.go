package null

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestFloat64ScanNull(t *testing.T) {
	val := Float64{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanString(t *testing.T) {
	val := Float64{}
	if err := val.Scan("1.1"); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanByte(t *testing.T) {
	val := Float64{}
	if err := val.Scan([]byte("1.1")); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanInt(t *testing.T) {
	val := Float64{}
	var i int = 1
	if err := val.Scan(i); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanInt8(t *testing.T) {
	val := Float64{}
	var i8 int8 = 1
	if err := val.Scan(i8); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanInt16(t *testing.T) {
	val := Float64{}
	var i16 int16 = 1
	if err := val.Scan(i16); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanInt32(t *testing.T) {
	val := Float64{}
	var i32 int32 = 1
	if err := val.Scan(i32); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanInt64(t *testing.T) {
	val := Float64{}
	var i64 int64 = 1
	if err := val.Scan(i64); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanFloat64(t *testing.T) {
	val := Float64{}
	var f64 float64 = 1.1
	if err := val.Scan(f64); err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64ScanStringParseError(t *testing.T) {
	val := Float64{}
	err := val.Scan("foo")
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestFloat64ScanStringByteError(t *testing.T) {
	val := Float64{}
	err := val.Scan([]byte("foo"))
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestFloat64ScanTypeError(t *testing.T) {
	val := Float64{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestFloat64ValueFloat(t *testing.T) {
	val := NewFloat64(1.1, true)
	got, err := val.Value()
	if got != 1.1 || err != nil {
		t.Fatalf("want %v, but %v:", 1.1, got)
	}
}

func TestFloat64ValueZero(t *testing.T) {
	val := NewFloat64(0, true)
	got, err := val.Value()
	if got != 0.0 || err != nil {
		t.Fatalf("want %v, but %v:", 0.0, got)
	}
}

func TestFloat64ValueNull(t *testing.T) {
	val := NewFloat64(0, false)
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", 0.0, got)
	}
}

func TestFloat64MarshalJSONFloat(t *testing.T) {
	val := NewFloat64(1.1, true)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := "1.1"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestFloat64MarshalJSONZero(t *testing.T) {
	val := NewFloat64(0, true)
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

func TestFloat64MarshalJSONNull(t *testing.T) {
	val := NewFloat64(0, false)
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

func TestFloat64UnmarshalJSONFloat(t *testing.T) {
	var val Float64
	err := json.NewDecoder(strings.NewReader("1.1")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64UnmarshalJSONZero(t *testing.T) {
	var val Float64
	err := json.NewDecoder(strings.NewReader("0")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(0, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64UnmarshalJSONNull(t *testing.T) {
	var val Float64
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewFloat64(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat64UnmarshalJSONError(t *testing.T) {
	val := Float64{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestFloat64IsNull(t *testing.T) {
	val := NewFloat64(0, true)
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = NewFloat64(0, false)
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
