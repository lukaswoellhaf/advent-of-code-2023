package day2_test

import (
	day2 "advent-of-code-2023/day_2"
	"advent-of-code-2023/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Part1(t *testing.T) {
	day2Part1Input, _ := tools.ReadFileAndTrimSpace("../inputs/day_2_part_1_example.txt")
	possibleGameIDsSum, getSumErr := day2.GetSumFromIDsByPossibleGames(day2Part1Input)
	assert.Equal(t, nil, getSumErr, "getSumErr should be nil")
	assert.Equal(t, 8, possibleGameIDsSum, "possibleGameIDsSum should be 8")
}
