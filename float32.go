package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Float32 represents a float32 that may be null.
type Float32 struct {
	Float32 float32
	Valid   bool
}

// Scan implements the Scanner interface.
func (f *Float32) Scan(value interface{}) error {
	if value == nil {
		f.Float32, f.Valid = 0, false
		return nil
	}

	f.Valid = true
	switch data := value.(type) {
	case string:
		f32, err := strconv.ParseFloat(data, 32)
		if err != nil {
			return err
		}
		f.Float32 = float32(f32)
		return nil
	case []byte:
		f32, err := strconv.ParseFloat(string(data), 32)
		if err != nil {
			return err
		}
		f.Float32 = float32(f32)
		return nil
	case int:
		f.Float32 = float32(data)
		return nil
	case int8:
		f.Float32 = float32(data)
		return nil
	case int16:
		f.Float32 = float32(data)
		return nil
	case int32:
		f.Float32 = float32(data)
		return nil
	case int64:
		f.Float32 = float32(data)
		return nil
	case float32:
		f.Float32 = data
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (f Float32) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float32, nil
}

// MarshalJSON encode the value to JSON.
func (f Float32) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return JSONMarshal(f.Float32)
}

// UnmarshalJSON decode data to the value.
func (f *Float32) UnmarshalJSON(data []byte) error {
	var f32 *float32
	if err := json.Unmarshal(data, &f32); err != nil {
		return err
	}
	f.Valid = f32 != nil
	if f.Valid {
		f.Float32 = *f32
	} else {
		f.Float32 = 0
	}
	return nil
}

// IsNull returns true if Valid is false.
func (f *Float32) IsNull() bool {
	return !f.Valid
}
