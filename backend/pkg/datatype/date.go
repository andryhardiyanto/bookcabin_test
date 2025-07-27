package datatype

import (
	"backend/pkg/fibers"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Date struct {
	Time  *time.Time
	Valid bool
}

func NewDate(t time.Time) Date {
	return Date{Time: &t, Valid: true}
}

func (d *Date) Scan(value any) error {
	if value == nil {
		d.Time = nil
		d.Valid = false
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		d.Time = &v
		d.Valid = true
		return nil
	case []byte:
		// for text format
		t, err := time.Parse(time.DateOnly, string(v))
		if err != nil {
			return err
		}
		d.Time = &t
		d.Valid = true
		return nil
	case string:
		t, err := time.Parse(time.DateOnly, v)
		if err != nil {
			return err
		}
		d.Time = &t
		d.Valid = true
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into Date", value)
	}
}

func (d Date) Value() (driver.Value, error) {
	if !d.Valid || d.Time == nil {
		return nil, nil
	}
	return *d.Time, nil
}
func (d Date) String() string {
	if !d.Valid || d.Time == nil {
		return ""
	}
	return d.Time.Format(time.DateOnly)
}

func (d Date) MarshalJSON() ([]byte, error) {
	if !d.Valid || d.Time == nil {
		return nil, nil
	}
	return json.Marshal(d.Time.Format(time.DateOnly))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "null" || str == "" {
		d.Time = nil
		d.Valid = false
		return nil
	}
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return fibers.ErrorInvalidDate
	}
	d.Time = &t
	d.Valid = true
	return nil
}

func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse(time.DateOnly, string(data))
	if err != nil {
		return fibers.ErrorInvalidDate
	}
	d.Time = &t
	d.Valid = true
	return nil
}
