package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"unsafe"
)

// Int16 represents a int16 that may be null.
type Int16 struct {
	Int16 int16
	Valid bool
}

// NewInt16 creates an instance of Int16.
func NewInt16(i16 int16, valid bool) Int16 {
	return Int16{Int16: i16, Valid: valid}
}

// Scan implements the Scanner interface.
func (i *Int16) Scan(value interface{}) error {
	if value == nil {
		i.Int16, i.Valid = 0, false
		return nil
	}

	i.Valid = true
	switch data := value.(type) {
	case string:
		i16, err := strconv.ParseInt(data, 10, 16)
		if err != nil {
			return err
		}
		i.Int16 = int16(i16)
		return nil
	case []byte:
		i16, err := strconv.ParseInt(*(*string)(unsafe.Pointer(&data)), 10, 16)
		if err != nil {
			return err
		}
		i.Int16 = int16(i16)
		return nil
	case int:
		i.Int16 = int16(data)
		return nil
	case int16:
		i.Int16 = int16(data)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (i Int16) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int16, nil
}

// MarshalJSON encode the value to JSON.
func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int16)
}

// UnmarshalJSON decode data to the value.
func (i *Int16) UnmarshalJSON(data []byte) error {
	var fl *int16
	if err := json.Unmarshal(data, &fl); err != nil {
		return err
	}
	i.Valid = fl != nil
	if fl == nil {
		i.Int16 = 0
	} else {
		i.Int16 = *fl
	}
	return nil
}

// IsZeroOrNull return true if int16 is 0 or Valid is false.
func (i *Int16) IsZeroOrNull() bool {
	return i.Int16 == 0 || !i.Valid
}

// String return string indicated the value.
func (i Int16) String() string {
	if !i.Valid {
		return "<null>"
	}
	return strconv.Itoa(int(i.Int16))
}
