package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseInput(fileName string) {
	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { 
		// Parse the file
	}
}

func partOne(fileName string) int {
	answer := 0

	return answer
}

func partTwo(fileName string) int {
	answer := 0

	return answer
}

func main() {
	fmt.Println(partOne("test.txt")) 

	fmt.Println(partTwo("test.txt"))
}
