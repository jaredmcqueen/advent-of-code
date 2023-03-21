package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestReadFile(t *testing.T) {
	filename := "sample.txt"

	got := readFile(filename)
	if len(got) <= 1 {
		t.Error("readfile", got)
	}
}

func TestParseText(t *testing.T) {
	data := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
	want := [][]string{
		{
			"vJrwpWtwJgWr", "hcsFMMfFFhFp",
		},
		{
			"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL",
		},
	}

	got := ParseText(data)

	for i := 0; i < len(want); i++ {
		if !slices.Equal(got[i], want[i]) {
			t.Errorf("ParseText() want: %+v, got: %+v\n", want, got)
		}
	}
}

func TestPart1(t *testing.T) {
	rucksacks := [][]string{
		{
			"vJrwpWtwJgWr", "hcsFMMfFFhFp",
		},
		{
			"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL",
		},
	}

	answer := Part1(rucksacks)
}
