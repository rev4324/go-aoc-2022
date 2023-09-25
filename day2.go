package main

import (
	"fmt"
	"strings"
)

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func move(first string, second string) int {
	switch second {
	case "X":
		if first == "A" {
			return scores["C"]
		}
		if first == "B" {
			return scores["A"]
		}
		if first == "C" {
			return scores["B"]
		}
	case "Y":
		// draw
		return 3 + scores[first]
	case "Z":
		// win
		if first == "A" {
			return 6 + scores["B"]
		}
		if first == "B" {
			return 6 + scores["C"]
		}
		if first == "C" {
			return 6 + scores["A"]
		}
	}
	return 0
}

func Day2() string {
	scanner := FileScanner()

	total := 0
	for scanner.Scan() {
		str := scanner.Text()
		letters := strings.Split(str, " ")

		total += move(letters[0], letters[1])
	}

	fmt.Printf("total: %v\n", total)
	return ""
}
