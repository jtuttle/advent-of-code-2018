package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func computeValue(licenseNums []int, ptr *int) int {
	childCount := licenseNums[*ptr]
	(*ptr)++
	
	metaCount := licenseNums[*ptr]
	(*ptr)++

	childValues := make(map[int]int)
	
	for i := 0; i < childCount; i++ {
		childValues[i+1] = computeValue(licenseNums, ptr)
	}

	value := 0

	for i := 0; i < metaCount; i++ {
		metadata := licenseNums[*ptr]
		(*ptr)++

		if childCount == 0 {
			value += metadata
		} else {
			if val, ok := childValues[metadata]; ok {
				value += val
			}
		}
	}

	return value
}

func main() {
	file, _ := os.Open("./day-08-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	license := scanner.Text()
	licenseArray := strings.Split(license, " ")

	licenseNums := make([]int, len(licenseArray))

	for i, numStr := range licenseArray {
		numInt, _ := strconv.Atoi(numStr)
		licenseNums[i] = numInt
	}

	idx := 0
	ptr := &idx
	metadataSum := computeValue(licenseNums, ptr)

	fmt.Println("Sum of metadata:", metadataSum)
}
