package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumMetadata(licenseNums []int, ptr *int) int {
	childCount := licenseNums[*ptr]
	(*ptr)++
	
	metaCount := licenseNums[*ptr]
	(*ptr)++

	metadataSum := 0

	for i := 0; i < childCount; i++ {
		metadataSum += sumMetadata(licenseNums, ptr)
	}

	for i := 0; i < metaCount; i++ {
		metadataSum += licenseNums[*ptr]
		(*ptr)++
	}

	return metadataSum
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
	
	metadataSum := sumMetadata(licenseNums, ptr)

	fmt.Println("Sum of metadata:", metadataSum)
}
