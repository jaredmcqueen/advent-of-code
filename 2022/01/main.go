package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fileString, _ := readfile("input.txt")
	elves := parseFile(fileString)
	fmt.Println(elves)
	fmt.Println("elf with most calories has", part1(elves))
	fmt.Println("top 3 elves combined are", part2(elves))
}

func readfile(fn string) (string, error) {
	f, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func parseFile(fs string) (elves [][]int) {
	for _, elf := range strings.Split(fs, "\n\n") {
		if elf == "" {
			continue
		}
		var foodItems []int
		for _, foodItem := range strings.Split(elf, "\n") {
			if foodItem == "" {
				continue
			}
			// fmt.Println(foodItem)
			foodCalories, err := strconv.Atoi(foodItem)
			if err != nil {
				log.Fatal(err)
			}
			foodItems = append(foodItems, foodCalories)
		}
		elves = append(elves, foodItems)
	}
	return elves
}

func sum(array []int) int {
	var total int
	for _, item := range array {
		total += item
	}
	return total
}

func part1(elves [][]int) int {
	totals := []int{}
	for _, elf := range elves {
		totals = append(totals, sum(elf))
	}

	sort.Ints(totals)
	return totals[len(totals)-1]
}

func part2(elves [][]int) int {
	totals := []int{}
	for _, elf := range elves {
		totals = append(totals, sum(elf))
	}

	sort.Ints(totals)
	last3 := totals[len(totals)-3:]
	return sum(last3)
}
