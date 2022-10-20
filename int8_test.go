package null

import (
	"bytes"
	"encoding/json"
	"math"
	"strings"
	"testing"
)

func TestInt8ScanNull(t *testing.T) {
	val := Int8{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := NewInt8(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8ScanString(t *testing.T) {
	val := Int8{}
	if err := val.Scan("1"); err != nil {
		t.Fatal(err)
	}

	want := NewInt8(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8ScanByte(t *testing.T) {
	val := Int8{}
	if err := val.Scan([]byte("1")); err != nil {
		t.Fatal(err)
	}

	want := NewInt8(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8ScanInt(t *testing.T) {
	val := Int8{}
	var i int = 1
	if err := val.Scan(i); err != nil {
		t.Fatal(err)
	}

	want := NewInt8(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8ScanMaximumValueOver(t *testing.T) {
	val := Int8{}
	var i int = math.MaxInt8 + 1
	err := val.Scan(i)
	if err == nil || err.Error() != "maximum or minimum value of Int16 exceeded: 128" {
		t.Fatalf("want %v, but %v:", "maximum or minimum value of Int16 exceeded: 128", err)
	}
}

func TestInt8ScanMinimumValueOver(t *testing.T) {
	val := Int8{}
	var i int = math.MinInt8 - 1
	err := val.Scan(i)
	if err == nil || err.Error() != "maximum or minimum value of Int16 exceeded: -129" {
		t.Fatalf("want %v, but %v:", "maximum or minimum value of Int16 exceeded: -129", err)
	}
}

func TestInt8ScanInt8(t *testing.T) {
	val := Int8{}
	var i8 int8 = 1
	if err := val.Scan(i8); err != nil {
		t.Fatal(err)
	}

	want := NewInt8(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8ScanStringParseError(t *testing.T) {
	val := Int8{}
	err := val.Scan("foo")
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestInt8ScanStringByteError(t *testing.T) {
	val := Int8{}
	err := val.Scan([]byte("foo"))
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestInt8ScanTypeError(t *testing.T) {
	val := Int8{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestInt8ValueInt(t *testing.T) {
	val := NewInt8(1, true)
	got, err := val.Value()
	if got != int64(1) || err != nil {
		t.Fatalf("want %v, but %v:", "1", got)
	}
}

func TestInt8ValueZero(t *testing.T) {
	val := NewInt8(0, true)
	got, err := val.Value()
	if got != int64(0) || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestInt8ValueNull(t *testing.T) {
	val := NewInt8(0, false)
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestInt8MarshalJSONInt(t *testing.T) {
	val := NewInt8(1, true)
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

func TestInt8MarshalJSONZero(t *testing.T) {
	val := NewInt8(0, true)
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

func TestInt8MarshalJSONNull(t *testing.T) {
	val := NewInt8(0, false)
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

func TestInt8UnmarshalJSONInt(t *testing.T) {
	var val Int8
	err := json.NewDecoder(strings.NewReader("1")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt8(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8UnmarshalJSONZero(t *testing.T) {
	var val Int8
	err := json.NewDecoder(strings.NewReader("0")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt8(0, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8UnmarshalJSONNull(t *testing.T) {
	var val Int8
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt8(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt8UnmarshalJSONError(t *testing.T) {
	val := Int8{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestInt8IsNull(t *testing.T) {
	val := NewInt8(0, true)
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = NewInt8(0, false)
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
