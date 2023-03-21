package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	ft, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("readFile", err)
	}
	return string(ft)
}

func generateSection(beginning, end int) []int {
	var result []int
	for i := beginning; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func parseLines(text string) [][][]int {
	var data [][][]int
	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}
		var sections [][]int
		for _, section := range strings.Split(line, ",") {
			if section == "" {
				continue
			}

			numbers := strings.Split(section, "-")

			begin, _ := strconv.Atoi(string(numbers[0]))
			end, _ := strconv.Atoi(string(numbers[1]))
			sections = append(sections, generateSection(begin, end))
		}
		data = append(data, sections)
	}
	return data
}

// func main() {
// 	ft := readFile("input.txt")
// 	sections := parseLines(ft)
// 	dupes1 := part1(sections)
// 	fmt.Println("part 1 dupes:", dupes1)
// 	dupes2 := part2(sections)
// 	fmt.Println("part 2 dupes:", dupes2)
// }

func part1(elfPairs [][][]int) int {
	dupes := 0
	for _, elfPair := range elfPairs {
		m := make(map[int]uint8)

		a := elfPair[0]
		b := elfPair[1]

		for _, id := range a {
			m[id] += 1
		}

		for _, id := range b {
			m[id] += 1
		}

		both := 0
		for _, v := range m {
			if v == 2 {
				both += 1
			}
		}

		var min int
		if len(a) < len(b) {
			min = len(a)
		} else {
			min = len(b)
		}

		if both == min {
			dupes += 1
		}

	}
	return dupes
}

func part2(elfPairs [][][]int) int {
	dupes := 0
	for _, elfPair := range elfPairs {
		m := make(map[int]uint8)

		a := elfPair[0]
		b := elfPair[1]

		for _, id := range a {
			m[id] += 1
		}

		for _, id := range b {
			m[id] += 1
		}

		for _, v := range m {
			if v == 2 {
				dupes += 1
				break
			}
		}

	}
	return dupes
}

func main() {
	d, _ := os.ReadFile("input.txt")
	pairs := strings.Split(string(d), "\n")
	c1, c2 := 0, 0
	for _, p := range pairs {
		if p == "" {
			continue
		}
		s1, s2 := strings.Split(p, ",")[0], strings.Split(p, ",")[1]
		start1, _ := strconv.Atoi(strings.Split(s1, "-")[0])
		start2, _ := strconv.Atoi(strings.Split(s2, "-")[0])
		end1, _ := strconv.Atoi(strings.Split(s1, "-")[1])
		end2, _ := strconv.Atoi(strings.Split(s2, "-")[1])

		if (start1 >= start2 && end1 <= end2) || (start2 >= start1 && end2 <= end1) {
			c1++
		}
		if (start1 <= end2) && (start2 <= end1) {
			// if (start1 <= end2) && (end1 >= start2) {
			c2++
		}
	}
	print("Part One: ", c1, " Part Two: ", c2)
}
