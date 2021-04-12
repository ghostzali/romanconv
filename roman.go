package romanconv

import (
	"regexp"
	"strings"
)

type RomanNumeralErr string
type RomanNumeral struct {
	Value  int
	Symbol string
}
type RomanNumerals []RomanNumeral

func (e RomanNumeralErr) Error() string {
	return string(e)
}

func (rs RomanNumerals) ValueOf(symbol string) int {
	for _, r := range rs {
		if r.Symbol == symbol {
			return r.Value
		}
	}
	return 0
}

func (rs RomanNumerals) Exists(symbol string) bool {
	for _, r := range rs {
		if r.Symbol == symbol {
			return true
		}
	}
	return false
}

const (
	PATTERN      = "^(_M){0,3}(_C_M|_C_D|(_D)?(_C){0,3})(_X_C|_X_L|(_L)?(_X){0,3})(_I_X|_I_V|(_V)?(_I){0,3}|(_X)?M{0,3}|(_V)?M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"
	LIMIT_VALUE  = 3999999
	FORMAT_ERROR = RomanNumeralErr("invalid roman format")
	VALUE_ERROR  = RomanNumeralErr("value over-over limit. 1 - 3999999")
)

var allRomanNumerals = RomanNumerals{
	{1000000, "_M"},
	{900000, "_C_M"},
	{500000, "_D"},
	{400000, "_C_D"},
	{100000, "_C"},
	{90000, "_X_C"},
	{50000, "_L"},
	{40000, "_X_L"},
	{10000, "_X"},
	{9000, "_I_X"},
	{5000, "_V"},
	{4000, "_I_V"},
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Convert takes roman numeral into arabic number
//
// Roman Characters : "I", "V", "X", "L", "C", "D", "M", "_I", "_V", "_X", "_L", "_C", "_D", "_M"
// Roman Numerals Chart reference: https://www.calculatorsoup.com/calculators/conversions/roman-numeral-converter.php
func Convert(roman string) (int, error) {
	roman = strings.ToUpper(roman)
	if !Validate(roman) {
		return 0, FORMAT_ERROR
	}
	total := 0
	for i := 0; i < len(roman); i++ {
		iLength := 1
		currentIndex := i
		if roman[currentIndex:currentIndex+iLength] == "_" {
			iLength = 2
			i++
		}
		symbol := roman[currentIndex : currentIndex+iLength]
		if currentIndex+(2*iLength)-1 < len(roman) && IsSubstractive(symbol) {
			potentialNumber := roman[currentIndex : currentIndex+(2*iLength)]
			value := allRomanNumerals.ValueOf(potentialNumber)
			if value != 0 {
				total += value
				i += iLength
				continue
			}
		}
		total += allRomanNumerals.ValueOf(symbol)
	}
	if total > LIMIT_VALUE {
		return total, VALUE_ERROR
	}
	return total, nil
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
	var roman strings.Builder

	for _, r := range allRomanNumerals {
		for value >= r.Value {
			roman.WriteString(r.Symbol)
			value -= r.Value
		}
	}

	return roman.String(), nil
}

// Validate check roman format string
func Validate(roman string) bool {
	romanValidPattern, _ := regexp.Compile(PATTERN)
	return romanValidPattern.MatchString(roman)
}

// IsSubstractive check if symbol can do substraction
func IsSubstractive(symbol string) bool {
	if symbol == "I" || symbol == "X" || symbol == "C" || symbol == "_I" || symbol == "_X" || symbol == "_C" {
		return true
	}
	return false
}
