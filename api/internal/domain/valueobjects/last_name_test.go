package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLastName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"Doe", "Doe", false},
		{"D", "", true}, // Too short
		{"", "", true},  // Too short
		{"ThisLastNameIsWayTooLongToBeValid", "", true}, // Too long
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ln, err := NewLastName(test.input)
			if (err != nil) != test.hasError {
				t.Errorf("expected error: %v, got: %v", test.hasError, err)
			}
			assert.Equal(t, test.expected, string(ln), "expected last name mismatch")
		})
	}
}

func Test_LastName_String(t *testing.T) {
	tests := []struct {
		input    LastName
		expected string
	}{
		{LastName("Doe"), "Doe"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			assert.Equal(t, test.expected, test.input.String(), "expected string representation mismatch")
		})
	}
}
