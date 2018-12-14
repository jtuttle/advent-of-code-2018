package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x int
	y int
}

func (coord *Coord) Add(other Coord) Coord {
	return Coord{ coord.x + other.x, coord.y + other.y }
}

func (coord *Coord) String() string {
	return fmt.Sprintf("%d,%d", coord.x, coord.y)
}

type Cart struct {
	position Coord
	direction rune
	turnCount int
}

func createCart(coord Coord, direction rune) *Cart {
	return &Cart{
		position: coord,
		direction: direction,
		turnCount: 0,
	}
}

func isCart(char rune) bool {
	return char == '^' || char == 'v' || char == '<' || char == '>'
}

func isTurnTrack(char rune) bool {
	return char == '/' || char == '\\' || char == '+'
}

func trackForCart(char rune) rune {
	if char == '^' || char == 'v' {
		return '|'
	} else if char == '<' || char == '>' {
		return '-'
	}

	panic(fmt.Sprintf("invalid cart char:", char))
}

func findCrash(carts []*Cart) *Coord {
	positions := make(map[Coord]bool)

	for _, cart := range carts {
		if positions[cart.position] {
			return &cart.position
		}
		positions[cart.position] = true
	}
	
	return nil
}

func nextCoordForCart(cart *Cart) Coord {
	position := cart.position
	direction := cart.direction

	if direction == '^' {
		return position.Add(Coord{ 0, -1 })
	} else if direction == 'v' {
		return position.Add(Coord{ 0, 1 })
	} else if direction == '<' {
		return position.Add(Coord{ -1, 0 })
	} else if direction == '>' {
		return position.Add(Coord{ 1, 0 })
	}

	panic(fmt.Sprintf("invalid cart direction:", direction))
}

func moveCart(tracks map[Coord]rune, cart *Cart) {
	nextCoord := nextCoordForCart(cart)
	cart.position = nextCoord

	track := tracks[nextCoord]

	if isTurnTrack(track) {
		turnCart(track, cart)
	}
}

func turnCart(track rune, cart *Cart) {
	direction := cart.direction
	
	if track == '/' {
		if direction == '^' {
			cart.direction = '>'
		} else if direction == 'v' {
			cart.direction = '<'
		} else if direction == '<' {
			cart.direction = 'v'
		} else if direction == '>' {
			cart.direction = '^'
		}
	} else if track == '\\' {
		if direction == '^' {
			cart.direction = '<'
		} else if direction == 'v' {
			cart.direction = '>'
		} else if direction == '<' {
			cart.direction = '^'
		} else if direction == '>' {
			cart.direction = 'v'
		}
	} else if track == '+' {
		turnCountMod := cart.turnCount % 3

		if turnCountMod == 0 {
			if direction == '^' {
				cart.direction = '<'
			} else if direction == 'v' {
				cart.direction = '>'
			} else if direction == '<' {
				cart.direction = 'v'
			} else if direction == '>' {
				cart.direction = '^'
			}
		} else if turnCountMod == 2 {
			if direction == '^' {
				cart.direction = '>'
			} else if direction == 'v' {
				cart.direction = '<'
			} else if direction == '<' {
				cart.direction = '^'
			} else if direction == '>' {
				cart.direction = 'v'
			}
		}

		cart.turnCount++
	}
}

func printState(tracks map[Coord]rune, carts []*Cart) {
	var xMax, yMax int

	for k, _ := range tracks {
		if k.x > xMax {
			xMax = k.x
		}

		if k.y > yMax {
			yMax = k.y
		}
	}

	cartMap := make(map[Coord]string)

	for _, cart := range carts {
		if cartMap[cart.position] != "" {
			cartMap[cart.position] = "X"
		} else {
			cartMap[cart.position] = string(cart.direction)
		}
	}
	
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			coord := Coord{ x, y }

			if cartMap[coord] != "" {
				fmt.Print(cartMap[coord])
			} else {
				if tracks[coord] == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print(string(tracks[coord]))
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("./day-13-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0

	tracks := make(map[Coord]rune)
	carts := make([]*Cart, 0)

	for scanner.Scan() {
		chars := []rune(scanner.Text())

		for x, char := range chars {
			if char == ' ' {
				continue
			}
			
			coord := Coord{ x, y }

			if isCart(char) {
				carts = append(carts, createCart(coord, char))
				tracks[coord] = trackForCart(char)
			} else {
				tracks[coord] = char
			}
		}

		y++
	}

	var crash *Coord

	for crash == nil {
		for _, cart := range carts {
			moveCart(tracks, cart)
			crash = findCrash(carts)

			if crash != nil {
				break
			}
		}
	}
	
	//printState(tracks, carts)

	fmt.Println("Crash occurred at:", crash)
}
