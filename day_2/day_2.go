package day2

import (
	"advent-of-code-2023/tools"
	"regexp"
	"strconv"
	"strings"
)

//--- Day 2: Cube Conundrum ---

// --- Part One ---
// Determine which games would have been possible if the bag had been loaded with only
// 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green --> Possible
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue --> Possible
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red --> NOT possible: 20 red in round 1
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red --> NOT possible: 14 blue in round 3
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green --> Possible

// Possible Game ID Sums = Game 1 + Game 2 + Game 5 = 8

// For each line:
// 1.1 Read/Remove Game <ID> from string
// 1.2 Split into game rounds by ';'
// 1.3 If feedback from each round possible (2.1) add game id to list
// 1.4 Sum all possible game ids

// For each round:
// 2.1 Determine possibility by limits

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

var (
	regexGameIdentifier = regexp.MustCompile(`^Game \d+:`)
	regexCubeColor      = regexp.MustCompile(`\b(?:red|green|blue)\b`)
	regexCubeValue      = regexp.MustCompile(`\d+`)
)

func GetSumFromIDsByPossibleGames(input string) (int, error) {
	var sumIDs int

	games := tools.SplitInputByNewLine(input)
	for gameID, game := range games {
		gameWithoutIdentifier := removeGameIdentifier(game)
		rounds := getAllRounds(gameWithoutIdentifier)
		for index, round := range rounds {
			isPossible, err := isRoundPossible(round)
			if err != nil {
				return -1, err
			}

			if !isPossible {
				break
			}

			if isPossible && index == len(rounds)-1 {
				sumIDs += gameID + 1
			}
		}
	}
	return sumIDs, nil
}

func removeGameIdentifier(game string) string {
	return regexGameIdentifier.ReplaceAllString(game, "")
}

func getAllRounds(gameWithoutIdentifier string) []string {
	return tools.SplitInputBySemicolon(gameWithoutIdentifier)
}

func isRoundPossible(round string) (bool, error) {
	cubes := getAllCubes(round)
	for _, cube := range cubes {
		cubeWithoutWhiteSpace := strings.TrimSpace(cube)
		cubeValue, getCubeErr := getCubeValue(cubeWithoutWhiteSpace)
		if getCubeErr != nil {
			return false, getCubeErr
		}
		cubeColor := getCubeColor(cubeWithoutWhiteSpace)

		switch cubeColor {
		case "red":
			if cubeValue > maxRed {
				return false, nil
			}
		case "green":
			if cubeValue > maxGreen {
				return false, nil
			}
		case "blue":
			if cubeValue > maxBlue {
				return false, nil
			}
		}
	}
	return true, nil
}

func getAllCubes(round string) []string {
	return tools.SplitInputByComma(round)
}

func getCubeColor(cube string) string {
	return regexCubeColor.FindString(cube)
}

func getCubeValue(cube string) (int, error) {
	cubeValueMatch := regexCubeValue.FindString(cube)
	return strconv.Atoi(cubeValueMatch)
}
