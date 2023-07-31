package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	num_increases := 0

	var vals []int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		vals = append(vals, i)
	}
	for x := 0; x < len(vals)-2; x++ {
		vals[x] = vals[x] + vals[x+1] + vals[x+2]
	}
	for x := 1; x < len(vals)-2; x++ {
		if vals[x] > vals[x-1] {
			num_increases += 1
		}
	}

	fmt.Println(num_increases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
