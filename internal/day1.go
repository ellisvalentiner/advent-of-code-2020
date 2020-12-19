package internal

import (
	"bufio"
	"path/filepath"
	"os"
	"strconv"
)

const target int = 2020

func Day1() (int, int) {
	basepath, _ := os.Getwd()
	absPath := filepath.Join(basepath, "data/input1.txt")

	file, _ := os.Open(absPath)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			lines = append(lines, scanner.Text())
	}

	var part1 int
	var part2 int

	for i := uint64(0); i < uint64(len(lines) - 2); i++ {
		for j := uint64(i + 1); j < uint64(len(lines) - 1); j++ {
			for k := uint64(j + 1); k < uint64(len(lines)); k++ {
				x, _ := strconv.Atoi(lines[i])
				y, _ := strconv.Atoi(lines[j])
				z, _ := strconv.Atoi(lines[k])
				if x + y == target {
					part1 = x * y
					break
				}
				if x + y + z == target {
					part2 = x * y * z
					break
				}
			}
		}
	}

	return part1, part2
}
