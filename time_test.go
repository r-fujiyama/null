package nulltype

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

var testTime time.Time = time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC)

func TestTimeScanNull(t *testing.T) {
	val := Time{}
	if err := val.Scan(nil); err != nil {
		t.Fatal(err)
	}

	want := Time{Time: time.Time{}, Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestTimeScanTime(t *testing.T) {
	val := Time{}
	if err := val.Scan(testTime); err != nil {
		t.Fatal(err)
	}

	want := Time{Time: testTime, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestTimeScanTypeError(t *testing.T) {
	val := Time{}
	err := val.Scan(struct{}{})
	if err == nil || err.Error() != "unsupported type: struct {}" {
		t.Fatalf("want %v, but %v:", "unsupported type: struct {}", err)
	}
}

func TestTimeValueTime(t *testing.T) {
	val := Time{Time: testTime, Valid: true}
	got, err := val.Value()
	if got != testTime || err != nil {
		t.Fatalf("want %v, but %v:", true, got)
	}
}

func TestTimeValueNull(t *testing.T) {
	val := Time{Time: time.Time{}, Valid: false}
	got, err := val.Value()
	if got != nil || err != nil {
		t.Fatalf("want %v, but %v:", "", got)
	}
}

func TestTimeMarshalJSONTime(t *testing.T) {
	val := Time{Time: testTime, Valid: true}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(val); err != nil {
		t.Fatal(err)
	}

	want := `"2022-12-31T23:59:59Z"`
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestTimeMarshalJSONNull(t *testing.T) {
	val := Time{Time: time.Time{}, Valid: false}
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

func TestTimeUnmarshalJSONTime(t *testing.T) {
	var val Time
	err := json.NewDecoder(strings.NewReader(`"2022-12-31T23:59:59Z"`)).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Time{Time: testTime, Valid: true}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestTimeUnmarshalJSONNull(t *testing.T) {
	var val Time
	err := json.NewDecoder(strings.NewReader("null")).Decode(&val)
	if err != nil {
		t.Fatal(err)
	}

	want := Time{Time: time.Time{}, Valid: false}
	if val != want {
		t.Fatalf("want %v, but %v:", want, val)
	}
}

func TestTimeUnmarshalJSONError(t *testing.T) {
	val := Time{}
	err := val.UnmarshalJSON([]byte("foo"))
	if err == nil {
		t.Fatal("no error message is output")
	}
}

func TestTimeIsNull(t *testing.T) {
	val := Time{Time: time.Time{}, Valid: true}
	if val.IsNull() {
		t.Fatal("it has to be not null")
	}

	val = Time{Time: time.Time{}, Valid: false}
	if !val.IsNull() {
		t.Fatal("it has to be not null")
	}
}
