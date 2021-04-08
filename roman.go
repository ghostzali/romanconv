package romanconv

import (
	"regexp"
	"strings"
)

type RomanNumeralErr string

func (e RomanNumeralErr) Error() string {
	return string(e)
}

const (
	PATTERN      = "^(_M){0,3}(_C_M|_C_D|(_D)?(_C){0,3})(_X_C|_X_L|(_L)?(_X){0,3})(_I_X|_I_V|(_V)?(_I){0,3}|(_X)?M{0,3}|(_V)?M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"
	LIMIT_VALUE  = 3999999
	FORMAT_ERROR = RomanNumeralErr("invalid roman format")
	VALUE_ERROR  = RomanNumeralErr("value over-over limit. 1 - 3999999")
)

var (
	romanMapValues = map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"_I": 1000,
		"_V": 5000,
		"_X": 10000,
		"_L": 50000,
		"_C": 100000,
		"_D": 500000,
		"_M": 1000000,
	}
	valueMapRomans = map[int]string{
		1000000: "_M",
		900000:  "_C_M",
		500000:  "_D",
		400000:  "_C_D",
		100000:  "_C",
		90000:   "_X_C",
		50000:   "_L",
		40000:   "_X_L",
		10000:   "_X",
		9000:    "_I_X",
		5000:    "_V",
		4000:    "_I_V",
		1000:    "M",
		900:     "CM",
		500:     "D",
		400:     "CD",
		100:     "C",
		90:      "XC",
		50:      "L",
		40:      "XL",
		10:      "X",
		9:       "IX",
		5:       "V",
		4:       "IV",
		1:       "I",
	}
	valueIndex = []int{1000000, 900000, 500000, 400000, 100000, 90000, 50000, 40000, 10000, 9000, 5000, 4000, 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

// Convert takes roman numeral into arabic number
//
// Roman Characters : "I", "V", "X", "L", "C", "D", "M", "_I", "_V", "_X", "_L", "_C", "_D", "_M"
// Roman Numerals Chart reference: https://www.calculatorsoup.com/calculators/conversions/roman-numeral-converter.php
func Convert(roman string) (int, error) {
	roman = strings.ToUpper(roman)
	if !Validate(roman) {
		return 0, FORMAT_ERROR
	}
	romanValue := 0
	latestValue := 0
	for i := 0; i < len(roman); i++ {
		key := roman[i : i+1]
		if key == "_" {
			key = roman[i : i+2]
			i++
		}
		value := romanMapValues[key]
		if latestValue != 0 && latestValue < value {
			romanValue += value - (2 * latestValue)
		} else {
			romanValue += value
		}
		latestValue = value
	}
	if romanValue > LIMIT_VALUE {
		return romanValue, VALUE_ERROR
	}
	return romanValue, nil
}

// Parse format arabic number into roman numeral
//
// Limit input : 3999999
// Roman Numerals Chart reference: https://www.calculatorsoup.com/calculators/conversions/roman-numeral-converter.php
func Parse(arabic int) (string, error) {
	if arabic > LIMIT_VALUE {
		return "", VALUE_ERROR
	}
	value := arabic
	roman := ""
	for value > 0 {
		key := 0
		for _, k := range valueIndex {
			if k <= value {
				key = k
				break
			}
		}
		roman += valueMapRomans[key]
		value -= key
	}

	return roman, nil
}

// Validate check roman format string
func Validate(roman string) bool {
	romanValidPattern, _ := regexp.Compile(PATTERN)
	return romanValidPattern.MatchString(roman)
}
