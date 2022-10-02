package null

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestFloat32ScanNull(t *testing.T) {
	val := Float32{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanString(t *testing.T) {
	val := Float32{}
	if err := val.Scan("1.1"); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanByte(t *testing.T) {
	val := Float32{}
	if err := val.Scan([]byte("1.1")); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanInt(t *testing.T) {
	val := Float32{}
	var i int = 1
	if err := val.Scan(i); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanInt8(t *testing.T) {
	val := Float32{}
	var i8 int8 = 1
	if err := val.Scan(i8); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanInt16(t *testing.T) {
	val := Float32{}
	var i16 int16 = 1
	if err := val.Scan(i16); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanInt32(t *testing.T) {
	val := Float32{}
	var i32 int32 = 1
	if err := val.Scan(i32); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanInt64(t *testing.T) {
	val := Float32{}
	var i64 int64 = 1
	if err := val.Scan(i64); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanFloat32(t *testing.T) {
	val := Float32{}
	var f32 float32 = 1.1
	if err := val.Scan(f32); err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32ScanStringParseError(t *testing.T) {
	val := Float32{}
	err := val.Scan("foo")
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestFloat32ScanStringByteError(t *testing.T) {
	val := Float32{}
	err := val.Scan([]byte("foo"))
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestFloat32ScanTypeError(t *testing.T) {
	val := Float32{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestFloat32ValueFloat(t *testing.T) {
	val := NewFloat32(1.1, true)
	got, err := val.Value()
	if got != float32(1.1) || err != nil {
		t.Fatalf("want %v, but %v:", 1.1, got)
	}
}

func TestFloat32ValueZero(t *testing.T) {
	val := NewFloat32(0, true)
	got, err := val.Value()
	if got != float32(0.0) || err != nil {
		t.Fatalf("want %v, but %v:", 0.0, got)
	}
}

func TestFloat32ValueNull(t *testing.T) {
	val := NewFloat32(0, false)
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", 0.0, got)
	}
}

func TestFloat32MarshalJSONFloat(t *testing.T) {
	val := NewFloat32(1.1, true)
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

func TestFloat32MarshalJSONZero(t *testing.T) {
	val := NewFloat32(0, true)
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

func TestFloat32MarshalJSONNull(t *testing.T) {
	val := NewFloat32(0, false)
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

func TestFloat32UnmarshalJSONFloat(t *testing.T) {
	var val Float32
	err := json.NewDecoder(strings.NewReader("1.1")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(1.1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32UnmarshalJSONZero(t *testing.T) {
	var val Float32
	err := json.NewDecoder(strings.NewReader("0")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(0, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32UnmarshalJSONNull(t *testing.T) {
	var val Float32
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewFloat32(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestFloat32UnmarshalJSONError(t *testing.T) {
	val := Float32{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestFloat32IsNull(t *testing.T) {
	val := NewFloat32(0, true)
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = NewFloat32(0, false)
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
