package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	_, err := readfile("test.txt")
	if err != nil {
		t.Error("file read")
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "part 1 sample",
			input: "test.txt",
			want:  24000,
		},
		{
			name:  "part 1 real",
			input: "input.txt",
			want:  69626,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs, err := readfile(tt.input)
			if err != nil {
				t.Error(err)
			}

			elves := parseFile(fs)
			if got := part1(elves); got != tt.want {
				t.Errorf("part1() = %v, want %v\n", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "part 2 sample",
			input: "test.txt",
			want:  45000,
		},
		{
			name:  "part 2 real",
			input: "input.txt",
			want:  206780,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs, err := readfile(tt.input)
			if err != nil {
				t.Error(err)
			}

			elves := parseFile(fs)
			if got := part2(elves); got != tt.want {
				t.Errorf("part2() = %v, want %v\n", got, tt.want)
			}
		})
	}
}
