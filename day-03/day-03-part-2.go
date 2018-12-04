package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Rect struct {
	id int
	topLeft Coord
	width, height int
}

func findId(line string) int {
	regex := regexp.MustCompile(`(?:#)(\d*)`)
	stringId := regex.FindStringSubmatch(line)[1]
	id, _ := strconv.Atoi(stringId)
	return id
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
		id: findId(line),
		topLeft: coord,
		width: dimensions[0],
		height: dimensions[1],
	}
}

func hasNoOverlap(rect Rect, material map[Coord]int) bool {
	for x := rect.topLeft.x; x < rect.topLeft.x + rect.width; x++ {
		for y := rect.topLeft.y; y < rect.topLeft.y + rect.height; y++ {
			coord := Coord{x, y}
			if material[coord] > 1 {
				return false
			}
		}
	}

	return true
}

func main() {
	file, _ := os.Open("./day-03-input.txt")
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	material := make(map[Coord]int)

	rects := make([]Rect, 0)
	
	for scanner.Scan() {
		rects = append(rects, parseRect(scanner.Text()))
	}

	for _, rect := range rects {
		for x := rect.topLeft.x; x < rect.topLeft.x + rect.width; x++ {
			for y := rect.topLeft.y; y < rect.topLeft.y + rect.height; y++ {
				coord := Coord{x, y}
				material[coord] += 1
			}
		}
	}

	for _, rect := range rects {
		if hasNoOverlap(rect, material) {
			fmt.Println("No overlap rectangle ID:", rect.id)
		}
	}
}
