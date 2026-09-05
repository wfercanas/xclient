package types

import (
	"encoding/json"
	"time"
)

type Date string

func NewDate(t time.Time) Date {
	return Date(t.Format("2006-01-02"))
}

func ToDate(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return "", err
	}

	return NewDate(t), nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(*d)
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*d = NewDate(t)
	return nil
}
