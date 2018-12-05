package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func removeReactions(polymer string) string {
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

func main() {
	file, _ := os.Open("./day-05-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	polymer := scanner.Text()
	oldPolymer := ""

	for polymer != oldPolymer {
		oldPolymer = polymer
		polymer = removeReactions(polymer)
	}

	fmt.Println("Remaining units:", len(polymer))
}
