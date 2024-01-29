package entity_test

import (
	"testing"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

func TestYearMonth(t *testing.T) {
	tests := map[string]struct {
		year        int
		month       int
		expectedErr error
	}{
		"Normal": {
			year:  1991,
			month: 7,
		},
		"ZeroValue": {
			year:  0,
			month: 0,
		},
		"InvalidMonth": {
			year:        1991,
			month:       0,
			expectedErr: entity.InvalidYearMonthError{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			yearMonth, err := entity.NewYearMonth(tt.year, tt.month)
			if err != tt.expectedErr {
				t.Errorf("tt.expectedErr: %+v, err: %+v", tt.expectedErr, err)
			}
			if tt.expectedErr != nil {
				return
			}
			if yearMonth.Year() != tt.year {
				t.Errorf("tt.year: %d, yearMonth.Year(): %d", tt.year, yearMonth.Year())
			}
			if yearMonth.Month() != tt.month {
				t.Errorf("tt.month: %d, yearMonth.Month(): %d", tt.month, yearMonth.Month())
			}
		})
	}
}
