package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func partOne() string {
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, "Day1", "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		var r1 string = ""
		var r2 string = ""
		for _, char := range scanner.Text() {
			if _, err := strconv.Atoi(string(char)); err == nil {
				if r1 == "" {
					r1 = string(char)
				}
				r2 = string(char)
			}
		}
		val, err := strconv.Atoi(r1 + r2)
		if err != nil {
			panic(err)
		}
		total += val
	}
	return strconv.Itoa(total)
}

func partTwo() string {
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, "Day1", "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Find the first and last digit in the string
		r1, r2 := "", ""
		r1Index, r2Index := -1, -1
		for i, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				if r1 == "" {
					r1 = string(char)
					r1Index = i
				}
				r2 = string(char)
				r2Index = i
			}
		}

		// Find the first and last number (spelled out) in the string
		numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		firstIndex := 999
		lastIndex := -1
		n1, n2 := -1, -1
		for i, n := range numbers {
			i1 := strings.Index(line, n)
			i2 := strings.LastIndex(line, n)

			if i1 != -1 && i1 < firstIndex {
				firstIndex = i1
				n1 = i
			}
			if i2 > lastIndex {
				lastIndex = i2
				n2 = i
			}
		}

		// Determine final result
		d1, d2 := "", ""
		if r1Index < firstIndex {
			d1 = r1
		} else {
			d1 = strconv.Itoa(n1 + 1)
		}

		if r2Index > lastIndex {
			d2 = r2
		} else {
			d2 = strconv.Itoa(n2 + 1)
		}

		val, err := strconv.Atoi(d1 + d2)
		if err != nil {
			panic(err)
		}
		total += val
	}
	return strconv.Itoa(total)
}

func main() {
	fmt.Println(partOne())

	fmt.Println(partTwo())
}
