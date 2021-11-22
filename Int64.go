package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Int64 represents a int64 that may be null.
type Int64 struct {
	Int64 int64
	Valid bool
}

// Scan implements the Scanner interface.
func (i *Int64) Scan(value interface{}) error {
	if value == nil {
		i.Int64, i.Valid = 0, false
		return nil
	}

	i.Valid = true
	switch data := value.(type) {
	case string:
		i64, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return err
		}
		i.Int64 = i64
		return nil
	case []byte:
		i64, err := strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return err
		}
		i.Int64 = i64
		return nil
	case int:
		i.Int64 = int64(data)
		return nil
	case int16:
		i.Int64 = int64(data)
		return nil
	case int32:
		i.Int64 = int64(data)
		return nil
	case int64:
		i.Int64 = data
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int64, nil
}

// MarshalJSON encode the value to JSON.
func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int64)
}

// UnmarshalJSON decode data to the value.
func (i *Int64) UnmarshalJSON(data []byte) error {
	var fl *int64
	if err := json.Unmarshal(data, &fl); err != nil {
		return err
	}
	i.Valid = fl != nil
	if fl == nil {
		i.Int64 = 0
	} else {
		i.Int64 = *fl
	}
	return nil
}

// IsZeroOrNull return true if int64 is 0 or Valid is false.
func (i *Int64) IsZeroOrNull() bool {
	return i.Int64 == 0 || !i.Valid
}

// String return string indicated the value.
func (i Int64) String() string {
	if !i.Valid {
		return "<null>"
	}
	return strconv.Itoa(int(i.Int64))
}
