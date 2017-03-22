package roman_test

import (
	"testing"

	"strconv"

	"../roman"
)

func TestRoman(t *testing.T) {
	number := 1
	romanNumeral := roman.DecodeToRomanNumeral(number)
	numberToString := strconv.Itoa(number)
	if romanNumeral != "I" {
		t.Error("Expected "+numberToString+" = I, got "+numberToString, romanNumeral)
	}
}
