package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"math"
)

func countWinners(winnerStrings []string, gameStrings []string) int {
	count := 0
	for _, n := range gameStrings {
		if n != "" && slices.Contains(winnerStrings, n) {
			count++
		}
	}
	return count
}

func partOne(fileName string) int {
	answer := 0

	// Load the input into an array of strings
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
			// Throw out text before the colon
			line = strings.Split(line, ":")[1]

			// Parse out numbers
			numberStrings := strings.Split(line, "|")
			winnerStrings := strings.Split(numberStrings[0], " ")
			gameStrings := strings.Split(numberStrings[1], " ")
			
			score := countWinners(winnerStrings, gameStrings)

			if score > 0 {
				score = int(math.Pow(2, float64(score-1)))
			}

			answer += score
		}
	}

	return answer
}

func partTwo(fileName string) int {
	answer := 0

	// Load the input into an array of strings
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Figure out how many lines we have so we can initialize a slice
	lineCount := 0
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			lineCount++
		}
	}
	
	// Create the multipliers slice, starting each multiplier at 1
	multipliers := make([]int, lineCount)
	for i := range multipliers {
		multipliers[i] = 1
	}

	file.Seek(0,0)
	scanner = bufio.NewScanner(file)
	loopIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			// Throw out text before the colon
			line = strings.Split(line, ":")[1]

			// Parse out numbers
			numberStrings := strings.Split(line, "|")
			winnerStrings := strings.Split(numberStrings[0], " ")
			gameStrings := strings.Split(numberStrings[1], " ")
			
			score := countWinners(winnerStrings, gameStrings)
			for i := 1; loopIndex + i < lineCount && i <= score; i++ {
				multipliers[loopIndex+i] += multipliers[loopIndex]
			}
			loopIndex++
		}
	}

	// The answer is the sum of all multipliers
	for _, m := range multipliers {
		answer += m
	}

	return answer
}

func main() {
	fmt.Printf("%d\n", partTwo("input.txt"))
}
