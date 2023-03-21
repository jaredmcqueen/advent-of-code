package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock     int = 1
	Paper    int = 2
	Scissors int = 3

	Lose int = 1
	Draw int = 2
	Win  int = 3
)

var lookupTablePart1 = map[string]int{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var lookupTablePart2 = map[string]int{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
	"A": Lose,
	"B": Draw,
	"C": Win,
}

var (
	scoreLose int = 0
	scoreDraw int = 3
	scoreWin  int = 6
)

func main() {
	f, _ := readFile("sample.txt")

	// part1
	game := parseInput(f, lookupTablePart1)
	score := part1(game)
	fmt.Println("part1 score", score)

	// part2
	game = parseInput(f, lookupTablePart2)
	score = part2(game)
	fmt.Println("part2 score", score)
}

func part1(games [][]int) int {
	totalScore := 0

	for _, game := range games {

		if game[0] == Rock {
			if game[1] == Rock {
				totalScore += scoreDraw + Rock
			}
			if game[1] == Paper {
				totalScore += scoreWin + Paper
			}
			if game[1] == Scissors {
				totalScore += scoreLose + Scissors
			}
		}

		if game[0] == Paper {
			if game[1] == Rock {
				totalScore += scoreLose + Rock
			}
			if game[1] == Paper {
				totalScore += scoreDraw + Paper
			}
			if game[1] == Scissors {
				totalScore += scoreWin + Scissors
			}
		}

		if game[0] == Scissors {
			if game[1] == Rock {
				totalScore += scoreWin + Rock
			}
			if game[1] == Paper {
				totalScore += scoreLose + Paper
			}
			if game[1] == Scissors {
				totalScore += scoreDraw + Scissors
			}
		}
	}

	return totalScore
}

func part2(games [][]int) int {
	totalScore := 0

	for _, game := range games {

		opponentMove := game[0]
		myMove := game[1]

		if myMove == Lose {
			if opponentMove == Rock {
				totalScore += scoreLose + Scissors
			}
			if opponentMove == Paper {
				totalScore += scoreLose + Rock
			}
			if opponentMove == Scissors {
				totalScore += scoreLose + Paper
			}
		}

		if myMove == Draw {
			if opponentMove == Rock {
				totalScore += scoreDraw + Rock
			}
			if opponentMove == Paper {
				totalScore += scoreDraw + Paper
			}
			if opponentMove == Scissors {
				totalScore += scoreDraw + Scissors
			}
		}
		if myMove == Win {
			if opponentMove == Rock {
				totalScore += scoreWin + Paper
			}
			if opponentMove == Paper {
				totalScore += scoreWin + Scissors
			}
			if opponentMove == Scissors {
				totalScore += scoreWin + Rock
			}
		}
	}

	return totalScore
}

func readFile(file string) (string, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func parseInput(text string, lt map[string]int) [][]int {
	var games [][]int
	for _, game := range strings.Split(text, "\n") {
		if game == "" {
			continue
		}
		// "A Y\nB X\nC Z\n"
		theirMove := lt[string(game[0])]
		myMove := lt[string(game[2])]

		games = append(games, []int{theirMove, myMove})
	}

	return games
}
