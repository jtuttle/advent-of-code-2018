// Note: I'm treating 0, 0 (an empty Coord struct) as the indicator that a grid
// coord is equidistant from two different input coordinates, meaning this won't
// work when 0, 0 is in the input.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func findClosestCoord(gridCoord Coord, coords []Coord) Coord {
	distances := make([]CoordDistance, len(coords))
	
	for i, coord := range coords {
		distance := manhattanDistance(coord, gridCoord)
		distances[i] = CoordDistance { coord, distance }
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	if distances[0].distance == distances[1].distance {
		return Coord{}
	} else {
		return distances[0].coord
	}
}

func computeClosestCoords(coords []Coord, boundingRect Rect) [][]Coord {
	topLeft := boundingRect.topLeft
	bottomRight := boundingRect.bottomRight

	dx := bottomRight.x - topLeft.x + 1
	dy := bottomRight.y - topLeft.y + 1
	
	closestCoords := make([][]Coord, dy)

	for y := 0; y < dy; y++ {
		closestCoords[y] = make([]Coord, dx)
		
		for x := 0; x < dx; x++ {		
			gridCoord := Coord{ topLeft.x + x, topLeft.y + y }
			closestCoords[y][x] = findClosestCoord(gridCoord, coords)
		}
	}

	return closestCoords
}

func computeAreas(coords []Coord, boundingRect Rect) map[Coord]int {
	areas := make(map[Coord]int)
	
	closestCoords := computeClosestCoords(coords, boundingRect)

	//printClosestCoordGrid(closestCoords)

	for _, col := range closestCoords {
		for _, coord := range col {
			if coord != (Coord{}) {
				areas[coord] += 1
			}
		}
	}

	return areas
}

func findFiniteAreas(areaMap1, areaMap2 map[Coord]int) map[Coord]int {
	finiteAreas := make(map[Coord]int)
	
	for k, v := range areaMap1 {
		if areaMap2[k] == v {
			finiteAreas[k] = v
		}
	}

	return finiteAreas
}

func printClosestCoordGrid(closestCoords [][]Coord) {
	for _, row := range closestCoords {
		fmt.Println(row)
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
	areas := computeAreas(coords, boundingRect)
	
	largerRect := Rect{
		topLeft: Coord{ boundingRect.topLeft.x - 5, boundingRect.topLeft.y - 5 },
		bottomRight: Coord { boundingRect.bottomRight.x + 5, boundingRect.bottomRight.y + 5 },
	}
	largerAreas := computeAreas(coords, largerRect)
	
	finiteAreas := findFiniteAreas(areas, largerAreas)

	var largestArea int

	for _, v := range finiteAreas {
		if v > largestArea {
			largestArea = v
		}
	}

	fmt.Println("Largest area:", largestArea)
}
