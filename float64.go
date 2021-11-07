package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"unsafe"
)

// Float64 represents a float64 that may be null.
type Float64 struct {
	Float64 float64
	Valid   bool
}

// NewFloat64 creates an instance of Float64.
func NewFloat64(f float64, valid bool) Float64 {
	return Float64{Float64: f, Valid: valid}
}

// Scan implements the Scanner interface.
func (f *Float64) Scan(value interface{}) error {
	if value == nil {
		f.Float64, f.Valid = 0, false
		return nil
	}

	f.Valid = true
	switch data := value.(type) {
	case string:
		f64, err := strconv.ParseFloat(data, 64)
		if err != nil {
			return err
		}
		f.Float64 = f64
		return nil
	case []byte:
		f64, err := strconv.ParseFloat(*(*string)(unsafe.Pointer(&data)), 64)
		if err != nil {
			return err
		}
		f.Float64 = f64
		return nil
	case int:
		f.Float64 = float64(data)
		return nil
	case int16:
		f.Float64 = float64(data)
		return nil
	case int32:
		f.Float64 = float64(data)
		return nil
	case int64:
		f.Float64 = float64(data)
		return nil
	case float64:
		f.Float64 = data
		return nil
	default:
		return fmt.Errorf("got data of type %T", value)
	}
}

// Value implements the driver Valuer interface.
func (f *Float64) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float64, nil
}

// MarshalJSON encode the value to JSON.
func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.Float64)
}

// UnmarshalJSON decode data to the value.
func (f *Float64) UnmarshalJSON(data []byte) error {
	var f64 *float64
	if err := json.Unmarshal(data, &f64); err != nil {
		return err
	}
	f.Valid = f64 != nil
	if f64 == nil {
		f.Float64 = 0
	} else {
		f.Float64 = *f64
	}
	return nil
}

// IsZeroOrEmpty return true if float64 is 0 or Valid is false.
func (f *Float64) IsZeroOrNull() bool {
	return f.Float64 == 0 || !f.Valid
}
