package day1

import (
	"advent-of-code-2023/tools"
	"errors"
	"regexp"
	"strconv"
)

// --- Day 1: Trebuchet?! ---

// --- Part One ---
// The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value
// that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last
// digit (in that order) to form a single two-digit number.

// 1.0 Split input by newline
// For each line:
// 2.0 Find first digit in input line
// 2.1 Find last digit in input line
// 2.2 Concatenate first digit + last digit
// 2.3 Store combined digit
// 3.0 Sum all stored digits

// --- Part Two ---
// Your calculation isn't quite right. It looks like some of the digits
// are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".
// Equipped with this new information, you now need to find the real first and last digit on each line.

// 1. Adapt regular expression to match digits 0-9 and words one, two, three, four, five, six, seven, eight, and nine
// 2. Add conversion for number words into integers

func FindCalibrationValueSum(input string) (int, error) {
	var calibrationValuesSum int
	allLines := tools.SplitInputByNewLine(input)
	for _, line := range allLines {
		calibrationValue, err := findFirstAndLastDigitCombined(line)
		if err != nil {
			return -1, err
		}
		calibrationValuesSum += calibrationValue
	}

	return calibrationValuesSum, nil
}

func findFirstAndLastDigitCombined(line string) (int, error) {
	regularExp := regexp.MustCompile(`(?:[0-9]|one|two|three|four|five|six|seven|eight|nine)`)
	matches := regularExp.FindAllString(line, -1)

	digit, combineErr := combineDigits(matches)
	if combineErr != nil {
		return -1, combineErr
	}

	return digit, nil
}

func mapNumberMatchToDigit(numberMatch string) string {
	var digit string

	if len(numberMatch) > 1 {
		switch numberMatch {
		case "one":
			digit = "1"
		case "two":
			digit = "2"
		case "three":
			digit = "3"
		case "four":
			digit = "4"
		case "five":
			digit = "5"
		case "six":
			digit = "6"
		case "seven":
			digit = "7"
		case "eight":
			digit = "8"
		case "nine":
			digit = "9"
		}
	} else {
		digit = numberMatch
	}

	return digit
}

func combineDigits(numberMatches []string) (int, error) {
	if len(numberMatches) == 0 {
		return -1, errors.New("numberMatches slice is empty")
	}

	combinedDigitAsInt, convErr := strconv.Atoi(mapNumberMatchToDigit(numberMatches[0]) + mapNumberMatchToDigit(numberMatches[len(numberMatches)-1]))
	if convErr != nil {
		return -1, convErr
	}

	return combinedDigitAsInt, nil
}
