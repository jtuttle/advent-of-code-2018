package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coord struct {
	x int
	y int
}

func getHundredsDigit(num int) int {
	if num < 100 {
		return 0
	}
	
	numStr := strconv.Itoa(num)
	hundredsDigitStr := string(numStr[len(numStr) - 3])
	hundredsDigit, _ := strconv.Atoi(hundredsDigitStr)

	return hundredsDigit
}

func computePower(coord Coord, serialNum int) int {
	rackId := coord.x + 10
	power := rackId * coord.y
	power += serialNum
	power *= rackId
	power = getHundredsDigit(power)
	power -= 5
	
	return power
}

func createGrid(serialNum int) map[Coord]int {
	grid := make(map[Coord]int)

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			coord := Coord{ x, y }
			grid[coord] = computePower(coord, serialNum)
		}
	}

	return grid
}

func computeSquarePower(topLeftCoord Coord, grid map[Coord]int) int {
	squarePower := 0
	
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			coord := Coord{ topLeftCoord.x + x, topLeftCoord.y + y }
			squarePower += grid[coord]
		}
	}

	return squarePower
}

func main() {
	file, _ := os.Open("./day-11-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	serialNum, _ := strconv.Atoi(scanner.Text())
	
	grid := createGrid(serialNum)
	squarePowers := make(map[Coord]int)

	for y := 1; y < 298; y++ {
		for x := 1; x < 298; x++ {
			coord := Coord{ x, y }
			squarePowers[coord] = computeSquarePower(coord, grid)
		}
	}

	largestCoord := Coord{ 0, 0 }
	
	for k, v := range squarePowers {
		if v > squarePowers[largestCoord] {
			largestCoord = k
		}
	}
	
	fmt.Println("Coord with largest square power:", largestCoord)
}
