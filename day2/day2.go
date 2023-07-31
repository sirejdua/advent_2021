package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	depth := 0
	horizontal := 0
	aim := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(words[1])
		switch words[0] {
		case "forward":
			horizontal += x
			depth += aim * x
		case "down":
			aim += x
		case "up":
			aim -= x
		}
	}

	fmt.Println(horizontal * depth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
