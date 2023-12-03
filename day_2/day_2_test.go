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

func TestDay2Part2(t *testing.T) {
	day2Part2Input, _ := tools.ReadFileAndTrimSpace("../inputs/day_2_part_2_example.txt")
	sumPowerOfMinCubes, sumPowerErr := day2.SumPowerOfMinCubesNeededForGames(day2Part2Input)
	assert.Equal(t, nil, sumPowerErr, "sumPowerErr should be nil")
	assert.Equal(t, 2286, sumPowerOfMinCubes, "sumPowerOfMinCubes should be 2286")
}
