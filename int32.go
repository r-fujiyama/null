package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Int32 represents a int32 that may be null.
type Int32 struct {
	Int32 int32
	Valid bool
}

// NewInt32 creates a new Int32
func NewInt32(i32 int32, valid bool) Int32 {
	return Int32{Int32: i32, Valid: valid}
}

// Scan implements the Scanner interface.
func (i *Int32) Scan(value interface{}) error {
	if value == nil {
		i.Int32, i.Valid = 0, false
		return nil
	}

	i.Valid = true
	switch data := value.(type) {
	case string:
		i32, err := strconv.ParseInt(data, 10, 32)
		if err != nil {
			return err
		}
		i.Int32 = int32(i32)
		return nil
	case []byte:
		i32, err := strconv.ParseInt(string(data), 10, 32)
		if err != nil {
			return err
		}
		i.Int32 = int32(i32)
		return nil
	case int:
		i.Int32 = int32(data)
		return nil
	case int8:
		i.Int32 = int32(data)
		return nil
	case int16:
		i.Int32 = int32(data)
		return nil
	case int32:
		i.Int32 = data
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (i Int32) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return int64(i.Int32), nil
}

// MarshalJSON encode the value to JSON.
func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return jsonMarshal(i.Int32)
}

// UnmarshalJSON decode data to the value.
func (i *Int32) UnmarshalJSON(data []byte) error {
	var i32 *int32
	if err := json.Unmarshal(data, &i32); err != nil {
		return err
	}
	i.Valid = i32 != nil
	if i.Valid {
		i.Int32 = *i32
	} else {
		i.Int32 = 0
	}
	return nil
}

// IsNull returns true if Valid is false.
func (i *Int32) IsNull() bool {
	return !i.Valid
}
