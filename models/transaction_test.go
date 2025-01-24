package models

import (
	"testing"
	"time"
)

func TestToRow(t *testing.T) {
	tests := []struct {
		name     string
		input    Transaction
		expected []interface{}
	}{
		{
			name: "valid transaction",
			input: Transaction{
				Id:          "1",
				Description: "Coffee",
				Amount:      4.99,
				Category:    "Food",
				Date:        time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
			},
			expected: []interface{}{"1", "Coffee", 4.99, "Food", "2023-10-10"},
		},
		{
			name: "empty description",
			input: Transaction{
				Id:          "2",
				Description: "",
				Amount:      10.50,
				Category:    "Transport",
				Date:        time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC),
			},
			expected: []interface{}{"2", "", 10.50, "Transport", "2023-03-15"},
		},
		{
			name: "zero amount",
			input: Transaction{
				Id:          "3",
				Description: "Free gift",
				Amount:      0.0,
				Category:    "Miscellaneous",
				Date:        time.Date(2023, 5, 20, 0, 0, 0, 0, time.UTC),
			},
			expected: []interface{}{"3", "Free gift", 0.0, "Miscellaneous", "2023-05-20"},
		},
		{
			name: "negative amount",
			input: Transaction{
				Id:          "4",
				Description: "Refund",
				Amount:      -25.75,
				Category:    "Shopping",
				Date:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: []interface{}{"4", "Refund", -25.75, "Shopping", "2023-01-01"},
		},
		{
			name: "empty fields",
			input: Transaction{
				Id:          "",
				Description: "",
				Amount:      0.0,
				Category:    "",
				Date:        time.Time{},
			},
			expected: []interface{}{"", "", 0.0, "", "0001-01-01"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.ToRow()
			for i, v := range got {
				if v != tt.expected[i] {
					t.Errorf("ToRow() = %v, expected %v", got, tt.expected)
					break
				}
			}
		})
	}
}
