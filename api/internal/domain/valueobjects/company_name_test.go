package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCompanyName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"Valid Company", "Valid Company", false},
		{"C", "", true}, // Too short
		{"", "", true},  // Too short
		{"ThisCompanyNameIsWayWayTooLongToBeValidAndShouldReturnAnError", "", true}, // Too long
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			cn, err := NewCompanyName(test.input)
			if (err != nil) != test.hasError {
				t.Errorf("expected error: %v, got: %v", test.hasError, err)
			}
			assert.Equal(t, test.expected, string(cn), "expected company name mismatch")
		})
	}
}

func Test_CompanyName_String(t *testing.T) {
	tests := []struct {
		input    CompanyName
		expected string
	}{
		{CompanyName("Valid Company"), "Valid Company"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			assert.Equal(t, test.expected, test.input.String(), "expected string representation mismatch")
		})
	}
}
