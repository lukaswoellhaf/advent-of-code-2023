package day1_test

import (
	day1 "advent-of-code-2023/day_1"
	"advent-of-code-2023/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	day1Part1Input, _ := tools.ReadFileAndTrimSpace("../inputs/day_1_part_1_example.txt")
	calibrationValueSum, calibrationErr := day1.FindCalibrationValueSum(day1Part1Input)
	assert.Equal(t, nil, calibrationErr, "calibrationErr should be nil")
	assert.Equal(t, 142, calibrationValueSum, "calibrationValueSum should be 142")
}

func TestDay1Part2(t *testing.T) {
	day1Part2Input, _ := tools.ReadFileAndTrimSpace("../inputs/day_1_part_2_example.txt")
	calibrationValueSum, calibrationErr := day1.FindCalibrationValueSum(day1Part2Input)
	assert.Equal(t, nil, calibrationErr, "calibrationErr should be nil")
	assert.Equal(t, 281, calibrationValueSum, "calibrationValueSum should be 281")
}
