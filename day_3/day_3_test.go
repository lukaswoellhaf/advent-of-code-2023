package day3_test

import (
	day3 "advent-of-code-2023/day_3"
	"advent-of-code-2023/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3Part1(t *testing.T) {
	day3Part1Input, _ := tools.ReadFileAndTrimSpace("../inputs/day_3_part_1_example.txt")
	sumOfPartNumbers, getSumOfPartNUmbersErr := day3.SumPartNumbersInEngineSemanticsBoard(day3Part1Input)
	assert.Equal(t, nil, getSumOfPartNUmbersErr, "getSumOfPartNUmbersErr should be nil")
	assert.Equal(t, 4361, sumOfPartNumbers, "sumOfPartNumbers should be 4361")
}
