package types

import (
	"time"

	"github.com/quinntas/go-rest-template/internal/api/utils/timeUtils"
)

type Date struct {
	Value time.Time
}

func NewDate(value time.Time) Date {
	return Date{
		Value: value,
	}
}

func (d *Date) Set(value string) {
	d.Value, _ = timeUtils.StringToTime(value)
}

func (d *Date) ToString() string {
	return timeUtils.TimeToString(d.Value)
}

func (d *Date) ToPointerString() *string {
	str := d.ToString()
	if str == "0001-01-01 00:00:00" {
		return nil
	}
	return &str
}
