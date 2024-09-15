package formatter

import "testing"

func TestCheckStartingDigit(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"a", false},
		{"1", true},
		{"_", false},
		{"a1", false},
		{"1a", true},
	}

	for _, test := range tests {
		result := CheckStartingDigit(test.input)
		if result != test.expected {
			t.Errorf("Expected CheckStartingDigit(%s) to be %t, got %t", test.input, test.expected, result)
		}
	}
}
