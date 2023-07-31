package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	m       map[int]int
	numbers [25]int
	marked  [25]bool
	solved  bool
}

func (b *Board) print(w io.Writer) {
	fmt.Fprintln(w, b.m)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Fprintf(w, "%d;%v  ", b.numbers[5*i+j], b.marked[5*i+j])
		}
		fmt.Fprint(w, "\n")
	}
}

func (b *Board) mark(val int) {
	if b.solved {
		return
	}
	if b.m == nil {
		b.m = make(map[int]int)
		for i, v := range b.numbers {
			b.m[v] = i
		}
	}
	if idx, ok := b.m[val]; ok {
		b.marked[idx] = true
	}
}

func (b *Board) isWinning() bool {
	for i := 0; i < 5; i++ {
		row_all_marked := true
		col_all_marked := true
		for j := 0; j < 5; j++ {
			if !b.marked[5*i+j] {
				row_all_marked = false
			}
			if !b.marked[5*j+i] {
				col_all_marked = false
			}
		}
		if row_all_marked || col_all_marked {
			b.solved = true
			return true
		}
	}
	return false
}

func (b *Board) computeScore(scoreMultiplier int) int {
	score := 0
	for val, idx := range b.m {
		if !b.marked[idx] {
			score += val
		}
	}
	score *= scoreMultiplier
	return score
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	called_numbers := strings.Split(scanner.Text(), ",")

	var boards []Board
	idx := 0
	for scanner.Scan() {
		if idx%25 == 0 {
			boards = append(boards, Board{})
		}
		x, _ := strconv.Atoi(scanner.Text())
		boards[len(boards)-1].numbers[idx%25] = x
		idx++
	}

	last_solved := -1
	last_called_for_solve := -1
	// range returns a copy !!
	for _, numString := range called_numbers {
		x, _ := strconv.Atoi(numString)
		for boardIdx := range boards {
			if boards[boardIdx].solved {
				continue
			}
			boards[boardIdx].mark(x)
			if boards[boardIdx].isWinning() {
				fmt.Fprintf(os.Stdout,
					"%v is the winning board.\n%v is the winning score",
					boardIdx, boards[boardIdx].computeScore(x))
				last_solved = boardIdx

				last_called_for_solve = x
			}
		}
	}

	fmt.Fprintf(os.Stdout,
		"%v is the last winning board.\n%v is the winning score",
		last_solved, boards[last_solved].computeScore(last_called_for_solve))
}
