package main

import (
	"fmt"
	"os"
	"strings"
)

type RingBuffer struct {
	v    []string
	size int
}

func main() {
	data, _ := os.ReadFile("input.txt")
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		rb := NewRingBuffer(14)
		for i, char := range string(line) {
			rb.AddChar(string(char))
			if rb.IsUnique() {
				fmt.Println(rb, rb.IsUnique(), i+1)
				break
			}
		}
	}
}

func NewRingBuffer(size int) *RingBuffer {
	r := RingBuffer{
		v:    []string{},
		size: size,
	}
	return &r
}

func (r *RingBuffer) AddChar(char string) {
	if len(r.v) < r.size {
		r.v = append(r.v, char)
		return
	}
	r.v = append(r.v[1:], char)
}

func (r *RingBuffer) String() string {
	return strings.Join(r.v, "")
}

func (r *RingBuffer) IsUnique() bool {
	// mjqj
	found := true
	if len(r.v) < r.size {
		return false
	}
loop:
	for i, l1 := range r.v {
		for j, l2 := range r.v {
			if j == i {
				continue
			}
			if l2 == l1 {
				found = false
				break loop
			}
		}
	}
	return found
}
