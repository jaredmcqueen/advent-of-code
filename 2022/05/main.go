package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) (stacksText, instructionsText string) {
	// split the stacks from the instructions
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("readFile", err)
	}

	stacksText = strings.Split(string(text), "\n\n")[0]
	instructionsText = strings.Split(string(text), "\n\n")[1]
	return stacksText, instructionsText
}

type Stack struct{}

func stacksParser(text string) map[int][]string {
	m := make(map[int][]string)
	for _, line := range strings.Split(text, "\n") {
		for i := 0; i < len(line); i += 4 {
			position := i/4 + 1
			character := line[i+1]
			if (character >= 65) && (character <= 91) {
				m[position] = append(m[position], string(character))
			}
		}
	}
	return m
}

func instructionsParser(text string) [][]int {
	var instructions [][]int

	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}
		var amount, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
		instructions = append(instructions, []int{amount, from, to})
	}

	return instructions
}

func main() {
	filename := "input.txt"

	stacksText, instructionsText := readFile(filename)
	stacks := stacksParser(stacksText)
	instructions := instructionsParser(instructionsText)

	part1Answer := part1(stacks, instructions)
	fmt.Println("part1 answer", part1Answer)

	stacksText, instructionsText = readFile(filename)
	stacks = stacksParser(stacksText)
	instructions = instructionsParser(instructionsText)

	part2Answer := part2(stacks, instructions)
	fmt.Println("part2 answer", part2Answer)
}

func part1(stacks map[int][]string, instructions [][]int) string {
	for _, instruction := range instructions {

		amount := instruction[0]
		from := instruction[1]
		to := instruction[2]

		// grab the letters to move
		lettersToMove := make([]string, amount)
		copy(lettersToMove, stacks[from][:amount])

		// remove from the FROM stack
		stacks[from] = stacks[from][amount:]

		// move the letter
		for _, letter := range lettersToMove {
			stacks[to] = append([]string{letter}, stacks[to]...)
		}
	}

	answer := make([]string, len(stacks))
	for k, v := range stacks {
		answer[k-1] = v[0]
	}

	return strings.Join(answer, "")
}

func part2(stacks map[int][]string, instructions [][]int) string {
	for _, instruction := range instructions {

		amount := instruction[0]
		from := instruction[1]
		to := instruction[2]

		// grab the letters to move
		lettersToMove := make([]string, amount)
		copy(lettersToMove, stacks[from][:amount])

		// remove from the FROM stack
		stacks[from] = stacks[from][amount:]

		// move the letter
		// for _, letter := range lettersToMove {
		stacks[to] = append(lettersToMove, stacks[to]...)
		// }
	}

	answer := make([]string, len(stacks))
	for k, v := range stacks {
		if len(v) > 0 {
			answer[k-1] = v[0]
		} else {
			answer[k-1] = " "
		}
	}

	return strings.Join(answer, "")
}
