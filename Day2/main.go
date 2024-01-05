package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	r = iota
	g = iota
	b = iota
)

func partOneMisinterpretation() string {
	answer := 0
	maxCubes := [3]int{12,13,14}

	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, "Day2", "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		totalCubes := [3]int{0,0,0}
		isValid := true
		// Determine sum of each cube
		s1 := strings.Split(line, ":")
		rounds := strings.Split(s1[1], ";")
		for _, round := range rounds {
			roundData := strings.Split(round, ",")
			for _, roundDatum := range roundData {
				roundArgs := strings.Split(strings.TrimSpace(roundDatum), " ")
				quantity, err := strconv.Atoi(roundArgs[0])
				if err != nil {
					panic(err)
				}
					
				color := roundArgs[1]
				if color == "red" {
					totalCubes[r] += quantity
				} else if color == "green" {
					totalCubes[g] += quantity
				} else if color == "blue" {
					totalCubes[b] += quantity
				} else {panic(color)}
			}
		}
		// If the game is valid, add index to total
		fmt.Printf("%d %d %d\n", totalCubes[0], totalCubes[1], totalCubes[2])
		for i, t := range totalCubes {
			if t > maxCubes[i] {
				isValid = false
			}
		}
		if isValid {
			s2 := strings.Split(strings.TrimSpace(s1[0]), " ")
			index, err := strconv.Atoi(s2[1])
			if err != nil {
				panic(err)
			}
			answer += index
		}
	}

	return strconv.Itoa(answer)
}

func partOne() string {
	answer := 0
	maxCubes := [3]int{12,13,14}

	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, "Day2", "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Determine sum of each cube
		isValid := true
		s1 := strings.Split(line, ":")
		rounds := strings.Split(s1[1], ";")
		for _, round := range rounds {
			roundData := strings.Split(round, ",")
			for _, roundDatum := range roundData {
				roundArgs := strings.Split(strings.TrimSpace(roundDatum), " ")
				quantity, err := strconv.Atoi(roundArgs[0])
				if err != nil {
					panic(err)
				}
					
				color := roundArgs[1]
				if color == "red" {
					if quantity > maxCubes[r] {
						isValid = false
						break
					}
				} else if color == "green" {
					if quantity > maxCubes[g] {
						isValid = false
						break
					}
				} else if color == "blue" {
					if quantity > maxCubes[b] {
						isValid = false
						break
					}
				} else {panic(color)}
			}
			if !isValid {
				break
			}
		}
		
		if isValid {
			s2 := strings.Split(strings.TrimSpace(s1[0]), " ")
			index, err := strconv.Atoi(s2[1])
			if err != nil {
				panic(err)
			}
			answer += index
		}
	}

	return strconv.Itoa(answer)
}

func partTwo() string {
	answer := 0

	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, "Day2", "input.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		maxCubes := [3]int{0,0,0}

		// Determine sum of each cube
		s1 := strings.Split(line, ":")
		rounds := strings.Split(s1[1], ";")
		for _, round := range rounds {
			roundData := strings.Split(round, ",")
			for _, roundDatum := range roundData {
				roundArgs := strings.Split(strings.TrimSpace(roundDatum), " ")
				quantity, err := strconv.Atoi(roundArgs[0])
				if err != nil {
					panic(err)
				}
					
				color := roundArgs[1]
				if color == "red" {
					if quantity > maxCubes[r] {
						maxCubes[r] = quantity
					}
				} else if color == "green" {
					if quantity > maxCubes[g] {
						maxCubes[g] = quantity
					}
				} else if color == "blue" {
					if quantity > maxCubes[b] {
						maxCubes[b] = quantity
					}
				} else {panic(color)}
			}
		}
		
		power := 1
		for _, m := range maxCubes {
			power *= m
		}
		answer += power
		
	}

	return strconv.Itoa(answer)
}

func main() {
	fmt.Println(partTwo())
}
