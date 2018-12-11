package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Light struct {
	position Vector2
	velocity Vector2
}

type Vector2 struct {
	x int
	y int
}

type Rect struct {
	topLeft Vector2
	bottomRight Vector2
}

func createVector2(coordStr string) Vector2 {
	coordsSplit := strings.Split(coordStr, ",")

	x, _ := strconv.Atoi(strings.TrimSpace(coordsSplit[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(coordsSplit[1]))
	
	return Vector2{ x, y }
}

func parseLight(line string) Light {
	positionRegex := regexp.MustCompile(`(?:position=<)([\s\-\d,]*)(>)`)
	positionStr := positionRegex.FindStringSubmatch(line)[1]
	
	velocityRegex := regexp.MustCompile(`(?:velocity=<)([\s\-\d,]*)(>)`)
	velocityStr := velocityRegex.FindStringSubmatch(line)[1]

	return Light{
		position: createVector2(positionStr),
		velocity: createVector2(velocityStr),
	}
}

func findBounds(lights []Light) Rect {
	var xMin, xMax, yMin, yMax *int
	
	for _, light := range lights {
		x := light.position.x
		y := light.position.y
		
		if xMin == nil || x < *xMin {
			xMin = &x
		}

		if xMax == nil || x > *xMax {
			xMax = &x
		}

		if yMin == nil || y < *yMin {
			yMin = &y
		}

		if yMax == nil || y > *yMax {
			yMax = &y
		}
	}

	return Rect{
		topLeft: Vector2{ *xMin, *yMin },
		bottomRight: Vector2{ *xMax, *yMax },
	}
}

func advanceLights(lights []Light) []Light {
	advanced := make([]Light, len(lights))
	
	for i, light := range lights {
		newPos := Vector2{
			x: light.position.x + light.velocity.x,
			y: light.position.y + light.velocity.y,
		}

		light.position = newPos

		advanced[i] = light
	}

	return advanced
}

func printLights(lights []Light) {
	bounds := findBounds(lights)
	
	topLeft := bounds.topLeft
	bottomRight := bounds.bottomRight

	lightMap := make(map[Vector2]bool)

	for _, light := range lights {
		lightMap[light.position] = true
	}
	
	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			coord := Vector2{ x, y }

			if lightMap[coord] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("./day-10-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lights := make([]Light, 0)
	
	for scanner.Scan() {
		line := scanner.Text()
		light := parseLight(line)
		lights = append(lights, light)
	}

	for true {
		bounds := findBounds(lights)
		boundsWidth := math.Abs(float64(bounds.bottomRight.x - bounds.topLeft.x))
		
		nextLights := advanceLights(lights)
		
		nextBounds := findBounds(nextLights)
		nextBoundsWidth := math.Abs(float64(nextBounds.bottomRight.x - nextBounds.topLeft.x))

		if nextBoundsWidth > boundsWidth {
			break
		}

		lights = nextLights
		bounds = nextBounds
	}
        
	printLights(lights)
}
