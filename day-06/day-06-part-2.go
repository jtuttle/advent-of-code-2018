package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Rect struct {
	topLeft, bottomRight Coord
}

type CoordDistance struct {
	coord Coord
	distance int
}

func parseCoord(line string) Coord {
	lineSplit := strings.Split(line, ", ")

	x, _ := strconv.Atoi(lineSplit[0])
	y, _ := strconv.Atoi(lineSplit[1])

	return Coord{ x, y }
}

func findBoundingRect(coords []Coord) Rect {
	var xMin, xMax, yMin, yMax *int

	for _, coord := range coords {
		x := coord.x
		y := coord.y
		
		if xMin == nil || x < *xMin {
			xMin = &x
		}

		if yMin == nil || y < *yMin {
			yMin = &y
		}
		
		if xMax == nil || x > *xMax {
			xMax = &x
		}

		if yMax == nil || y > *yMax {
			yMax = &y
		}
	}

	return Rect{
		topLeft: Coord{ *xMin, *yMin },
		bottomRight: Coord{ *xMax, *yMax },
	}
}

func manhattanDistance(coord1 Coord, coord2 Coord) int {
	dy := math.Abs(float64(coord2.y - coord1.y))
	dx := math.Abs(float64(coord2.x - coord1.x))

	return int(dy + dx)
}

func computeDistanceSum(gridCoord Coord, coords []Coord) int {
	sum := 0

	for _, coord := range coords {
		sum += manhattanDistance(gridCoord, coord)
	}

	return sum
}

func computeDistanceSums(coords []Coord, boundingRect Rect) [][]int {
	topLeft := boundingRect.topLeft
	bottomRight := boundingRect.bottomRight

	dx := bottomRight.x - topLeft.x + 1
	dy := bottomRight.y - topLeft.y + 1
	
	closestCoords := make([][]int, dy)

	for y := 0; y < dy; y++ {
		closestCoords[y] = make([]int, dx)
		
		for x := 0; x < dx; x++ {		
			gridCoord := Coord{ topLeft.x + x, topLeft.y + y }
			closestCoords[y][x] = computeDistanceSum(gridCoord, coords)
		}
	}

	return closestCoords
}

func computeSafestRegionSize(distanceSums [][]int, threshold int) int {
	regionSize := 0
	
	for _, col := range distanceSums {
		for _, sum := range col {
			if sum < threshold {
				regionSize++
			}
		}
	}

	return regionSize
}

func printSafestRegion(grid [][]int, threshold int) {
	for _, row := range grid {
		for _, sum := range row {
			if sum < threshold {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("./day-06-input.txt")
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	coords := make([]Coord, 0)
	
	for scanner.Scan() {
		coords = append(coords, parseCoord(scanner.Text()))
	}

	boundingRect := findBoundingRect(coords)
	distanceSums := computeDistanceSums(coords, boundingRect)

	//printSafestRegion(distanceSums, 10000)

	regionSize := computeSafestRegionSize(distanceSums, 10000)

	fmt.Println("Region size:", regionSize)
}
