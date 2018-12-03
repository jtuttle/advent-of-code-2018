package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

type Coord struct {
	x, y int
}

type Rect struct {
	topLeft Coord
	width, height int
}

func findNumPair(line string, delimiter string) []int {
	regex := regexp.MustCompile(`\d*` + delimiter + `\d*`)
	strings := strings.Split(regex.FindString(line), delimiter)

	one, _ := strconv.Atoi(strings[0])
	two, _ := strconv.Atoi(strings[1])
	
	return []int { one, two }
}

func parseRect(line string) Rect {
	coordSlice := findNumPair(line, ",")
	dimensions := findNumPair(line, "x")

	coord := Coord{coordSlice[0], coordSlice[1]}
	
	return Rect{
		topLeft: coord,
		width: dimensions[0],
		height: dimensions[1],
	}
}

func main() {
	file, _ := os.Open("./day-03-input.txt")
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	material := make(map[Coord]int)
	
	for scanner.Scan() {
		rect := parseRect(scanner.Text())

		for x := rect.topLeft.x; x < rect.topLeft.x + rect.width; x++ {
			for y := rect.topLeft.y; y < rect.topLeft.y + rect.height; y++ {
				coord := Coord{x, y}
				material[coord] += 1
			}
		}
	}

	overlapCount := 0

	for _, v := range material {
		if v > 1 {
			overlapCount++
		}
	}

	fmt.Println("Number of overlapping square inches:", overlapCount)
}
