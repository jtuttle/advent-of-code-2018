package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	frequency := 0
	
	file, _ := os.Open("./day-01-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		delta, _ := strconv.Atoi(scanner.Text())
		frequency += delta
	}

	fmt.Println("Final frequency:", frequency)
}
