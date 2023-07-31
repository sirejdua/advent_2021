package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	commaIdx := bytes.IndexByte(data, ',')
	if commaIdx > 0 {
		buffer := data[:commaIdx]
		return commaIdx + 1, bytes.TrimSpace(buffer), nil
	}
	if atEOF {
		if len(data) > 0 {
			return len(data), bytes.TrimSpace(data), nil
		}
	}
	return 0, nil, nil
}

type State struct {
	data [9]int
	day  int
}

func (s *State) step() {
	var s_next State
	s_next.day = s.day + 1
	for i := 0; i < 9; i++ {
		s_next.data[(i-1+9)%9] = s.data[i]
	}
	s_next.data[6] += s.data[0]

	*s = s_next
}

func (s State) numFish() int {
	sum := 0
	for _, f := range s.data {
		sum += f
	}
	return sum
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(scanCommas)

	idx := 0
	var s State
	for scanner.Scan() {

		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if x < 0 || x >= len(s.data) {
			log.Fatal("x is out of range", x)
		}
		s.data[x]++
		idx++
	}

	fmt.Println(s)
	for i := 0; i < 256; i++ {
		s.step()
		fmt.Println(s, s.numFish())
	}
}
