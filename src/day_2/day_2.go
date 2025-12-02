package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(filename string) {
	f, _ := os.Open(filename)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	invalidSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.SplitSeq(line, ",")
		for id := range ids {
			splitIds := strings.Split(id, "-")
			start, _ := strconv.Atoi(splitIds[0])
			end, _ := strconv.Atoi(splitIds[1])
			for i := start; i <= end; i += 1 {
				strI := strconv.Itoa(i)
				if len(strI)%2 == 0 && strI[0:len(strI)/2] == strI[len(strI)/2:] {
					invalidSum += i
				}
			}
		}
	}

	fmt.Println(invalidSum)
}

func part2(filename string) {
	f, _ := os.Open(filename)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	invalidSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.SplitSeq(line, ",")
		for id := range ids {
			splitIds := strings.Split(id, "-")
			start, _ := strconv.Atoi(splitIds[0])
			end, _ := strconv.Atoi(splitIds[1])
			for i := start; i <= end; i += 1 {
				strI := strconv.Itoa(i)
				for j := 1; j <= len(strI)/2; j += 1 {
					part := strI[0:j]
					if strI == strings.Repeat(part, len(strI)/j) {
						invalidSum += i
						break
					}
				}
			}
		}
	}

	fmt.Println(invalidSum)
}

func main() {
	part1(os.Args[1])
	part2(os.Args[1])
}
