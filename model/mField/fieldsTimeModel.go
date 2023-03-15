package mField

import (
	"database/sql/driver"
	"time"
)

type FieldsTimeModel struct {
	CreatedAt LocalTime `OnlyRead:"true"`
	UpdatedAt LocalTime `OnlyRead:"true"`
}

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 || string(data) == "null" {
		*t = LocalTime(time.Time{})
		return
	}
	now, err := time.Parse(time.DateTime, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(time.DateTime)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(time.DateTime)
}

func (t LocalTime) Unix() int64 {
	return time.Time(t).Unix()
}
