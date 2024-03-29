package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// Int8 represents a int8 that may be null.
type Int8 struct {
	Int8  int8
	Valid bool
}

// NewInt8 creates a new Int8
func NewInt8(i8 int8, valid bool) Int8 {
	return Int8{Int8: i8, Valid: valid}
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
		if data > math.MaxInt8 || data < math.MinInt8 {
			return fmt.Errorf("maximum or minimum value of Int16 exceeded: %d", data)
		}
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
	return jsonMarshal(i.Int8)
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
