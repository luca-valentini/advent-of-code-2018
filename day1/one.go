package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	resultingFrequency := 0

	for scanner.Scan() {
		frequencyChange, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		resultingFrequency += frequencyChange
	}

	fmt.Println("Sum:", resultingFrequency)
}
