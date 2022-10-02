package null

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

// NewBool creates a new Bool
func NewBool(b bool, valid bool) Bool {
	return Bool{Bool: b, Valid: valid}
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
	case uint8:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case uint16:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case uint32:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case uint64:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case int:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case int8:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case int16:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case int32:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
		b.Bool = data == 1
		return nil
	case int64:
		if data != 0 && data != 1 {
			return fmt.Errorf("unsupported bool value: %d", value)
		}
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
	return jsonMarshal(b.Bool)
}

// UnmarshalJSON decode data to the value.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var bb *bool
	if err := json.Unmarshal(data, &bb); err != nil {
		return err
	}
	b.Valid = bb != nil
	if b.Valid {
		b.Bool = *bb
	} else {
		b.Bool = false
	}
	return nil
}

// IsNull returns true if Valid is false.
func (b *Bool) IsNull() bool {
	return !b.Valid
}
