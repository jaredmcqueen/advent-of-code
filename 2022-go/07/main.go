package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Directory int = iota + 1
	File
)

type Node struct {
	Tag      int
	Name     string
	Size     int
	Children []*Node
}

func main() {
	// 924098 is too low
	input := readFile("sample.txt")
	answer1 := Part1(input)
	fmt.Println("Part 1 answer is", answer1)
}

func Part1(instructions []string) int {
	root := NewDir("/", Directory)
	pos := root
	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}
		if string(instruction[0]) == "$" {
			var command, arg string
			fmt.Sscanf(instruction, "$ %s %s", &command, &arg)
			if command == "cd" {
				if arg == ".." {
					pos = FindNode(root, pos.Name, Directory)
					continue
				}
				pos = FindNode(root, arg, Directory)
			}
			continue
		}
		left := strings.Split(instruction, " ")[0]
		right := strings.Split(instruction, " ")[1]
		size, err := strconv.Atoi(left)
		if err != nil {
			pos.Children = append(pos.Children, NewDir(right, Directory))
			continue
		}
		pos.Children = append(pos.Children, NewFile(right, size, File))
	}

	result := 0
	r := FindDirs(root)
	sort.Strings(r)
	for _, n := range r {
		fmt.Println(n)
	}

	return result
}

func NewDir(name string, tag int) *Node {
	node := Node{
		Tag:  tag,
		Name: name,
	}
	return &node
}

func NewFile(name string, size, tag int) *Node {
	node := Node{
		Tag:  tag,
		Size: size,
		Name: name,
	}
	return &node
}

func SumDir(root *Node) int {
	result := 0
	queue := make([]*Node, 0) // why zero?
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0] // grab first element
		queue = queue[1:]  // remove first element
		if nextUp.Tag == File {
			result += nextUp.Size
		}
		if len(nextUp.Children) > 0 {
			queue = append(queue, nextUp.Children...)
		}
	}
	return result
}

// depth first
func FindDirs(root *Node) []string {
	result := []string{}
	queue := make([]*Node, 0) // why zero?
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0] // grab first element
		queue = queue[1:]  // remove first element
		if nextUp.Tag == Directory {
			result = append(result, nextUp.Name)
		}
		if len(nextUp.Children) > 0 {
			queue = append(queue, nextUp.Children...)
		}
	}
	return result
}

func FindNode(root *Node, name string, tag int) *Node {
	queue := make([]*Node, 0) // why zero?
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0] // grab first element
		queue = queue[1:]  // remove first element
		if nextUp.Name == name && nextUp.Tag == tag {
			return nextUp
		}
		if len(nextUp.Children) > 0 {
			queue = append(queue, nextUp.Children...)
		}
	}
	return nil
}

func readFile(filename string) []string {
	var result []string
	f, _ := os.ReadFile(filename)
	lines := strings.Split(string(f), "\n")
	result = append(result, lines...)
	return result
}
