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
	HIGH_CARD int = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type hand struct {
	cards string
	bid int
}

func parseInput(fileName string) []hand {
	var hands []hand
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
			a := strings.Fields(line)
			bid, err := strconv.Atoi(a[1])
			if err != nil {
				panic(err)
			}
			hands = append(hands, hand{
				cards: a[0],
				bid: bid,
			})
		}
	}
	return hands
}

func calculateScore(hand string, jokers bool) int {
	runeCounts := make(map[rune]int)
	// Determine type of hand
	for _, c := range hand {
		if _, ok := runeCounts[c]; ok {
			runeCounts[c]++
		} else {
			runeCounts[c] = 1
		}
	}
	highestCount := 0
	secondHighest := 0
	for char, count := range runeCounts {
		// For part 2, consider jokers separately 
		if jokers && char == 'J' {
			continue
		}
		if count > highestCount {
			secondHighest = highestCount
			highestCount = count
		} else if count > secondHighest {
			secondHighest = count
		}
	}
	// For part 2, consider jokers as whatever we have the most of
	if jokers {
		if jokerCount, ok := runeCounts['J']; ok {
			highestCount += jokerCount
		}
	}
	switch highestCount {
	case 5: return FIVE_OF_A_KIND
	case 4: return FOUR_OF_A_KIND
	case 3: 
		if secondHighest == 2 {
			return FULL_HOUSE
		} else {
			return THREE_OF_A_KIND
		}
	case 2:
		if secondHighest == 2 {
			return TWO_PAIR
		} else {
			return ONE_PAIR
		}
	default: return HIGH_CARD
	}
}

func cardToInt(c rune, jokers bool) int {
	// For part 2, consider jokers as the worst individual card
	if jokers && c == 'J' {
		return 0
	}

	switch c {
	case 'T': return 10
	case 'J': return 11
	case 'Q': return 12
	case 'K': return 13
	case 'A': return 14
	default: 
		i, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		return i
	}
}

func tieBreaker(a hand, b hand, jokers bool) int {
	for i := 0; i < len(a.cards); i++ {
		if cardToInt(rune(a.cards[i]), jokers) > cardToInt(rune(b.cards[i]), jokers) {
			return 1
		} else if cardToInt(rune(a.cards[i]), jokers) < cardToInt(rune(b.cards[i]), jokers) {
			return -1
		}
	}
	return 0
}

func solve(fileName string, jokers bool) int {
	answer := 0
	organizedHands := make(map[int][]hand)

	hands := parseInput(fileName)

	for _, h := range hands {
		// Determine a numeric score for the hand (higher is better)
		score := calculateScore(h.cards, jokers)
		// Each score is its own list in the map
		// Each list is in order from worst -> best
		if _, ok := organizedHands[score]; ok {
			for i, prevHand := range organizedHands[score] {
				// If the previous hand was better, this hand goes first
				if tieBreaker(prevHand, h, jokers) > 0 {
					organizedHands[score] = append(organizedHands[score][:i], append([]hand{h}, organizedHands[score][i:]...)...)
					break
				} else if i == len(organizedHands[score])-1 {
					organizedHands[score] = append(organizedHands[score], h)
				}
			}
		} else {
			organizedHands[score] = []hand{h}
		}
	}
	// fmt.Println(organizedHands)


	// In order from worst to best hand, multiply bid by rank
	rank := 1
	for i := 0; i <= FIVE_OF_A_KIND; i++ {
		// fmt.Printf("%d - %d\n", i, len(organizedHands[i]))
		for _, h := range organizedHands[i] {
			// fmt.Printf("%s\n", h.cards)
			answer += (rank * h.bid)
			rank++
		}
	}

	return answer
}

func main() {
	fmt.Println(solve("input.txt", false)) 
	fmt.Println(solve("input.txt", true))
}
