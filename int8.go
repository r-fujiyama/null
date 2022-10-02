package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Int8 represents a int8 that may be null.
type Int8 struct {
	Int8  int8
	Valid bool
}

// Scan implements the Scanner interface.
func (i *Int8) Scan(value interface{}) error {
	if value == nil {
		i.Int8, i.Valid = 0, false
		return nil
	}

	i.Valid = true
	switch data := value.(type) {
	case string:
		i8, err := strconv.ParseInt(data, 10, 8)
		if err != nil {
			return err
		}
		i.Int8 = int8(i8)
		return nil
	case []byte:
		i8, err := strconv.ParseInt(string(data), 10, 8)
		if err != nil {
			return err
		}
		i.Int8 = int8(i8)
		return nil
	case int:
		i.Int8 = int8(data)
		return nil
	case int8:
		i.Int8 = int8(data)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (i Int8) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return int64(i.Int8), nil
}

// MarshalJSON encode the value to JSON.
func (i Int8) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return JSONMarshal(i.Int8)
}

// UnmarshalJSON decode data to the value.
func (i *Int8) UnmarshalJSON(data []byte) error {
	var i8 *int8
	if err := json.Unmarshal(data, &i8); err != nil {
		return err
	}
	i.Valid = i8 != nil
	if i.Valid {
		i.Int8 = *i8
	} else {
		i.Int8 = 0
	}
	return nil
}

// IsNull returns true if Valid is false.
func (i *Int8) IsNull() bool {
	return !i.Valid
}
