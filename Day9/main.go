package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseInput(fileName string) [][]int {
	sequences := make([][]int, 0)

	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { 
		line := scanner.Text()
		if len(line) > 0 {
			splitLine := strings.Split(line, " ")
			numbers := make([]int, len(splitLine))
			for i, s := range splitLine {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				numbers[i] = n
			}
			sequences = append(sequences, numbers)
		}
	}
	return sequences
}

func isAllZeroes(s []int) bool {
	for _, n := range s {
		if n != 0 {
			return false
		}
	}
	return true
}

func solve(fileName string, partTwo bool) int {
	answer := 0

	// Read all the numbers in the file
	input := parseInput(fileName)

	// Find the predicted next number for each sequence
	for _, sequence := range input {
		cascade := make([][]int, len(sequence))

		// Find differences between sequence items until the differences are all zero
		cascade[0] = sequence
		var i int
		for i = 0; !isAllZeroes(cascade[i]); i++ {
			newSequence := make([]int, len(cascade[i])-1)
			for j, n := range cascade[i] {
				if j == len(cascade[i])-1 {break}
				newSequence[j] = cascade[i][j+1] - n
			}
			cascade[i+1] = newSequence
		}
		cascade = cascade[:i+1]
		
		if !partTwo {
			// Once the cascade is complete, walk backwards through it until we find our predicted value
			cascade[len(cascade)-1] = append(cascade[len(cascade)-1], 0)
			for i := len(cascade)-1; i > 0; i-- {
				nextVal := cascade[i-1][len(cascade[i-1])-1] + cascade[i][len(cascade[i])-1]
				cascade[i-1] = append(cascade[i-1], nextVal)
			}
			// fmt.Println(cascade[0][len(cascade[0])-1])

			// Add the predicted value to the final answer
			answer += cascade[0][len(cascade[0])-1]
		} else {
			// Once the cascade is complete, walk backwards through it until we find our predicted value
			cascade[len(cascade)-1] = append([]int{0}, cascade[len(cascade)-1]...)
			for i := len(cascade)-1; i > 0; i-- {
				prevVal := cascade[i-1][0] - cascade[i][0]
				cascade[i-1] = append([]int{prevVal}, cascade[i-1]...)
			}
			// fmt.Println(cascade[0][0])

			// Add the predicted value to the final answer
			answer += cascade[0][0]
		}
	}

	return answer
}

func main() {
	fmt.Println(solve("input.txt", false)) 

	fmt.Println(solve("input.txt", true))
}
