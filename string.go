package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"unsafe"
)

// String represents a string that may be null.
type String struct {
	String string
	Valid  bool
}

// NewString creates an instance of String.
func NewString(str string) String {
	return String{String: str, Valid: str != ""}
}

// Scan implements the Scanner interface.
func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.String, s.Valid = "", false
		return nil
	}

	s.Valid = true
	switch data := value.(type) {
	case string:
		s.String = data
		return nil
	case []byte:
		s.String = *(*string)(unsafe.Pointer(&data))
		return nil
	default:
		return fmt.Errorf("got data of type %T", value)
	}
}

// Value implements the driver Valuer interface.
func (s *String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.String, nil
}

// MarshalJSON encode the value to JSON.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

// UnmarshalJSON decode data to the value.
func (s *String) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	s.Valid = str != nil
	if str == nil {
		s.String = ""
	} else {
		s.String = *str
	}
	return nil
}

// IsEmpty returns true if String is "" or Valid is false.
func (s *String) IsEmpty() bool {
	return s.String == "" || !s.Valid
}
