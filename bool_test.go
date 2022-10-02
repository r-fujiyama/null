package null

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestBoolScanNull(t *testing.T) {
	val := Bool{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: false, Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanString(t *testing.T) {
	val := Bool{}
	if err := val.Scan("true"); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanByte(t *testing.T) {
	val := Bool{}
	if err := val.Scan([]byte("true")); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanInt(t *testing.T) {
	val := Bool{}
	if err := val.Scan(1); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanUint8(t *testing.T) {
	val := Bool{}
	if err := val.Scan(uint8(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanUint16(t *testing.T) {
	val := Bool{}
	if err := val.Scan(uint16(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanUint32(t *testing.T) {
	val := Bool{}
	if err := val.Scan(uint32(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}
func TestBoolScanUint64(t *testing.T) {
	val := Bool{}
	if err := val.Scan(uint64(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanInt8(t *testing.T) {
	val := Bool{}
	if err := val.Scan(int8(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanInt16(t *testing.T) {
	val := Bool{}
	if err := val.Scan(int16(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanInt32(t *testing.T) {
	val := Bool{}
	if err := val.Scan(int32(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}
func TestBoolScanInt64(t *testing.T) {
	val := Bool{}
	if err := val.Scan(int64(1)); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanBool(t *testing.T) {
	val := Bool{}
	if err := val.Scan(true); err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolScanStringParseError(t *testing.T) {
	val := Bool{}
	err := val.Scan("parse error")
	if err == nil {
		t.Fatalf("no parse error is output")
	}
}

func TestBoolScanIntValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(2)
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanInt8ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(int8(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanInt16ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(int16(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanInt32ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(int32(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanInt64ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(int64(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanUint8ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(uint8(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanUint16ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(uint16(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanUint32ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(uint32(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanUint64ValueError(t *testing.T) {
	val := Bool{}
	err := val.Scan(uint64(2))
	if err == nil {
		t.Fatalf("no value error is output")
	}
}

func TestBoolScanByteParseError(t *testing.T) {
	val := Bool{}
	err := val.Scan([]byte("parse error"))
	if err == nil {
		t.Fatalf("no parse error is output")
	}
}

func TestBoolScanTypeError(t *testing.T) {
	val := Bool{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestBoolValueBool(t *testing.T) {
	val := Bool{Bool: true, Valid: true}
	got, err := val.Value()
	if got != true || err != nil {
		t.Fatalf("want %v, but %v:", true, got)
	}
}

func TestBoolValueNull(t *testing.T) {
	val := Bool{Bool: false, Valid: false}
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", "", got)
	}
}

func TestBoolMarshalJSONBool(t *testing.T) {
	val := Bool{Bool: true, Valid: true}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := "true"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestBoolMarshalJSONNull(t *testing.T) {
	val := Bool{Bool: false, Valid: false}
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

func TestBoolUnmarshalJSONBool(t *testing.T) {
	var val Bool
	err := json.NewDecoder(strings.NewReader("true")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: true, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolUnmarshalJSONNull(t *testing.T) {
	var val Bool
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Bool{Bool: false, Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestBoolUnmarshalJSONError(t *testing.T) {
	val := Bool{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestBoolIsNull(t *testing.T) {
	val := Bool{Bool: true, Valid: true}
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = Bool{Bool: false, Valid: false}
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
