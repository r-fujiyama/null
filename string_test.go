package nulltype

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestStringNewString(t *testing.T) {
	s := NewString("")
	want := String{String: "", Valid: false}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}

	s = NewString("foo")
	want = String{String: "foo", Valid: true}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringScanString(t *testing.T) {
	s := String{}
	if err := s.Scan("foo"); err != nil {
		t.Fatal(err)
	}

	want := String{String: "foo", Valid: true}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringScanByte(t *testing.T) {
	s := String{}
	if err := s.Scan([]byte("foo")); err != nil {
		t.Fatal(err)
	}

	want := String{String: "foo", Valid: true}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringScanEmpty(t *testing.T) {
	s := String{}
	if err := s.Scan([]byte("")); err != nil {
		t.Fatal(err)
	}

	want := String{String: "", Valid: true}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringScanNull(t *testing.T) {
	s := String{}
	if err := s.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := String{String: "", Valid: false}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringScanError(t *testing.T) {
	s := String{}
	err := s.Scan(1)
	if err == nil || err.Error() != "got data of type int" {
		t.Fatalf("want %v, but %v:", "got data of type int", err)
	}
}

func TestStringValueString(t *testing.T) {
	s := String{String: "foo", Valid: true}
	str, err := s.Value()
	if str != "foo" || err != nil {
		t.Fatalf("want %v, but %v:", "foo", str)
	}
}

func TestStringValueEmpty(t *testing.T) {
	s := String{String: "", Valid: true}
	str, err := s.Value()
	if str != "" || err != nil {
		t.Fatalf("want %v, but %v:", "", str)
	}
}

func TestStringValueNull(t *testing.T) {
	s := String{String: "foo", Valid: false}
	str, err := s.Value()
	if str != nil || err != nil {
		t.Fatalf("want %v, but %v:", "", str)
	}
}

func TestStringMarshalJSONString(t *testing.T) {
	s := String{String: "foo", Valid: true}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(s); err != nil {
		t.Fatal(err)
	}

	want := `"foo"`
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestStringMarshalJSONEmpty(t *testing.T) {
	s := String{String: "", Valid: true}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(s); err != nil {
		t.Fatal(err)
	}

	want := `""`
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestStringMarshalJSONNull(t *testing.T) {
	s := String{String: "foo", Valid: false}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(s); err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestStringUnmarshalJSONString(t *testing.T) {
	var s String
	err := json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	want := String{String: "foo", Valid: true}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringUnmarshalJSONEmpty(t *testing.T) {
	var s String
	err := json.NewDecoder(strings.NewReader(`""`)).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	want := String{String: "", Valid: true}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringUnmarshalJSONNull(t *testing.T) {
	var s String
	err := json.NewDecoder(strings.NewReader("null")).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	want := String{String: "", Valid: false}
	if s != want {
		t.Fatalf("want %v, but %v:", want, s)
	}
}

func TestStringUnmarshalJSONError(t *testing.T) {
	nullString := String{}
	err := nullString.UnmarshalJSON([]byte("a"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestStringIsEmptyString(t *testing.T) {
	s := String{String: "foo", Valid: true}
	if s.IsEmpty() {
		t.Fatal("should not be empty")
	}
}

func TestStringIsEmptyEmpty(t *testing.T) {
	s := String{String: "", Valid: true}
	if !s.IsEmpty() {
		t.Fatal("it has to be empty")
	}
}

func TestStringIsEmptyNull(t *testing.T) {
	s := String{String: "", Valid: false}
	if !s.IsEmpty() {
		t.Fatal("it has to be empty")
	}
}
