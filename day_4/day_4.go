package day4

import (
	"advent-of-code-2023/tools"
	"math"
	"regexp"
	"strings"
)

//--- Day 4: Scratchcards ---

// --- Part One ---
// The Elf leads you over to the pile of colorful cards. There, you discover dozens of scratchcards,
// all with their opaque covering already scratched off. Picking one up, it looks like each card has
// two lists of numbers separated by a vertical bar (|): a list of winning numbers and then a list of
// numbers you have. You organize the information into a table (your puzzle input).

// As far as the Elf has been able to figure out, you have to figure out which of the numbers you have
// appear in the list of winning numbers. The first match makes the card worth one point and each match
// after the first doubles the point value of that card.

// Take a seat in the large pile of colorful cards. How many points are they worth in total?

// For each card:
//	Split winning numbers and actual numbers by `|`
//	Split winning numbers and actual numbers by ` `
//	Find common elements between winning numbers and actual numbers
//	Calculate card points: 2 ** len(common element) // 2
// Sum the card points

func GetTotalPointsOfScratchCards(input string) int {
	scratchCards := tools.SplitInputByNewLine(input)
	var totalPoints int

	for _, element := range scratchCards {
		card := removeCardIdentifier(element)
		winningNumbers, actualNumbers := splitWinningAndActualNumbers(card)
		points := getPointsOfScratchCard(winningNumbers, actualNumbers)
		totalPoints += points
	}
	return totalPoints
}

func removeCardIdentifier(element string) string {
	regexCardIdentifier := regexp.MustCompile(`^Card \d+:`)
	return regexCardIdentifier.ReplaceAllString(element, "")
}

func splitWinningAndActualNumbers(card string) ([]string, []string) {
	winningAndActualNumbers := strings.Split(card, "|")
	winningNumbers := winningAndActualNumbers[0]
	actualNumbers := winningAndActualNumbers[1]
	return strings.Fields(winningNumbers), strings.Fields(actualNumbers)
}

func getPointsOfScratchCard(winningNumbers, actualNumbers []string) int {
	commonNumbers := tools.FindCommonElements(winningNumbers, actualNumbers)
	return int(math.Pow(2, float64(len(commonNumbers))) / 2)
}
