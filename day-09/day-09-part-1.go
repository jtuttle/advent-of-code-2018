package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Marble struct {
	value int
	prev *Marble
	next *Marble
}

func createMarble(value int, prev *Marble, next *Marble) *Marble {
	newMarble := &Marble {
		value: value,
		prev: prev,
		next: next,
	}

	return newMarble
}

func printMarbles(marble *Marble) {
	values := []int {
		marble.value,
	}
	
	currentMarble := marble.next

	for marble != currentMarble {
		values = append(values, currentMarble.value)
		currentMarble = currentMarble.next
	}

	fmt.Println(values)
}

func findHighestScore(scores map[int]int) int {
	highest := 0

	for _, v := range scores {
		if v > highest {
			highest = v
		}
	}
	
	return highest
}

func main() {
	file, _ := os.Open("./day-09-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	input := scanner.Text()

	playersRegex := regexp.MustCompile(`(\d*)(?:\splayers)`)
	playersStr := playersRegex.FindStringSubmatch(input)[1]
	players, _ := strconv.Atoi(playersStr)

	pointsRegex := regexp.MustCompile(`(\d*)(?:\spoints)`)
	pointsStr := pointsRegex.FindStringSubmatch(input)[1]
	points, _ := strconv.Atoi(pointsStr)

	scores := make(map[int]int)

	currentMarble := createMarble(0, nil, nil)
	currentMarble.prev = currentMarble
	currentMarble.next = currentMarble

	for i := 1; i <= points; i++ {
		player := ((i - 1) % players) + 1
		
		if i % 23 == 0 {
			remove := currentMarble

			for i := 0; i < 7; i++ {
				remove = remove.prev
			}

			scores[player] += (i + remove.value)

			oldNext := remove.next
			remove.prev.next = remove.next
			oldNext.prev = remove.prev

			currentMarble = oldNext
		} else {
			newMarble := createMarble(
				i,
				currentMarble.next,
				currentMarble.next.next,
			)

			newMarble.prev.next = newMarble
			newMarble.next.prev = newMarble

			currentMarble = newMarble
		}
	}

	fmt.Println("Highest score:", findHighestScore(scores))
}
