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
		t := strings.Trim(lines[i], "R")
		t = strings.Trim(t, "L")
		a, _ := strconv.Atoi(t)
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
		t := strings.Trim(lines[i], "R")
		t = strings.Trim(t, "L")
		a, _ := strconv.Atoi(t)
		for range a {
			if left {
				curr -= 1
			} else {
				curr += 1
			}
			if curr%100 == 0 {
				cnt += 1
			}

		}
	}
	fmt.Printf("Answer to day 1 part 2: %d\n", cnt)
}
