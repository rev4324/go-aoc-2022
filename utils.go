package main

import (
	"bufio"
	"os"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func FileScanner() *bufio.Scanner {
	if len(os.Args) < 2 {
		panic("specify a filepath")
	}
	fd, err := os.Open(os.Args[1])
	Check(err)

	return bufio.NewScanner(fd)
}
