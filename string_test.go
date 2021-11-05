package null

import "testing"

func TestNullString(t *testing.T) {
	t.Run("New nullString", func(t *testing.T) {
		nullString := NewString("")
		if nullString.String != "" || nullString.Valid {
			t.Fatalf("null.String wrong. expected=%#v, got=%#v", String{String: "", Valid: false}, nullString)
		}
	})

	t.Run("New notNullString", func(t *testing.T) {
		nullString := NewString("test")
		if nullString.String != "test" || !nullString.Valid {
			t.Fatalf("null.String wrong. expected=%#v, got=%#v", String{String: "test", Valid: true}, nullString)
		}
	})

	t.Run("Scan string", func(t *testing.T) {
		nullString := String{}
		err := nullString.Scan([]byte("test"))
		if err != nil {
			t.Fatalf("error message wrong. expected=%v, got=%s", nil, err)
		}
		if nullString.String != "test" || !nullString.Valid {
			t.Fatalf("null.String wrong. expected=%#v, got=%#v", String{String: "test", Valid: true}, nullString)
		}
	})

	t.Run("Scan empty string", func(t *testing.T) {
		nullString := String{}
		err := nullString.Scan([]byte(""))
		if err != nil {
			t.Fatalf("error message wrong. expected=%v, got=%s", nil, err)
		}
		if nullString.String != "" || !nullString.Valid {
			t.Fatalf("null.String wrong. expected=%#v, got=%#v", String{String: "", Valid: true}, nullString)
		}
	})

	t.Run("Scan nil", func(t *testing.T) {
		nullString := String{}
		err := nullString.Scan(nil)
		if err != nil {
			t.Fatalf("error message wrong. expected=%v, got=%s", nil, err)
		}
		if nullString.String != "" && nullString.Valid {
			t.Fatalf("null.String. expected=%#v, got=%#v", String{String: "", Valid: false}, nullString)
		}
	})

	t.Run("Scan no byte data", func(t *testing.T) {
		nullString := String{}
		err := nullString.Scan(1)
		if err == nil || err.Error() != "got data of type int but wanted []uint8" {
			t.Fatalf("error wrong. expected=%s, got=%s", "got data of type int but wanted []uint8", err)
		}
	})

	t.Run("Value string", func(t *testing.T) {
		nullString := String{String: "test", Valid: true}
		str, err := nullString.Value()
		if str != "test" || err != nil {
			t.Fatalf("str or err wrong. expected=(str=test, err=nil) , got=(str=%v,err=%v)", str, err)
		}
	})

	t.Run("Value empty", func(t *testing.T) {
		nullString := String{String: "", Valid: true}
		str, err := nullString.Value()
		if str != "" || err != nil {
			t.Fatalf("str or err wrong. expected=(str=, err=<nil>) , got=(str=%v,err=%v)", str, err)
		}
	})

	t.Run("Value null", func(t *testing.T) {
		nullString := String{String: "", Valid: false}
		str, err := nullString.Value()
		if str != nil || err != nil {
			t.Fatalf("str or err wrong. expected=(str=<nil>, err=<nil>) , got=(str=%v,err=%v)", str, err)
		}
	})

	t.Run("MarshalJSON string", func(t *testing.T) {
		nullString := String{String: "test", Valid: true}
		jsonValue, err := nullString.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if string(jsonValue) != "\"test\"" {
			t.Fatalf("json value wrong. expected=(jsonValue=\"test\") , got=(jsonValue=%s)", jsonValue)
		}
	})

	t.Run("MarshalJSON empty", func(t *testing.T) {
		nullString := String{String: "", Valid: true}
		jsonValue, err := nullString.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if string(jsonValue) != "null" {
			t.Fatalf("json value wrong. expected=(jsonValue=null) , got=(jsonValue=%s)", jsonValue)
		}
	})

	t.Run("MarshalJSON null", func(t *testing.T) {
		nullString := String{String: "", Valid: false}
		jsonValue, err := nullString.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if string(jsonValue) != "null" {
			t.Fatalf("json value wrong. expected=(jsonValue=null) , got=(jsonValue=%s)", jsonValue)
		}
	})

	t.Run("UnmarshalJSON string", func(t *testing.T) {
		nullString := String{}
		err := nullString.UnmarshalJSON([]byte("\"test\""))
		if err != nil {
			t.Fatal(err)
		}
		if nullString.String != "test" || !nullString.Valid {
			t.Fatalf("null.String wrong. expected=%#v, got=%#v", String{String: "test", Valid: true}, nullString)
		}
	})

	t.Run("UnmarshalJSON empty", func(t *testing.T) {
		nullString := String{}
		err := nullString.UnmarshalJSON([]byte("\"\""))
		if err != nil {
			t.Fatal(err)
		}
		if nullString.String != "" || nullString.Valid {
			t.Fatalf("null.String wrong. expected=%#v, got=%#v", String{String: "", Valid: false}, nullString)
		}
	})

	t.Run("UnmarshalJSON null", func(t *testing.T) {
		nullString := String{}
		err := nullString.UnmarshalJSON([]byte("null"))
		if err != nil {
			t.Fatal(err)
		}
		if nullString.String != "" || nullString.Valid {
			t.Fatalf("null.String wrong. expected=%#v, got=%#v", String{String: "", Valid: false}, nullString)
		}
	})

	t.Run("UnmarshalJSON error", func(t *testing.T) {
		nullString := String{}
		err := nullString.UnmarshalJSON([]byte("a"))
		if err == nil {
			t.Fatal("no error message is output")
		}
	})

	t.Run("IsEmpty string", func(t *testing.T) {
		nullString := String{String: "test", Valid: true}
		if nullString.IsEmpty() {
			t.Fatal("should not be empty")
		}
	})

	t.Run("IsEmpty empty", func(t *testing.T) {
		nullString := String{String: "", Valid: true}
		if !nullString.IsEmpty() {
			t.Fatal("it has to be empty")
		}
	})

	t.Run("IsEmpty null", func(t *testing.T) {
		nullString := String{String: "", Valid: false}
		if !nullString.IsEmpty() {
			t.Fatal("it has to be empty")
		}
	})
}
