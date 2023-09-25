package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkContains(first, second []string) bool {
	first0, err := strconv.Atoi(first[0])
	Check(err)

	second0, err := strconv.Atoi(second[0])
	Check(err)

	first1, err := strconv.Atoi(first[1])
	Check(err)

	second1, err := strconv.Atoi(second[1])
	Check(err)

	return first0 <= second0 && first1 >= second1
}

func checkContainsPart2(first, second []string) bool {
	first0, err := strconv.Atoi(first[0])
	Check(err)

	second0, err := strconv.Atoi(second[0])
	Check(err)

	first1, err := strconv.Atoi(first[1])
	Check(err)

	second1, err := strconv.Atoi(second[1])
	Check(err)

	for i := first0; i <= first1; i++ {
		if i <= second1 && i >= second0 {
			return true
		}
	}

	return false
}

func Day4() {
	scanner := FileScanner()

	counter := 0
	for scanner.Scan() {
		str := scanner.Text()
		pairs := strings.Split(str, ",")
		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")

		if checkContains(first, second) || checkContains(second, first) {
			counter++
		}
	}

	fmt.Printf("counter: %v\n", counter)
}

func Day4Part2() {
	scanner := FileScanner()

	counter := 0
	for scanner.Scan() {
		str := scanner.Text()
		pairs := strings.Split(str, ",")
		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")

		if checkContainsPart2(first, second) || checkContainsPart2(second, first) {
			counter++
		}
	}

	fmt.Printf("counter: %v\n", counter)
}
