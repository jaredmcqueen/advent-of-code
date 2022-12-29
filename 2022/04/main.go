package main

import (
	"fmt"
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
	for i := beginning; i < end+1; i++ {
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

func main() {
	ft := readFile("input.txt")
	sections := parseLines(ft)
	dupes1 := part1(sections)
	fmt.Println("part 1 dupes:", dupes1)
	dupes2 := part2(sections)
	fmt.Println("part 2 dupes:", dupes2)
}

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
