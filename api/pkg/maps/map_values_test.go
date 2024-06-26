package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapValues(t *testing.T) {
	emptyMap := make(map[int]string)
	emptyResult := MapValues(emptyMap)
	assert.ElementsMatch(t, []string{}, emptyResult, "expected empty result")

	nonEmptyMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	expectedResult := []string{"one", "two", "three"}
	result := MapValues(nonEmptyMap)

	assert.ElementsMatch(t, expectedResult, result, "expected values mismatch")
}
