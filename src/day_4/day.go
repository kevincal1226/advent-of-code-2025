package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(filename string) {
	f, _ := os.Open(filename)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := []string{}

	numCols := 0

	for scanner.Scan() {
		text := scanner.Text()

		if len(grid) == 0 {
			numCols = len(text) + 2
			grid = append(grid, strings.Repeat(".", numCols))
		}

		grid = append(grid, "."+text+".")
	}

	grid = append(grid, strings.Repeat(".", numCols))

	numGoodBois := 0

	dirs := [][]int{}

	for i := -1; i <= 1; i += 1 {
		for j := -1; j <= 1; j += 1 {
			if i == j && j == 0 {
				continue
			}

			tmp := []int{i, j}
			dirs = append(dirs, tmp)
		}
	}

	for r, row := range grid {
		if r == 0 || r == len(grid)-1 {
			continue
		}
		for c, cha := range row {
			if c == 0 || c == numCols-1 || cha == '.' {
				continue
			}

			cnt := 0
			for _, dir := range dirs {
				if grid[r+dir[0]][c+dir[1]] == '@' {
					cnt += 1
				}
			}

			if cnt < 4 {
				numGoodBois += 1
			}

		}
	}

	fmt.Println(numGoodBois)
}

func part2(filename string) {
	f, _ := os.Open(filename)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := [][]byte{}

	numCols := 0

	for scanner.Scan() {
		text := scanner.Text()

		if len(grid) == 0 {
			numCols = len(text) + 2
			grid = append(grid, []byte(strings.Repeat(".", numCols)))
		}

		grid = append(grid, []byte("."+text+"."))
	}

	grid = append(grid, []byte(strings.Repeat(".", numCols)))

	numGoodBois := 0
	oldGoodBois := -1

	dirs := [][]int{}

	for i := -1; i <= 1; i += 1 {
		for j := -1; j <= 1; j += 1 {
			if i == j && j == 0 {
				continue
			}

			tmp := []int{i, j}
			dirs = append(dirs, tmp)
		}
	}

	for oldGoodBois != numGoodBois {
		oldGoodBois = numGoodBois

		for r, row := range grid {
			if r == 0 || r == len(grid)-1 {
				continue
			}
			for c, cha := range row {
				if c == 0 || c == numCols-1 || cha == '.' {
					continue
				}

				cnt := 0
				for _, dir := range dirs {
					if grid[r+dir[0]][c+dir[1]] == '@' {
						cnt += 1
					}
				}

				if cnt < 4 {
					numGoodBois += 1
					grid[r][c] = '.'
				}

			}
		}
	}

	fmt.Println(numGoodBois)
}

func main() {
	part1(os.Args[1])
	part2(os.Args[1])
}
