package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Bool represents a bool that may be null.
type Bool struct {
	Bool  bool
	Valid bool
}

// Scan implements the Scanner interface.
func (b *Bool) Scan(value interface{}) error {
	if value == nil {
		b.Bool, b.Valid = false, false
		return nil
	}

	b.Valid = true
	switch data := value.(type) {
	case string:
		toBool, err := strconv.ParseBool(data)
		if err != nil {
			return err
		}
		b.Bool = toBool
		return nil
	case []byte:
		toBool, err := strconv.ParseBool(string(data))
		if err != nil {
			return err
		}
		b.Bool = toBool
		return nil
	case int:
		b.Bool = data == 1
		return nil
	case int16:
		b.Bool = data == 1
		return nil
	case int32:
		b.Bool = data == 1
		return nil
	case int64:
		b.Bool = data == 1
		return nil
	case bool:
		b.Bool = data
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (b Bool) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}
	return b.Bool, nil
}

// MarshalJSON encode the value to JSON.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return JSONMarshal(b.Bool)
}

// UnmarshalJSON decode data to the value.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var bb *bool
	if err := json.Unmarshal(data, &bb); err != nil {
		return err
	}
	b.Valid = bb != nil
	if bb == nil {
		b.Bool = false
	} else {
		b.Bool = *bb
	}
	return nil
}
