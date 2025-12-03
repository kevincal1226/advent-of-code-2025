package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(filename string) {
	f, _ := os.Open(filename)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		biggestBoy, _ := strconv.Atoi(line[0:1])
		best := 0
		for i, c := range line {
			if i == 0 {
				continue
			}
			s, _ := strconv.Atoi(string(c))
			best = max(best, biggestBoy*10+s)
			if biggestBoy < s {
				biggestBoy = s
			}
		}
		sum += best
	}

	fmt.Println(sum)
}

func part2(filename string) {
	f, _ := os.Open(filename)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		stack := []int{}
		lineLen := len(line)
		for i, c := range line {
			s, _ := strconv.Atoi(string(c))
			for len(stack) > 0 && len(stack)+(lineLen-i-1) >= 12 && stack[len(stack)-1] < s {
				stack = stack[0 : len(stack)-1]
			}
			if len(stack) < 12 {
				stack = append(stack, s)
			}
		}
		best := 0
		for _, i := range stack {
			best = best*10 + i
		}
		sum += best
	}

	fmt.Println(sum)
}

func main() {
	part1(os.Args[1])
	part2(os.Args[1])
}
