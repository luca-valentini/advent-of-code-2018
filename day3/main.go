package main

import (
	"bufio"
	"fmt"
	"os"
)

type matrix [1000][1000]int

type entry struct {
	x      int
	y      int
	width  int
	height int
}

func parseEntries(scanner *bufio.Scanner) map[int]entry {
	entries := map[int]entry{}
	for scanner.Scan() {
		var id int
		var entry entry
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &entry.x, &entry.y, &entry.width, &entry.height)
		entries[id] = entry
	}
	return entries
}

func partOne(entries map[int]entry) int {

	var land matrix
	counter := 0

	for _, entry := range entries {
		for x := 0; x < entry.width; x++ {
			for y := 0; y < entry.height; y++ {
				if land[x+entry.x][y+entry.y] == 1 {
					counter++
				}
				land[x+entry.x][y+entry.y]++
			}
		}

	}
	return counter
}

func partTwo(entries map[int]entry) int {

	var land matrix
	ids := map[int]bool{}

	for id, entry := range entries {
		append := true
		for x := 0; x < entry.width; x++ {
			for y := 0; y < entry.height; y++ {
				if land[x+entry.x][y+entry.y] != 0 {
					append = false
					delete(ids, land[x+entry.x][y+entry.y])
				}
				land[x+entry.x][y+entry.y] = id
			}
		}
		if append == true {
			ids[id] = true
		}

	}

	result := 0
	for id := range ids {
		result = id
	}
	return result
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	entries := parseEntries(scanner)

	count := partOne(entries)
	fmt.Println(count)

	id := partTwo(entries)
	fmt.Println(id)
}
