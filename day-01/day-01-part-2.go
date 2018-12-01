package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}

	return lines, scanner.Err()
}

func main() {
	lines, _ := readLines("./day-01-input.txt")

	frequency := 0
	nums := make(map[int]bool)
	i := 0
	var firstRepeat *int
	
	for firstRepeat == nil {
		frequency += lines[i]

		if nums[frequency] {
			firstRepeat = &frequency
		}

		nums[frequency] = true
		
		i = (i + 1) % len(lines)
	}

	fmt.Println("First repeated frequency:", *firstRepeat)
}
