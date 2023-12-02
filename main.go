package main

import (
	day1 "advent-of-code-2023/day_1"
	day2 "advent-of-code-2023/day_2"
	"advent-of-code-2023/tools"
	"log"
)

func main() {
	greeting, _ := tools.ReadFile("inputs/startup_message.txt")
	println(greeting)

	println("--- Day 1: Trebuchet?! Part 1 ---")
	day1Part1Input, _ := tools.ReadFileAndTrimSpace("inputs/day_1_part_1_puzzle.txt")
	calibrationValueSum, calibrationErr := day1.FindCalibrationValueSum(day1Part1Input)
	if calibrationErr != nil {
		log.Fatal(calibrationErr)
	}
	println("Result:", calibrationValueSum)
	println("--- Day 1: Trebuchet?! Part 2 ---")
	day1Part2Input, _ := tools.ReadFileAndTrimSpace("inputs/day_1_part_2_puzzle.txt")
	calibrationValueSum, calibrationErr = day1.FindCalibrationValueSum(day1Part2Input)
	if calibrationErr != nil {
		log.Fatal(calibrationErr)
	}
	println("Result:", calibrationValueSum)

	println("--- Day 2: Cube Conundrum Part 1 ---")
	day2Part1Input, _ := tools.ReadFileAndTrimSpace("inputs/day_2_part_1_puzzle.txt")
	possibleGameIDsSum, getSumErr := day2.GetSumFromIDsByPossibleGames(day2Part1Input)
	if getSumErr != nil {
		log.Fatal(getSumErr)
	}
	println("Result:", possibleGameIDsSum)
}
