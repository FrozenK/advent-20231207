package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    []string
	bid      int
	strength int
}

// By is the type of a "less" function that defines the ordering of its Hand arguments.
type By func(p1, p2 *Hand) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(hands []Hand) {
	ps := &handSorter{
		hands: hands,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// handSorter joins a By function and a slice of Hands to be sorted.
type handSorter struct {
	hands []Hand
	by    func(p1, p2 *Hand) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *handSorter) Len() int {
	return len(s.hands)
}

// Swap is part of sort.Interface.
func (s *handSorter) Swap(i, j int) {
	s.hands[i], s.hands[j] = s.hands[j], s.hands[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *handSorter) Less(i, j int) bool {
	return s.by(&s.hands[i], &s.hands[j])
}

func getStrength(h []string) int {
	/*
	   0 = Five of a kind, where all five cards have the same label: AAAAA
	   1 = Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	   2 = Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	   3 = Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	   4 = Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	   5 = One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	   6 = High card, where all cards' labels are distinct: 23456
	   7 = other
	*/
	cardCount := make(map[string]int)
	for _, c := range h {
		cardCount[c] += 1
	}

	for c, count := range cardCount {
		if count == 5 {
			return 0
		}
		if count == 4 {
			return 1
		}
		if count == 3 {
			for x, xcount := range cardCount {
				if x != c && xcount == 2 {
					return 2
				}
			}
			return 3
		}
		if count == 2 {
			for x, xcount := range cardCount {
				if x != c && xcount == 3 {
					return 2
				}
				if x != c && xcount == 2 {
					return 4
				}
			}
			return 5
		}
	}

	// check if all card are distincts
	for _, count := range cardCount {
		if count > 1 {
			return 7
		}
	}
	return 6
}

func main() {
	f, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var hands []Hand
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, Hand{
			cards:    strings.Split(fields[0], ""),
			bid:      bid,
			strength: getStrength(strings.Split(fields[0], "")),
		})
	}

	value := func(p1, p2 *Hand) bool {
		if p1.strength < p2.strength {
			return false
		}
		if p1.strength > p2.strength {
			return true
		}

		values := []string{
			"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2",
		}
		for i := 0; i < 5; i++ {
			if p1.cards[i] == p2.cards[i] {
				continue
			}

			if slices.Index(values, p1.cards[i]) < slices.Index(values, p2.cards[i]) {
				return false
			}
			return true
		}
		return true
	}
	By(value).Sort(hands)

	sum := 0
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}
	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
