package null

import (
	"bytes"
	"encoding/json"
	"math"
	"strings"
	"testing"
)

func TestInt16ScanNull(t *testing.T) {
	val := Int16{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := NewInt16(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16ScanString(t *testing.T) {
	val := Int16{}
	if err := val.Scan("1"); err != nil {
		t.Fatal(err)
	}

	want := NewInt16(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16ScanByte(t *testing.T) {
	val := Int16{}
	if err := val.Scan([]byte("1")); err != nil {
		t.Fatal(err)
	}

	want := NewInt16(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16ScanInt(t *testing.T) {
	val := Int16{}
	var i int = 1
	if err := val.Scan(i); err != nil {
		t.Fatal(err)
	}

	want := NewInt16(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16ScanMaximumValueOver(t *testing.T) {
	val := Int16{}
	var i int = math.MaxInt16 + 1
	err := val.Scan(i)
	if err == nil || err.Error() != "maximum or minimum value of Int16 exceeded: 32768" {
		t.Fatalf("want %v, but %v:", "maximum or minimum value of Int16 exceeded: 32768", err)
	}
}

func TestInt16ScanMinimumValueOver(t *testing.T) {
	val := Int16{}
	var i int = math.MinInt16 - 1
	err := val.Scan(i)
	if err == nil || err.Error() != "maximum or minimum value of Int16 exceeded: -32769" {
		t.Fatalf("want %v, but %v:", "maximum or minimum value of Int16 exceeded: -32769", err)
	}
}

func TestInt16ScanInt8(t *testing.T) {
	val := Int16{}
	var i16 int8 = 1
	if err := val.Scan(i16); err != nil {
		t.Fatal(err)
	}

	want := NewInt16(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16ScanInt16(t *testing.T) {
	val := Int16{}
	var i16 int16 = 1
	if err := val.Scan(i16); err != nil {
		t.Fatal(err)
	}

	want := NewInt16(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16ScanStringParseError(t *testing.T) {
	val := Int16{}
	err := val.Scan("foo")
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestInt16ScanStringByteError(t *testing.T) {
	val := Int16{}
	err := val.Scan([]byte("foo"))
	if err == nil {
		t.Fatalf("no error is output")
	}
}

func TestInt16ScanTypeError(t *testing.T) {
	val := Int16{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestInt16ValueInt(t *testing.T) {
	val := NewInt16(1, true)
	got, err := val.Value()
	if got != int64(1) || err != nil {
		t.Fatalf("want %v, but %v:", "1", got)
	}
}

func TestInt16ValueZero(t *testing.T) {
	val := NewInt16(0, true)
	got, err := val.Value()
	if got != int64(0) || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestInt16ValueNull(t *testing.T) {
	val := NewInt16(0, false)
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", 0, got)
	}
}

func TestInt16MarshalJSONInt(t *testing.T) {
	val := NewInt16(1, true)
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

func TestInt16MarshalJSONZero(t *testing.T) {
	val := NewInt16(0, true)
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

func TestInt16MarshalJSONNull(t *testing.T) {
	val := NewInt16(0, false)
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

func TestInt16UnmarshalJSONInt(t *testing.T) {
	var val Int16
	err := json.NewDecoder(strings.NewReader("1")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt16(1, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16UnmarshalJSONZero(t *testing.T) {
	var val Int16
	err := json.NewDecoder(strings.NewReader("0")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt16(0, true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16UnmarshalJSONNull(t *testing.T) {
	var val Int16
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewInt16(0, false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestInt16UnmarshalJSONError(t *testing.T) {
	val := Int16{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestInt16IsNull(t *testing.T) {
	val := NewInt16(0, true)
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = NewInt16(0, false)
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
