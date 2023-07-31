package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func (p1 Point) eq(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

type Line struct {
	p1        Point
	p2        Point
	direction Point
}

type LineIterator struct {
	l    Line
	p    Point
	last bool
}

// what if the line has p1 == p2
func (l Line) iter() LineIterator {
	return LineIterator{l, l.p1, false}
}

func (li *LineIterator) next() {
	if li.p.eq(li.l.p2) {
		li.last = true
		return
	}

	li.p.x += li.l.direction.x
	li.p.y += li.l.direction.y
}

func (li LineIterator) hasNext() bool {
	return !li.last
}

type Grid struct {
	occupancy map[Point]int
}

func axis_aligned(p1, p2 Point) bool {
	return p1.x == p2.x || p1.y == p2.y
}

func sign(x, y int) int {
	if x > y {
		return 1
	} else if y > x {
		return -1
	} else {
		return 0
	}
}

func makeLine(p1, p2 Point) Line {
	return Line{p1, p2, Point{sign(p2.x, p1.x), sign(p2.y, p1.y)}}
}

func (g *Grid) insert_line(p1, p2 Point) {
	fmt.Fprintf(os.Stdout,
		"%v,%v -> %v,%v\n",
		p1.x, p1.y, p2.x, p2.y)

	if g.occupancy == nil {
		g.occupancy = make(map[Point]int)
	}

	l := makeLine(p1, p2)
	for li := l.iter(); li.hasNext(); li.next() {
		g.occupancy[li.p]++
	}
}

func parse_point(p string) Point {
	coords := strings.Split(p, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{x, y}
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var g Grid

	idx := 0
	for scanner.Scan() {
		scanner.Text()
		p1 := parse_point(scanner.Text())
		scanner.Scan()
		scanner.Scan()
		p2 := parse_point(scanner.Text())

		g.insert_line(p1, p2)

		idx++
	}
	overlap := 0
	for _, count := range g.occupancy {
		if count >= 2 {
			overlap++
		}
	}

	fmt.Println("Number of points where there are at least two lines overlapping is", overlap)
}
