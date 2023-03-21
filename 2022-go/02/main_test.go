package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestReadFile(t *testing.T) {
	f, err := readFile("sample.txt")
	if err != nil {
		t.Error("readFile", err)
	}

	if len(f) <= 1 {
		t.Error("file no bytes")
	}
}

func TestParseFilePart1(t *testing.T) {
	sampleInput := "A Y\nB X\nC Z\n"

	want := [][]int{
		{Rock, Paper},
		{Paper, Rock},
		{Scissors, Scissors},
	}

	got := parseInput(sampleInput, lookupTablePart1)

	for i := 0; i < len(want); i++ {
		if !slices.Equal(want[i], got[i]) {
			t.Errorf("parseFile want:%+v got:%+v\n", want, got)
		}
	}
}

func TestParseFilePart2(t *testing.T) {
	sampleInput := "A Y\nB X\nC Z\n"

	want := [][]int{
		{Rock, Draw},
		{Paper, Lose},
		{Scissors, Win},
	}

	got := parseInput(sampleInput, lookupTablePart2)

	for i := 0; i < len(want); i++ {
		if !slices.Equal(want[i], got[i]) {
			t.Errorf("parseFile want:%+v got:%+v\n", want, got)
		}
	}
}

func TestPart1(t *testing.T) {
	game := [][]int{
		{Rock, Paper},
		{Paper, Rock},
		{Scissors, Scissors},
	}

	got := part1(game)
	want := 15

	if got != want {
		t.Errorf("part1() got: %+v, want: %+v\n", got, want)
	}
}

// func TestPart2(t *testing.T) {
// 	game := [][]int{
// 		{Rock, Paper},
// 		{Paper, Rock},
// 		{Scissors, Scissors},
// 	}
// 	got := part2(game, lookupTablePart2)
// 	want := 12
//
// 	if got != want {
// 		t.Errorf("part2() got: %+v, want: %+v\n", got, want)
// 	}
// }
