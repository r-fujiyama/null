package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// Int16 represents a int16 that may be null.
type Int16 struct {
	Int16 int16
	Valid bool
}

// NewInt16 creates a new Int16
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
		i16, err := strconv.ParseInt(string(data), 10, 16)
		if err != nil {
			return err
		}
		i.Int16 = int16(i16)
		return nil
	case int:
		if data > math.MaxInt16 || data < math.MinInt16 {
			return fmt.Errorf("maximum or minimum value of Int16 exceeded: %d", data)
		}
		i.Int16 = int16(data)
		return nil
	case int8:
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
	return int64(i.Int16), nil
}

// MarshalJSON encode the value to JSON.
func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return jsonMarshal(i.Int16)
}

// UnmarshalJSON decode data to the value.
func (i *Int16) UnmarshalJSON(data []byte) error {
	var i16 *int16
	if err := json.Unmarshal(data, &i16); err != nil {
		return err
	}
	i.Valid = i16 != nil
	if i.Valid {
		i.Int16 = *i16
	} else {
		i.Int16 = 0
	}
	return nil
}

// IsNull returns true if Valid is false.
func (i *Int16) IsNull() bool {
	return !i.Valid
}
