package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLetters(line string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range []rune(line) {
		counts[char] += 1
	}

	return counts
}

func hasNLetters(counts map[rune]int, n int) (bool) {
	for _, count := range counts {
		if count == n {
			return true
		}
	}

	return false
}

func main() {
	file, _ := os.Open("./day-02-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	twoCount := 0
	threeCount := 0
	
	for scanner.Scan() {
		line := scanner.Text()

		counts := countLetters(line)

		if hasNLetters(counts, 2) {
			twoCount++
		}

		if hasNLetters(counts, 3) {
			threeCount++
		}
	}

	fmt.Println("Checksum:", twoCount * threeCount)
}
