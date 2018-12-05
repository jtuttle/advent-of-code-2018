package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func reactOnce(polymer string) string {
	newPolymer := make([]rune, 0)
	polymerUnits := []rune(polymer)

	skipNext := false
	
	for i, unit := range polymerUnits {
		if i < len(polymer) - 1 {
			// Skip the matched unit if we found a match last loop
			if skipNext {
				skipNext = false
				continue
			}
			
			nextUnit := polymerUnits[i+1]

			runeDiff := float64(unit) - float64(nextUnit)
			
			if math.Abs(runeDiff) == 32 {
				skipNext = true
			} else {
				newPolymer = append(newPolymer, unit)
			}
		} else {
			// Append the final letter unless it was matched on last loop
			if !skipNext {
				newPolymer = append(newPolymer, unit)
			}
		}
	}

	return string(newPolymer)
}

func reactUntilStable(polymer string) string {
	oldPolymer := ""
	
	for polymer != oldPolymer {
		oldPolymer = polymer
		polymer = reactOnce(polymer)
	}

	return polymer
}

func findShortestSubbedLength(polymer string) int {
	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	minLength := len(polymer)
	
	for _, letter := range alphabet {
		polymerSubbed := strings.Replace(polymer, string(letter), "", -1)
		polymerSubbed = strings.Replace(polymerSubbed, string(unicode.ToLower(letter)), "", -1)

		polymerSubbed = reactUntilStable(polymerSubbed)

		length := len(polymerSubbed)

		if length < minLength {
			minLength = length
		}
	}

	return minLength
}

func main() {
	file, _ := os.Open("./day-05-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	polymer := scanner.Text()

	shortestSubbedLength := findShortestSubbedLength(polymer)
	
	fmt.Println("Minimum remaining units:", shortestSubbedLength)
}
