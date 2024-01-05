package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type numberRange struct {
	destStart int
	sourceStart int
	length int
}

type seedRange struct {
	start int
	length int
}

func parseInput1(fileName string) ([]int, [][]numberRange) {
	var seeds []int
	seedMaps := make([][]numberRange, 7)
	
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	// For the sake of simplicity, I'm assuming the maps are always in the given order
	scanner := bufio.NewScanner(file)

	// Start by parsing the seed numbers
	scanner.Scan()
	seedInput := scanner.Text()
	splitSeedInput := strings.Split(seedInput, " ")
	for i := 1; i < len(splitSeedInput); i++ {
		n, err := strconv.Atoi(splitSeedInput[i])
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, n)
	}

	// Then we can parse out the maps
	mapIndex := -1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			// A line containing the word "to" indicates the start of the next map
			if strings.Contains(line, "to") {
				mapIndex++
			} else {
				// If a line is not blank and does not indicate a new map, the line is a map
				mapData := strings.Split(line, " ")
				destStart, err0 := strconv.Atoi(mapData[0])
				sourceStart, err1 := strconv.Atoi(mapData[1])
				length, err2 := strconv.Atoi(mapData[2])
				if err0 != nil || err1 != nil || err2 != nil {
					panic(line)
				}

				nr := numberRange{
					destStart: destStart,
					sourceStart: sourceStart,
					length: length,
				}
				seedMaps[mapIndex] = append(seedMaps[mapIndex], nr)
			}
		}
	}

	return seeds, seedMaps
}

func parseInput2(fileName string) ([]seedRange, [][]numberRange) {
	var seeds []seedRange
	seedMaps := make([][]numberRange, 7)
	
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	// For the sake of simplicity, I'm assuming the maps are always in the given order
	scanner := bufio.NewScanner(file)

	// Start by parsing the seed numbers
	scanner.Scan()
	seedInput := scanner.Text()
	splitSeedInput := strings.Split(seedInput, " ")
	for i := 1; i < len(splitSeedInput); i+=2 {
		start, err1 := strconv.Atoi(splitSeedInput[i])
		length, err2 := strconv.Atoi(splitSeedInput[i+1])
		if err1 != nil || err2 != nil {
			panic(fmt.Sprintf("%s %s", splitSeedInput[i], splitSeedInput[i+1]))
		}
		n := seedRange{
			start: start,
			length: length,
		}
		seeds = append(seeds, n)
	}

	// Then we can parse out the maps
	mapIndex := -1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			// A line containing the word "to" indicates the start of the next map
			if strings.Contains(line, "to") {
				mapIndex++
			} else {
				// If a line is not blank and does not indicate a new map, the line is a map
				mapData := strings.Split(line, " ")
				destStart, err0 := strconv.Atoi(mapData[0])
				sourceStart, err1 := strconv.Atoi(mapData[1])
				length, err2 := strconv.Atoi(mapData[2])
				if err0 != nil || err1 != nil || err2 != nil {
					panic(line)
				}

				nr := numberRange{
					destStart: destStart,
					sourceStart: sourceStart,
					length: length,
				}
				seedMaps[mapIndex] = append(seedMaps[mapIndex], nr)
			}
		}
	}

	return seeds, seedMaps
}

func findMinimumLocation(seeds []int, seedMaps [][]numberRange) int {
	answer := -1
	// Process each seed to determine the lowest location number
	var targetNumber int
	for _, sourceNumber := range seeds {
		for _, conversionRanges := range seedMaps {
			// If the sourceNumber isn't in a range, the conversion is 1:1
			targetNumber = sourceNumber

			// If the sourceNumber is in a range, use that range to find targetNumber
			for _, conversionRange := range conversionRanges {
				if sourceNumber >= conversionRange.sourceStart && sourceNumber < (conversionRange.sourceStart + conversionRange.length) {
					difference := conversionRange.destStart - conversionRange.sourceStart
					targetNumber = sourceNumber + difference
				}
			}

			// Prepare for the next conversion1
			sourceNumber = targetNumber
		}

		// After all of the conversions, you're left with the location number
		if answer == -1 || targetNumber < answer {
			answer = targetNumber
		}
	}
	return answer
}

func partOne(fileName string) int {
	// Parse the input
	seeds, seedMaps := parseInput1(fileName)

	// Determine the minimum location number
	return findMinimumLocation(seeds, seedMaps)
}

func partTwoBrute(fileName string) int {
	// Parse the input
	seedRanges, seedMaps := parseInput2(fileName)

	answer := -1
	// Process each seed to determine the lowest location number
	var targetNumber int
	for _, seedRange := range seedRanges {
		for seedNumber := seedRange.start; seedNumber < (seedRange.start + seedRange.length); seedNumber++ {
			sourceNumber := seedNumber
			for _, conversionRanges := range seedMaps {
				// If the sourceNumber isn't in a range, the conversion is 1:1
				targetNumber = sourceNumber

				// If the sourceNumber is in a range, use that range to find targetNumber
				for _, conversionRange := range conversionRanges {
					if sourceNumber >= conversionRange.sourceStart && sourceNumber < (conversionRange.sourceStart + conversionRange.length) {
						difference := conversionRange.destStart - conversionRange.sourceStart
						targetNumber = sourceNumber + difference
					}
				}

				// Prepare for the next conversion
				sourceNumber = targetNumber
			}
			
			// After all of the conversions, you're left with the location number
			if answer == -1 || targetNumber < answer {
				answer = targetNumber
			}
		}
	}

	return answer
}

func partTwo(fileName string) int {
	// Parse the input
	seedRanges, seedMaps := parseInput2(fileName)

	// Test each possible location value until we find one in a seed range
	answer := 0
	var sourceNumber int
	for true {
		// Walk backwards through the maps to find the corresponding seed number
		targetNumber := answer
		for i := len(seedMaps)-1; i >= 0; i-- {
			// If the targetNumber isn't in a range, the conversion is 1:1
			sourceNumber = targetNumber

			// If the targetNumber is in a range, use that range to find sourceNumber
			for _, conversionRange := range seedMaps[i] {
				if targetNumber >= conversionRange.destStart && targetNumber < (conversionRange.destStart + conversionRange.length) {
					difference := conversionRange.sourceStart - conversionRange.destStart
					sourceNumber = targetNumber + difference
				}
			}

			// Prepare for the next conversion
			targetNumber = sourceNumber
		}

		// Check to see if the seed number is included in a seed range
		for _, seedRange := range seedRanges {
			if sourceNumber >= seedRange.start && sourceNumber < (seedRange.start + seedRange.length) {
				return answer
			}
		}
		answer++
	}

	return answer
}

func main() {
	fmt.Println(partOne("input.txt")) 

	fmt.Println(partTwo("input.txt"))
}
