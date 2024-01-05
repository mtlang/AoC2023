package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"math"
)

type raceData struct {
	time int
	record int
}

func parseInputOne(fileName string) []raceData{
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the data from the file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeStrings := strings.Fields(scanner.Text())
	scanner.Scan()
	recordStrings := strings.Fields((scanner.Text()))

	// Return the data as a slice of raceData's
	raceRecords := make([]raceData, len(timeStrings)-1)
	for i := 1; i < len(timeStrings); i++ {
		time, err0 := strconv.Atoi(timeStrings[i])
		record, err1 := strconv.Atoi(recordStrings[i])
		if err0 != nil || err1 != nil {
			panic(fmt.Sprintf("%s - %s", timeStrings[i], recordStrings[i]))
		}

		raceRecords[i-1] = raceData{
			time: time,
			record: record,
		}
	}

	return raceRecords
}

func partOne(fileName string) int {
	answer := 1

	// Parse the input data
	raceRecords := parseInputOne(fileName)

	// For each race, determine the number of winning strategies
	for _, raceRecord := range raceRecords {
		numStrategies := 0
		// Test strategies starting with the middle, until one loses
		for i := (raceRecord.time+1)/2; i < raceRecord.time; i++ {
			distance := i * (raceRecord.time-i)
			if distance > raceRecord.record {
				numStrategies++
			} else {
				break
			}
		}
		// True number of strategies is double what we found
		numStrategies *= 2
		// If the race is an even number of seconds, we're off by one
		if raceRecord.time % 2 == 0 {
			numStrategies -= 1
		}

		// Factor this into the final answer
		answer *= numStrategies
	}

	return answer
}

func parseInputTwo(fileName string) raceData{
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the data from the file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeStrings := strings.Fields(scanner.Text())
	scanner.Scan()
	recordStrings := strings.Fields((scanner.Text()))

	// Return the data as a raceData
	time, err0 := strconv.Atoi(strings.Join(timeStrings[1:],""))
	record, err1 := strconv.Atoi(strings.Join(recordStrings[1:],""))
	if err0 != nil || err1 != nil {
		panic(fmt.Sprintf("%s - %s", timeStrings, recordStrings))
	}

	return raceData{
		time: time,
		record: record,
	}
}

func quadratic(a float64, b float64, c float64) []float64{
	// x = (-b +- sqrt((b^2)-4ac))/(2a)
	x1 := (-1*b + math.Sqrt(math.Pow(b, 2) - 4*a*c))/(2*a)
	x2 := (-1*b - math.Sqrt(math.Pow(b, 2) - 4*a*c))/(2*a)

	// To make future logic easier, we'll always return the smaller answer first
	if x1 < x2 {
		return []float64{x1, x2}
	}
	return []float64{x2, x1}
}

func partTwo(fileName string) int {
	answer := 0

	raceRecord := parseInputTwo(fileName)
	
	// Now that we're working with much bigger numbers, we'll be smarter about finding the answer
	// First, we'll determine possible times to charge in order to tie the record
	// This can be done using the quadratic equation:
	// distance = timeCharged * (time - timeCharged)
	// Therfore we can solve for timeCharged in the following form:
	// -tC^2 + tC*t - d
	// AKA a quadratic equation with a = -1, b = time, and c = -(distance)
	timesCharged := quadratic(float64(-1), float64(raceRecord.time), float64(-1 * raceRecord.record))

	// If we round these times down, the number of winning strategies is their difference
	answer = int(timesCharged[1]) - int(timesCharged[0])

	return answer
}

func main() {
	// fmt.Println(partOne("input.txt")) 

	fmt.Println(partTwo("input.txt"))
}
