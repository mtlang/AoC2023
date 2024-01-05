package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const DIGITS string = "0123456789"

func isSymbol(i byte) bool {
	return i < 46 || i == 47 || i > 57
}

func partOne() int {
	answer := 0
	var input []string

	// Load the input into an array of strings
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			input = append(input, line)
		}
	}

	// Parse out part numbers
	for lIndex, line := range input {
		startIndex := strings.IndexAny(line, DIGITS)
		for startIndex > -1 {

			endIndex := startIndex + 1
			for endIndex < len(line) && strings.Contains(DIGITS, string(line[endIndex])) {
				endIndex++
			}

			// Search for a special character around the number
			foundSymbol := false
			// Previous line
			if lIndex-1 >= 0 {
				// Diagonal up-left
				if startIndex >= 1 {
					foundSymbol = isSymbol(input[lIndex-1][startIndex-1])
				}

				// Directly above
				for i := startIndex; i < endIndex && !foundSymbol; i++ {
					foundSymbol = isSymbol(input[lIndex-1][i])
				}

				// Diagonal up-right
				if !foundSymbol && endIndex < len(line) {
					foundSymbol = isSymbol(input[lIndex-1][endIndex])
				}
			}

			// Left
			if !foundSymbol && startIndex > 0 {
				foundSymbol = isSymbol(line[startIndex-1])
			}

			// Right
			if !foundSymbol && endIndex < len(line) {
				foundSymbol = isSymbol(line[endIndex])
			}

			// Next line
			if !foundSymbol && lIndex+1 < len(input) {
				// Diagonal down-left
				if startIndex >= 1 {
					foundSymbol = isSymbol(input[lIndex+1][startIndex-1])
				}

				// Directly below
				for i := startIndex; i < endIndex && !foundSymbol; i++ {
					foundSymbol = isSymbol(input[lIndex+1][i])
				}

				// Diagonal down-right
				if !foundSymbol && endIndex < len(line) {
					foundSymbol = isSymbol(input[lIndex+1][endIndex])
				}
			}

			if foundSymbol {
				n, err := strconv.Atoi(line[startIndex:endIndex])
				if err != nil {
					panic(err)
				}
				answer += n
			}

			// Prepare for the next number in line (or move on to the next line)
			startIndex = strings.IndexAny(line[endIndex:], DIGITS)
			if startIndex > -1 {
				startIndex += endIndex
			}
		}
	}

	return answer
}

func partTwo() int {
	answer := 0
	var input []string

	// Load the input into an array of strings
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			input = append(input, line)
		}
	}

	// Parse out gear ratios
	for lIndex, line := range input {
		gearIndex := strings.Index(line, "*")
		for gearIndex > -1 {
			// Find numbers around the gear
			var foundNumbers []int
			var tempNumber string
			var finalNumber int

			// Check previous line
			if lIndex-1 >= 0 {
				// If there's a number directly above, it's the only number
				if strings.Contains(DIGITS, string(input[lIndex-1][gearIndex])) {
					tempNumber = string(input[lIndex-1][gearIndex])
					// Look back
					for i := gearIndex - 1; i >= 0; i-- {
						if strings.Contains(DIGITS, string(input[lIndex-1][i])) {
							tempNumber = string(input[lIndex-1][i]) + tempNumber
						} else {
							break
						}
					}
					// Look forward
					for i := gearIndex + 1; i < len(input[lIndex-1]); i++ {
						if strings.Contains(DIGITS, string(input[lIndex-1][i])) {
							tempNumber += string(input[lIndex-1][i])
						} else {
							break
						}
					}

					// Add found number
					finalNumber, err = strconv.Atoi(tempNumber)
					if err != nil {
						panic(err)
					}
					foundNumbers = append(foundNumbers, finalNumber)
				} else {
					// If not directly above, check upwards diagonals
					// Up-left
					if gearIndex > 0 && strings.Contains(DIGITS, string(input[lIndex-1][gearIndex-1])) {
						tempNumber = ""
						// Look back
						for i := gearIndex - 1; i >= 0; i-- {
							if strings.Contains(DIGITS, string(input[lIndex-1][i])) {
								tempNumber = string(input[lIndex-1][i]) + tempNumber
							} else {
								break
							}
						}
						// Add found number
						finalNumber, err = strconv.Atoi(tempNumber)
						if err != nil {
							panic(err)
						}
						foundNumbers = append(foundNumbers, finalNumber)
					}

					// Up-right
					if gearIndex < len(input[lIndex-1])-1 && strings.Contains(DIGITS, string(input[lIndex-1][gearIndex+1])) {
						tempNumber = ""
						// Look forward
						for i := gearIndex + 1; i < len(input[lIndex-1]); i++ {
							if strings.Contains(DIGITS, string(input[lIndex-1][i])) {
								tempNumber += string(input[lIndex-1][i])
							} else {
								break
							}
						}

						// Add found number
						finalNumber, err = strconv.Atoi(tempNumber)
						if err != nil {
							panic(err)
						}
						foundNumbers = append(foundNumbers, finalNumber)
					}
				}
			}

			// Check behind gear
			if gearIndex > 0 && strings.Contains(DIGITS, string(line[gearIndex-1])) {
				tempNumber = ""
				// Look back
				for i := gearIndex - 1; i >= 0; i-- {
					if strings.Contains(DIGITS, string(line[i])) {
						tempNumber = string(line[i]) + tempNumber
					} else {
						break
					}
				}
				// Add found number
				finalNumber, err = strconv.Atoi(tempNumber)
				if err != nil {
					panic(err)
				}
				foundNumbers = append(foundNumbers, finalNumber)
			}

			// From here on out, we can stop looking if we've already found 3 numbers
			// Check ahead of gear
			if len(foundNumbers) < 3 && gearIndex < len(line)-1 && strings.Contains(DIGITS, string(line[gearIndex+1])) {
				tempNumber = ""
				// Look forward
				for i := gearIndex + 1; i < len(line); i++ {
					if strings.Contains(DIGITS, string(line[i])) {
						tempNumber += string(line[i])
					} else {
						break
					}
				}

				// Add found number
				finalNumber, err = strconv.Atoi(tempNumber)
				if err != nil {
					panic(err)
				}
				foundNumbers = append(foundNumbers, finalNumber)
			}

			// Check next line
			if len(foundNumbers) < 3 && lIndex+1 < len(input) {
				// If there's a number directly below, it's the only number
				if strings.Contains(DIGITS, string(input[lIndex+1][gearIndex])) {
					tempNumber = string(input[lIndex+1][gearIndex])
					// Look back
					for i := gearIndex - 1; i >= 0; i-- {
						if strings.Contains(DIGITS, string(input[lIndex+1][i])) {
							tempNumber = string(input[lIndex+1][i]) + tempNumber
						} else {
							break
						}
					}
					// Look forward
					for i := gearIndex + 1; i < len(input[lIndex+1]); i++ {
						if strings.Contains(DIGITS, string(input[lIndex+1][i])) {
							tempNumber += string(input[lIndex+1][i])
						} else {
							break
						}
					}

					// Add found number
					finalNumber, err = strconv.Atoi(tempNumber)
					if err != nil {
						panic(err)
					}
					foundNumbers = append(foundNumbers, finalNumber)
				} else {
					// If not directly above, check upwards diagonals
					// Down-left
					if gearIndex > 0 && strings.Contains(DIGITS, string(input[lIndex+1][gearIndex-1])) {
						tempNumber = ""
						// Look back
						for i := gearIndex - 1; i >= 0; i-- {
							if strings.Contains(DIGITS, string(input[lIndex+1][i])) {
								tempNumber = string(input[lIndex+1][i]) + tempNumber
							} else {
								break
							}
						}
						// Add found number
						finalNumber, err = strconv.Atoi(tempNumber)
						if err != nil {
							panic(err)
						}
						foundNumbers = append(foundNumbers, finalNumber)
					}

					// Down-right
					if len(foundNumbers) < 3 && gearIndex < len(input[lIndex+1])-1 && strings.Contains(DIGITS, string(input[lIndex+1][gearIndex+1])) {
						tempNumber = ""
						// Look forward
						for i := gearIndex + 1; i < len(input[lIndex+1]); i++ {
							if strings.Contains(DIGITS, string(input[lIndex+1][i])) {
								tempNumber += string(input[lIndex+1][i])
							} else {
								break
							}
						}

						// Add found number
						finalNumber, err = strconv.Atoi(tempNumber)
						if err != nil {
							panic(err)
						}
						foundNumbers = append(foundNumbers, finalNumber)
					}
				}
			}

			// If we've got exactly two numbers, add the gear ratio
			if len(foundNumbers) == 2 {
				answer += foundNumbers[0] * foundNumbers[1]
			}

			// Find the next gear in line (or move on to the next line)
			if gearIndex >= len(line)-1 {
				break
			}
			oldIndex := gearIndex
			gearIndex = strings.Index(line[gearIndex+1:], "*")
			if gearIndex > -1 {
				gearIndex += oldIndex + 1
			}
		}
	}
	return answer
}

func main() {
	fmt.Printf("%d\n", partTwo())
}
