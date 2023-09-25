package main

import (
	"fmt"
	"strings"
)

func getShared(left string, right string) []byte {
	var result []byte
	for i := 0; i < len(right); i++ {
		if strings.Contains(left, string(right[i])) && !strings.Contains(string(result), string(right[i])) {
			result = append(result, right[i])
		}
	}
	return result
}

func getPriority(val rune) int {
	if int(val) > 96 {
		return int(val) - 96
	} else {
		return int(val) - 38
	}
}

func RemoveIndex(s [3]string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func handleThreeLines(first, second, third string) int {
	strs := [3]string{first, second, third}
	longest := 0
	for i, l := range strs {
		if len(l) > len(strs[longest]) {
			longest = i
		}
	}
	withoutLongest := RemoveIndex(strs, longest)

	var total int
	var shared string
	for _, r := range strs[longest] {
		if strings.Contains(withoutLongest[0], string(r)) && strings.Contains(withoutLongest[1], string(r)) && !strings.Contains(shared, string(r)) {
			shared += string(r)
			total += getPriority(r)
		}
	}

	return total
}

func Day3() {
	scanner := FileScanner()

	var total int
	for scanner.Scan() {
		str := scanner.Text()
		l := len(str)

		left := str[0 : l/2]
		right := str[l/2 : l]
		shared := string(getShared(left, right))

		for _, s := range shared {
			total += getPriority(s)
		}

	}

	fmt.Printf("total: %v\n", total)
}

func Day3PartTwo() {
	scanner := FileScanner()

	var total, lineNo int
	var first, second string
	for scanner.Scan() {
		str := scanner.Text()
		lineNo++

		switch lineNo % 3 {
		case 1:
			first = str
		case 2:
			second = str
		case 0:
			total += handleThreeLines(first, second, str)
		}

	}

	fmt.Printf("total: %v\n", total)
}
