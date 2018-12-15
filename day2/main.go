package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseEntries(scanner *bufio.Scanner) []string {
	boxIDs := []string{}
	for scanner.Scan() {
		boxID := scanner.Text()
		boxIDs = append(boxIDs, boxID)
	}
	return boxIDs
}

func countOccurrencies(boxID string) map[rune]int {
	occurrencies := map[rune]int{}
	for _, character := range boxID {
		if _, ok := occurrencies[character]; ok {
			occurrencies[character]++
		} else {
			occurrencies[character] = 1
		}
	}

	return occurrencies
}

func checkTwoAndThree(occurrencies map[rune]int) (int, int) {
	two := 0
	three := 0
	for _, count := range occurrencies {
		if count == 2 {
			two = 1
		}
		if count == 3 {
			three = 1
		}
	}
	return two, three
}

func partOne(boxIDs []string) int {
	twoCount := 0
	threeCount := 0
	for _, boxID := range boxIDs {
		two, three := checkTwoAndThree(countOccurrencies(boxID))
		twoCount = twoCount + two
		threeCount = threeCount + three
	}

	return twoCount * threeCount
}

func almostEqual(first, second string) (bool, string) {
	found := false
	common := ""
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			if found == false {
				common = string(append([]rune(second)[:i], []rune(second)[i+1:]...))
				found = true
			} else {
				found = false
				break
			}
		}
	}
	return found, string(common)
}

func partTwo(boxIDs []string) string {

	found := false
	result := ""

	for _, first := range boxIDs {
		for _, second := range boxIDs {
			found, result = almostEqual(first, second)
			if found == true {
				return result
			}
		}
	}
	return ""
}

func main() {

	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	boxIDs := parseEntries(scanner)

	checksum := partOne(boxIDs)

	fmt.Println(checksum)

	common := partTwo(boxIDs)

	fmt.Println(common)

}
