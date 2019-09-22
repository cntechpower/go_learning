package main

import "fmt"

type MyDate struct {
	Year  int
	Month month
	Day   int
}

type month int

const (
	January month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func NewMyDate(y int, m month, d int) *MyDate {
	return &MyDate{
		Year:  y,
		Month: m,
		Day:   d,
	}
}

func (d MyDate) IsLeapYear() bool {
	return d.Year%4 == 0
}

func (d MyDate) GetTotalDays() int {
	if d.IsLeapYear() {
		return 366
	}
	return 365
}

func (d MyDate) getMonthDays(m month) int {
	switch m {
	case February:
		if d.IsLeapYear() {
			return 29
		}
		return 28
	case January, March, May, July, August, October, December:
		return 31
	default:
		return 30
	}
}

func (d MyDate) GetExpireDays() int {
	days := 0
	for m := January; m < d.Month; m++ {
		days = days + d.getMonthDays(m)
	}
	return days + d.Day

}

func (d MyDate) GetRemainDays() int {
	days := 0
	var m month
	for m = December; d.Month < m; m-- {
		days = days + d.getMonthDays(m)
	}
	return days + d.getMonthDays(m) - d.Day
}

func (d MyDate) CalculatorDaysBetweenDate(d1 *MyDate) int {
	days := 0
	days = days + d.GetRemainDays()
	for y := d.Year + 1; y < d1.Year; y++ {
		days = days + NewMyDate(y, 0, 0).GetTotalDays()
	}
	return days + d1.GetExpireDays()
}

func main() {
	d1 := NewMyDate(1999, May, 10)
	d2 := NewMyDate(2006, March, 8)
	fmt.Println(d1.CalculatorDaysBetweenDate(d2))
}
