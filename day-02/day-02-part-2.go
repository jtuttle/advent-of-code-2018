package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func commonLetters(one string, two string) string {
	var common []rune

	oneRunes := []rune(one)
	twoRunes := []rune(two)
	
	for i := range one {
		if oneRunes[i] == twoRunes[i] {
			common = append(common, oneRunes[i])
		}
	}
	
	return string(common)
}

func main() {
	lines, _ := readLines("./day-02-input.txt")

	for i := range lines {
		for j := i + 1; j < len(lines); j++ {
			if i != j {
				common := commonLetters(lines[i], lines[j])
				
				if len(common) == len(lines[i]) - 1 {
					fmt.Println("Common letters:", common)
				}
			}
		}
	}
	
}
