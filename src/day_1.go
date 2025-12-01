package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day1Part1(filename string) {
	file, _ := os.ReadFile(filename)
	data := string(file)

	lines := strings.Split(data, "\n")

	curr := 50
	cnt := 0

	for i := range len(lines) {
		if len(lines[i]) == 0 {
			continue
		}
		left := lines[i][0] == 'L'
		a, _ := strconv.Atoi(lines[i][1:])
		if left {
			curr -= a
		} else {
			curr += a
		}
		if curr%100 == 0 {
			cnt += 1
		}

	}
	fmt.Printf("Answer to day 1 part 1: %d\n", cnt)
}

func day1Part2(filename string) {
	file, _ := os.ReadFile(filename)
	data := string(file)

	lines := strings.Split(data, "\n")

	curr := 50
	cnt := 0

	for i := range len(lines) {
		if len(lines[i]) == 0 {
			continue
		}

		left := lines[i][0] == 'L'
		difference, _ := strconv.Atoi(lines[i][1:])
		prev := curr

		// left and negative: should help you by the -x mod 100
		// left and positive: should help you by the -x mod 100
		// right and positive: should help you by the x mod 100
		// right and negative: should not help you by the x mod 100

		offset := 0
		if left {
			offset = (-prev%100 + 100) % 100
			curr -= difference
		} else {
			offset = (prev%100 + 100) % 100
			curr += difference
		}

		difference += offset

		cnt += difference / 100

	}
	fmt.Printf("Answer to day 1 part 2: %d\n", cnt)
}
