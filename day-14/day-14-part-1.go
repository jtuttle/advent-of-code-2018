package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findNextSum(scores []int, elfIndices []int) int {
	sum := 0
	
	for i, elfIndex := range elfIndices {
		numSteps := elfIndices[i] + (scores[elfIndex] + 1)
		elfIndices[i] = numSteps % len(scores)
		sum += scores[elfIndices[i]]
	}

	return sum
}

func appendSum(scores []int, sum int) []int {
	digits := strings.Split(strconv.Itoa(sum), "")

	for _, digit := range digits {
		digitInt, _ := strconv.Atoi(digit)
		scores = append(scores, digitInt)
	}
	
	return scores
}

func lastTenToString(scores []int, recipeCount int) string {
	lastTen := make([]string, 10)
	
	for i := recipeCount; i < recipeCount + 10; i++ {
		score := strconv.Itoa(scores[i])
		lastTen = append(lastTen, score)
	}
	
	return strings.Join(lastTen, "")
}

func main() {
	file, _ := os.Open("./day-14-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	
	recipeCount, _ := strconv.Atoi(scanner.Text())

	scores := []int { 3, 7 }
	elfIndices := []int { 0, 1 }

	for len(scores) < (recipeCount + 10) {
		sum := findNextSum(scores, elfIndices)
		scores = appendSum(scores, sum)
	}

	fmt.Println("Last 10 recipes:", lastTenToString(scores, recipeCount))
}
