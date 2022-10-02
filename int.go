package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Int represents a int that may be null.
type Int struct {
	Int   int
	Valid bool
}

// Scan implements the Scanner interface.
func (i *Int) Scan(value interface{}) error {
	if value == nil {
		i.Int, i.Valid = 0, false
		return nil
	}

	i.Valid = true
	switch data := value.(type) {
	case string:
		integer, err := strconv.Atoi(data)
		if err != nil {
			return err
		}
		i.Int = integer
		return nil
	case []byte:
		integer, err := strconv.Atoi(string(data))
		if err != nil {
			return err
		}
		i.Int = integer
		return nil
	case int:
		i.Int = data
		return nil
	case int8:
		i.Int = int(data)
		return nil
	case int16:
		i.Int = int(data)
		return nil
	case int32:
		i.Int = int(data)
		return nil
	case int64:
		i.Int = int(data)
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (i Int) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return int64(i.Int), nil
}

// MarshalJSON encode the value to JSON.
func (i Int) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return JSONMarshal(i.Int)
}

// UnmarshalJSON decode data to the value.
func (i *Int) UnmarshalJSON(data []byte) error {
	var integer *int
	if err := json.Unmarshal(data, &integer); err != nil {
		return err
	}
	i.Valid = integer != nil
	if i.Valid {
		i.Int = *integer
	} else {
		i.Int = 0
	}
	return nil
}

// IsNull returns true if Valid is false.
func (i *Int) IsNull() bool {
	return !i.Valid
}
