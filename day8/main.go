package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Node struct {
	id string
	left string
	right string
}

const (
	START = "AAA"
	DEST = "ZZZ"
)

func parseInput(fileName string) (string, map[string]Node) {
	// Return values
	var instructions string
	nodeMap := make(map[string]Node)

	wd, err := os.Getwd()
	file, err := os.Open(filepath.Join(wd, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// First read in the instructions
	scanner.Scan()
	instructions = strings.TrimSpace(scanner.Text())

	for scanner.Scan() { 
		line := scanner.Text()
		if len(line) > 0 {
			// Read in a node (assuming IDs is always 3 letters)
			split := strings.Split(line, " ")
			n := Node{
				id: split[0],
				left: split[2][1:4],
				right: split[3][:3],
			}
			nodeMap[n.id] = n
		}
	}

	return instructions, nodeMap
}

func partOne(fileName string) int {
	answer := 0

	instructions, nodeMap := parseInput(fileName)
	currentNode := nodeMap[START]

	for currentNode.id != DEST {
		// Determine direction and increment step counter
		direction := instructions[answer % len(instructions)]
		answer++
		
		// Move to the next node based on direction
		if direction == 'L' {
			currentNode = nodeMap[currentNode.left]
		} else if direction == 'R' {
			currentNode = nodeMap[currentNode.right]
		} else {
			panic(fmt.Sprintf("ERROR - Invalid direction: %c", rune(direction)))
		}
	}

	return answer
}

func allAtZ (nodes []Node) bool {
	for _, n := range nodes {
		if n.id[2] != 'Z' {
			return false
		}
	}
	return true
}

func gcd(a int, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func lcm(a int, b int) int {
	return (a*b)/gcd(a,b)
}

func lcmList(args []int) int {
	var retval int
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}
	retval = lcm(args[0], args[1])
	for i := 2; i < len(args); i++ {
		retval = lcm(retval, args[i])
	}
	return retval
}

func partTwo(fileName string) int {
	var currentNodes []Node

	instructions, nodeMap := parseInput(fileName)

	// Start out at all nodes ending in A
	for id, n := range nodeMap {
		if id[2] == 'A' {
			currentNodes = append(currentNodes, n)
		}
	}

	// Assuming the end always loops directly back to the beginning, you'll arrive at all ends at the least common multiple of all answers
	answers := make([]int, len(currentNodes))

	for i := range currentNodes {
		for currentNodes[i].id[2] != 'Z' {
			// Determine direction and increment step counter
			direction := instructions[answers[i] % len(instructions)]
			answers[i]++

			// Move each node to the next based on direction
			if direction == 'L' {
				currentNodes[i] = nodeMap[currentNodes[i].left]
			} else if direction == 'R' {
				currentNodes[i] = nodeMap[currentNodes[i].right]
			} else {
				panic(fmt.Sprintf("ERROR - Invalid direction: %c", rune(direction)))
			}
		}	
	}

	return lcmList(answers)
}

func main() {
	fmt.Println(partOne("input.txt")) 

	fmt.Println(partTwo("input.txt"))
}
