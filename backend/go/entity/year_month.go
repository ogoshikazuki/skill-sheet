package entity

type YearMonth interface {
	Year() int
	Month() int
}

type yearMonth struct {
	year  int
	month int
}

type InvalidYearMonthError struct{}

func (e InvalidYearMonthError) Error() string {
	return "invalid year month"
}

func NewYearMonth(year, month int) (YearMonth, error) {
	yearMonth := yearMonth{
		year:  year,
		month: month,
	}
	if !yearMonth.isValid() {
		return nil, InvalidYearMonthError{}
	}

	return yearMonth, nil
}

func (y yearMonth) isZero() bool {
	return y.year == 0 && y.month == 0
}

func (y yearMonth) isValid() bool {
	if y.isZero() {
		return true
	}

	if y.month < 1 || y.month > 12 {
		return false
	}

	return true
}

func (y yearMonth) Year() int {
	return y.year
}

func (y yearMonth) Month() int {
	return y.month
}
