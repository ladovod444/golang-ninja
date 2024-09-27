package calendar

import "errors"

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 0 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 0 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (d Date) Day() int {
	return d.day
}

func (d Date) Month() int {
	return d.month
}

func (d Date) Year() int {
	return d.year
}

