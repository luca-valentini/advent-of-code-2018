package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseEntries(scanner *bufio.Scanner) ([]int, error) {
	frequencyChanges := []int{}
	for scanner.Scan() {
		change, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return frequencyChanges, err
		}
		frequencyChanges = append(frequencyChanges, change)
	}
	return frequencyChanges, nil
}

func partOne(frequencyChanges []int) (int, error) {
	resultingFrequency := 0

	for _, change := range frequencyChanges {
		resultingFrequency += change
	}

	return resultingFrequency, nil
}

func partTwo(frequencyChanges []int) (int, error) {
	resultingFrequency := 0
	frequencies := map[int]bool{resultingFrequency: true}

	for {
		for _, change := range frequencyChanges {
			resultingFrequency += change
			if frequencies[resultingFrequency] {
				return resultingFrequency, nil
			}

			frequencies[resultingFrequency] = true
		}
	}
}

func main() {

	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	frequencyChanges, err := parseEntries(scanner)

	if err != nil {
		panic(err)
	}

	result, err := partOne(frequencyChanges)
	fmt.Println("[Part one] Result is:", result)

	result, err = partTwo(frequencyChanges)
	fmt.Println("[Part two] Result is:", result)

}
