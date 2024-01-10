package entity_test

import (
	"testing"
	"time"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

func TestDateString(t *testing.T) {
	date := entity.Date("1991-07-01")

	if date.String() != "1991-07-01" {
		t.Error()
	}
}

func TestDateIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid",
			input:    "1991-07-01",
			expected: true,
		},
		{
			name:     "invalid",
			input:    "not date",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if entity.Date(tt.input).IsValid() != tt.expected {
				t.Errorf("input: %s, expected: %t", tt.input, tt.expected)
			}
		})
	}
}

func TestNewDateFromTime(t *testing.T) {
	time, _ := time.Parse("2006-01-02", "1991-07-01")
	date := entity.NewDateFromTime(time)
	if date != "1991-07-01" {
		t.Error()
	}
}
