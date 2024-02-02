package day4_test

import (
	day4 "advent-of-code-2023/day_4"
	"advent-of-code-2023/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4Part(t *testing.T) {
	day4Part1Input, _ := tools.ReadFileAndTrimSpace("../inputs/day_4_part_1_example.txt")
	totalPoints := day4.GetTotalPointsOfScratchCards(day4Part1Input)
	assert.Equal(t, 13, totalPoints, "totalPoints should be 13")
}
