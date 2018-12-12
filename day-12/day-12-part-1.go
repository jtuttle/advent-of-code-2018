package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findPlantBounds(state map[int]bool) (int, int) {
	var first, last *int

	for k, v := range state {		
		if v {
			if first == nil || k < *first {
				newFirst := k
				first = &newFirst
			}

			if last == nil || k > *last {
				newLast := k
				last = &newLast
			}
		}
	}

	return *first, *last
}

func getLocalState(state map[int]bool, index int) string {
	localState := make([]rune, 0)
	
	for i := index - 2; i <= index + 2; i++ {
		if state[i] {
			localState = append(localState, '#')
		} else {
			localState = append(localState, '.')
		}
	}

	return string(localState)
}

func nextState(state map[int]bool, rules map[string]bool) map[int]bool {
	nextState := make(map[int]bool)
	
	firstPlant, lastPlant := findPlantBounds(state)

	for i := firstPlant - 2; i <= lastPlant + 2; i++ {
		substr := getLocalState(state, i)

		if val, ok := rules[substr]; ok {
			nextState[i] = val
		}
	}

	return nextState
}

func computePlantPotSum(state map[int]bool) int {
	sum := 0

	for k, v := range state {
		if v {
			sum += k
		}
	}

	return sum
}

func printState(state map[int]bool, firstIndex int, lastIndex int) {
	for i := firstIndex; i <= lastIndex; i++ {
		if state[i] {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	file, _ := os.Open("./day-12-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	firstLine := scanner.Text()

	stateStr := strings.Split(firstLine, " ")[2]
	state := make(map[int]bool)

	for i := 0; i < len(stateStr); i++ {
		state[i] = (string(stateStr[i]) == "#")
	}

	scanner.Scan()

	rules := make(map[string]bool)
	
	for scanner.Scan() {
		ruleSplit := strings.Split(scanner.Text(), " => ")
		rule := ruleSplit[0]

		if ruleSplit[1] == "#" {
			rules[rule] = true
		} else {
			rules[rule] = false
		}
	}

	for i := 0; i < 20; i++ {
		state = nextState(state, rules)
	}

	fmt.Println("Plant pot sum:", computePlantPotSum(state))
}
