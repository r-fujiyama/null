package null

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestStringScanNull(t *testing.T) {
	val := String{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := NewString("", false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringScanString(t *testing.T) {
	val := String{}
	if err := val.Scan("foo"); err != nil {
		t.Fatal(err)
	}

	want := NewString("foo", true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringScanStringEmpty(t *testing.T) {
	val := String{}
	if err := val.Scan(""); err != nil {
		t.Fatal(err)
	}

	want := NewString("", true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringScanByte(t *testing.T) {
	val := String{}
	if err := val.Scan([]byte("foo")); err != nil {
		t.Fatal(err)
	}

	want := NewString("foo", true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringScanByteEmpty(t *testing.T) {
	val := String{}
	if err := val.Scan([]byte("")); err != nil {
		t.Fatal(err)
	}

	want := NewString("", true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringScanError(t *testing.T) {
	val := String{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestStringValueString(t *testing.T) {
	val := NewString("foo", true)
	got, err := val.Value()
	if got != "foo" || err != nil {
		t.Fatalf("want %v, but %v:", "foo", got)
	}
}

func TestStringValueEmpty(t *testing.T) {
	val := NewString("", true)
	got, err := val.Value()
	if got != "" || err != nil {
		t.Fatalf("want %v, but %v:", "", got)
	}
}

func TestStringValueNull(t *testing.T) {
	val := NewString("foo", false)
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", "", got)
	}
}

func TestStringMarshalJSONString(t *testing.T) {
	val := NewString("foo", true)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := `"foo"`
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestStringMarshalJSONEmpty(t *testing.T) {
	val := NewString("", true)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := `""`
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestStringMarshalJSONNull(t *testing.T) {
	val := NewString("foo", false)
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

func TestStringUnmarshalJSONString(t *testing.T) {
	var val String
	err := json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewString("foo", true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringUnmarshalJSONEmpty(t *testing.T) {
	var val String
	err := json.NewDecoder(strings.NewReader(`""`)).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewString("", true)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringUnmarshalJSONNull(t *testing.T) {
	var val String
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := NewString("", false)
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestStringUnmarshalJSONError(t *testing.T) {
	val := String{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestStringIsEmptyString(t *testing.T) {
	val := NewString("foo", true)
	if val.IsEmpty() {
		t.Fatal("should not be empty")
	}
}

func TestStringIsEmptyEmpty(t *testing.T) {
	val := NewString("", true)
	if !val.IsEmpty() {
		t.Fatal("it has to be empty")
	}
}

func TestStringIsEmptyNull(t *testing.T) {
	val := NewString("", false)
	if !val.IsEmpty() {
		t.Fatal("it has to be empty")
	}
}

func TestStringIsNull(t *testing.T) {
	val := NewString("", true)
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = NewString("", false)
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
