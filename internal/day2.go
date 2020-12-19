package internal

import (
	"bufio"
	"index/suffixarray"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func Day2() (int, int) {
	basePath, _ := os.Getwd()
	absPath := filepath.Join(basePath, "data/input2.txt")

	file, _ := os.Open(absPath)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var part1 int
	var part2 int

	re := regexp.MustCompile("(-| |:)+")

	part1 = 0
	part2 = 0
	for i := 0; i < len(lines); i++ {
		split := re.Split(lines[i], -1)
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])
		char := split[2]
		pwd := split[3]
		r := regexp.MustCompile(char)
		count := len(suffixarray.New([]byte(pwd)).FindAllIndex(r, -1))
		if min <= count && count <= max {
			part1++
		}
		if pwd[min-1] == char[0] && pwd[max-1] != char[0] {
			part2++
		}
		if pwd[min-1] != char[0] && pwd[max-1] == char[0] {
			part2++
		}
	}

	return part1, part2
}
