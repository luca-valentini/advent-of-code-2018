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

func main() {

	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)

	twoCount := 0
	threeCount := 0
	for _, boxID := range parseEntries(scanner) {
		two, three := checkTwoAndThree(countOccurrencies(boxID))
		twoCount = twoCount + two
		threeCount = threeCount + three
	}

	checksum := twoCount * threeCount
	fmt.Println(checksum)
}
