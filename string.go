package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// String represents a string that may be null.
type String struct {
	String string
	Valid  bool
}

// NewString creates a new String
func NewString(str string, valid bool) String {
	return String{String: str, Valid: valid}
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
		s.String = string(data)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (s String) Value() (driver.Value, error) {
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
	return jsonMarshal(s.String)
}

// UnmarshalJSON decode data to the value.
func (s *String) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	s.Valid = str != nil
	if s.Valid {
		s.String = *str
	} else {
		s.String = ""
	}
	return nil
}

// IsEmpty return true if String is "" or Valid is false.
func (s *String) IsEmpty() bool {
	return s.String == "" || !s.Valid
}

// IsNull returns true if Valid is false.
func (s *String) IsNull() bool {
	return !s.Valid
}
