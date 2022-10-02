package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// String represents a string that may be null.
type String struct {
	Str   string
	Valid bool
}

// Scan implements the Scanner interface.
func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.Str, s.Valid = "", false
		return nil
	}

	s.Valid = true
	switch data := value.(type) {
	case string:
		s.Str = data
		return nil
	case []byte:
		s.Str = string(data)
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
	return s.Str, nil
}

// MarshalJSON encode the value to JSON.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return JSONMarshal(s.Str)
}

// UnmarshalJSON decode data to the value.
func (s *String) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	s.Valid = str != nil
	if str == nil {
		s.Str = ""
	} else {
		s.Str = *str
	}
	return nil
}

// IsEmpty return true if String is "" or Valid is false.
func (s *String) IsEmpty() bool {
	return s.Str == "" || !s.Valid
}

// String return string indicated the value.
func (s String) String() string {
	if !s.Valid {
		return "<nil>"
	}
	return s.Str
}
