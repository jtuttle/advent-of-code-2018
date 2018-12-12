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

type SquarePower struct {
	coord Coord
	squareSize int
	value int
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

func computeEdgePower(topLeftCoord Coord, grid map[Coord]int, squareSize int) int {
	edgePower := 0
	
	for y := 0; y < squareSize; y++ {
		coord := Coord{ topLeftCoord.x + (squareSize - 1), topLeftCoord.y + y }
		edgePower += grid[coord]
	}

	for x := 0; x < squareSize - 1; x++ {
		coord := Coord{ topLeftCoord.x + x, topLeftCoord.y + (squareSize - 1) }
		edgePower += grid[coord]
	}
	
	return edgePower
}

func computeSquarePowers(grid map[Coord]int, squareSize int, squarePowers map[Coord]int) {
	lastIndex := 300 - squareSize - 1
	
	for y := 1; y < lastIndex; y++ {
		for x := 1; x < lastIndex; x++ {
			coord := Coord{ x, y }
			edgePower := computeEdgePower(coord, grid, squareSize)
			squarePowers[coord] += edgePower
		}
	}
}

func printGrid(grid map[Coord]int, topLeftCoord Coord, squareSize int) {
	for y := topLeftCoord.y; y < topLeftCoord.y + squareSize; y++ {
		for x := topLeftCoord.x; x < topLeftCoord.x + squareSize; x++ {
			coord := Coord { x, y }

			if grid[coord] >= 0 && grid[coord] < 10 {
				fmt.Print(" ")
			}
			
			fmt.Print(grid[coord], " ")
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("./day-11-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	serialNum, _ := strconv.Atoi(scanner.Text())
	
	grid := createGrid(serialNum)
	
	var squarePowers = make(map[Coord]int)
	var largestSquarePower SquarePower
	
	for i := 1; i <= 300; i++ {
		computeSquarePowers(grid, i, squarePowers)

		for k, v := range squarePowers {
			if largestSquarePower == (SquarePower{}) || v > largestSquarePower.value {
				largestSquarePower = SquarePower{
					coord: k,
					squareSize: i,
					value: v,
				}
			}
		}
	}
	
	fmt.Println("Coord with largest square power:", largestSquarePower)
}
