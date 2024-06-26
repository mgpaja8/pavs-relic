package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewFirstName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"John", "John", false},
		{"J", "", true}, // Too short
		{"", "", true},  // Too short
		{"ThisFirstNameIsWayTooLongToBeValid", "", true}, // Too long
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			fn, err := NewFirstName(test.input)
			if (err != nil) != test.hasError {
				t.Errorf("expected error: %v, got: %v", test.hasError, err)
			}
			assert.Equal(t, test.expected, string(fn), "expected first name mismatch")
		})
	}
}

func Test_FirstName_String(t *testing.T) {
	tests := []struct {
		input    FirstName
		expected string
	}{
		{FirstName("John"), "John"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			assert.Equal(t, test.expected, test.input.String(), "expected string representation mismatch")
		})
	}
}
