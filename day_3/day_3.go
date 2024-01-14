package day3

import (
	"advent-of-code-2023/tools"
	"regexp"
	"strconv"
	"strings"
)

//--- Day 3: Gear Ratios ---
// The engine schematic (your puzzle input) consists of a visual representation of the engine.
// There are lots of numbers and symbols you don't really understand, but apparently any number adjacent
// to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not
// count as a symbol.)

//  What is the sum of all of the part numbers in the engine schematic?

// Parse input into 2D array
// For each x y in 2D array:
// 2.1 Find digit
// 2.2 If digit found: look for neighbour digits
// 2.3 Save coordinates of found digits
// 2.4 Create hit box for found digits of specific length
// For each digit in hitbox:
// 2.4.1 Hitbox must search for matches: (x0 y+1; x0 y-1; x+1 y0; x-1 y0)
// 2.4.2 Save Part Number as soon as a symbol is found in a hitbox of a digit
// 2.5 Mark found Part Numbers as '.' to be not considered anymore

//2.6 Sum stored PartNumbers

func SumPartNumbersInEngineSemanticsBoard(input string) (int, error) {
	engineSemanticsBoard := initEngineSemanticsBoard(input)
	var sum int

	digitGroupCoordinates := getDigitGroupCoordinates(engineSemanticsBoard)
	partNumbers, getPartNumbersErr := getPartNumbers(engineSemanticsBoard, digitGroupCoordinates)
	if getPartNumbersErr != nil {
		return -1, getPartNumbersErr
	}

	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	return sum, nil
}

func initEngineSemanticsBoard(input string) [][]string {
	engineSemanticRows := tools.SplitInputByNewLine(input)
	engineSemanticsBoard := make([][]string, len(engineSemanticRows))

	for x := range engineSemanticsBoard {
		engineSemanticsBoard[x] = strings.Split(engineSemanticRows[x], "")
	}

	return engineSemanticsBoard
}

func getDigitGroupCoordinates(engineSemanticsBoard [][]string) map[int][][]int {
	digitGroupCoordinates := make(map[int][][]int, 0)
	var digitsOnly = regexp.MustCompile("[0-9]")
	digitGroupIndex := 0
	digitGroupWasPrevious := false

	for x := range engineSemanticsBoard {
		for y := range engineSemanticsBoard[x] {
			if isDigitFound := digitsOnly.MatchString(engineSemanticsBoard[x][y]); isDigitFound {
				coordinate := []int{x, y}
				digitGroupCoordinates[digitGroupIndex] = append(digitGroupCoordinates[digitGroupIndex], coordinate)
				digitGroupWasPrevious = true
			} else if digitGroupWasPrevious {
				digitGroupIndex++
				digitGroupWasPrevious = false
			}
		}
	}

	return digitGroupCoordinates
}

func getPartNumbers(engineSemanticsBoard [][]string, digitGroupCoordinates map[int][][]int) ([]int, error) {
	var partNumbersFound []int
	partNumbersOnly := regexp.MustCompile("[^0-9.]")

	for digitGroup, digitGroupCoordinate := range digitGroupCoordinates {

		adjacentCoordinates := getAdjacentCoordinates(digitGroupCoordinate, engineSemanticsBoard)
		for _, adjacentCoordinate := range adjacentCoordinates {
			x := adjacentCoordinate[0]
			y := adjacentCoordinate[1]

			if isPartNumber := partNumbersOnly.MatchString(engineSemanticsBoard[x][y]); isPartNumber {
				digitCoordinates := digitGroupCoordinates[digitGroup]
				var partNumberAsString string
				for _, digitCoordinate := range digitCoordinates {
					x := digitCoordinate[0]
					y := digitCoordinate[1]

					digitAsString := engineSemanticsBoard[x][y]
					partNumberAsString += digitAsString
				}

				partNumber, conversionErr := strconv.Atoi(partNumberAsString)
				if conversionErr != nil {
					return nil, conversionErr
				}

				partNumbersFound = append(partNumbersFound, partNumber)
				break
			}
		}
	}

	return partNumbersFound, nil
}

func getAdjacentCoordinates(coordinates [][]int, engineSemanticsBoard [][]string) [][]int {
	isValidCoordinate := func(x, y int) bool {
		return x >= 0 && x < len(engineSemanticsBoard) && y >= 0 && y < len(engineSemanticsBoard[x])
	}

	var adjacentCoordinates [][]int
	for _, coordinate := range coordinates {
		x := coordinate[0]
		y := coordinate[1]

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {

				if dx == 0 && dy == 0 {
					continue
				}

				adjacentX := x + dx
				adjacentY := y + dy

				if isValidCoordinate(adjacentX, adjacentY) {
					adjacentCoordinates = append(adjacentCoordinates, []int{adjacentX, adjacentY})
				}
			}
		}
	}
	return adjacentCoordinates
}
