package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	text := readFile("input.txt")
	rucksacks := ParseText(text)
	score := Part1(rucksacks)
	fmt.Println("part 1 score is", score)

	score = Part2(rucksacks)
	fmt.Println("part 2 score is", score)
}

func findDuplicateLetterThree(rucksack []string) rune {
	e1 := rucksack[0]
	e2 := rucksack[1]
	e3 := rucksack[2]
	var duplicateLetter rune
out:
	for _, e1letter := range e1 {
		for _, e2letter := range e2 {
			for _, e3letter := range e3 {
				if e1letter == e2letter && e1letter == e3letter {
					duplicateLetter = e1letter
					break out
				}
			}
		}
	}
	return duplicateLetter
}

func findDuplicateLetterTwo(rucksack []string) rune {
	c1 := rucksack[0]
	c2 := rucksack[1]
	var duplicateLetter rune
out:
	for _, c1letter := range c1 {
		for _, c2letter := range c2 {
			if c1letter == c2letter {
				duplicateLetter = c1letter
				break out
			}
		}
	}
	return duplicateLetter
}

func Part1(rucksacks [][]string) int {
	priority := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priorityMap := make(map[rune]int)
	for i, v := range priority {
		priorityMap[v] = i + 1
	}

	score := 0

	for _, rucksack := range rucksacks {
		score += priorityMap[findDuplicateLetterTwo(rucksack)]
	}

	return score
}

func Part2(rucksacks [][]string) int {
	priority := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priorityMap := make(map[rune]int)
	for i, v := range priority {
		priorityMap[v] = i + 1
	}

	score := 0

	var trio []string
	for _, rucksack := range rucksacks {
		trio = append(trio, strings.Join(rucksack, ""))
		if len(trio) == 3 {
			score += priorityMap[findDuplicateLetterThree(trio)]
			trio = []string{}
		}
	}

	return score
}

func readFile(filename string) string {
	fb, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("readFile", err)
	}

	return string(fb)
}

func ParseText(text string) [][]string {
	var rucksacks [][]string
	for _, rucksack := range strings.Split(text, "\n") {

		if rucksack == "" {
			continue
		}
		rucksacks = append(rucksacks, []string{
			rucksack[:len(rucksack)/2],
			rucksack[len(rucksack)/2:],
		})
	}
	return rucksacks
}
