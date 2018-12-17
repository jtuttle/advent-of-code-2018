package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func scoresToString(scores []int) string {
	last := make([]string, len(scores))
	
	for i := 0; i < len(scores); i++ {
		last[i] = strconv.Itoa(scores[i])
	}
	
	return strings.Join(last, "")
}

func main() {
	file, _ := os.Open("./day-14-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	
	sequence := scanner.Text()

	scores := []int { 3, 7 }
	elfIndices := []int { 0, 1 }

	for len(scores) < 30000000 {
		sum := findNextSum(scores, elfIndices)
		scores = appendSum(scores, sum)
	}
		
	scoresStr := scoresToString(scores)

	regex := regexp.MustCompile(sequence)
	match := regex.FindStringIndex(scoresStr)

	fmt.Println("Recipes before sequence:", match[0])
}
