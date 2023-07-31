package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  string
	size  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data string, depth int) {
	t.root = insert_helper(t.root, data, depth)
}

func insert_helper(n *BinaryNode, data string, depth int) *BinaryNode {
	if n == nil {
		n = &BinaryNode{left: nil, right: nil, data: "", size: 0}
	}
	n.insert(data, depth)
	return n
}

func (n *BinaryNode) insert(data string, depth int) {
	if depth == len(data) {
		n.data = data
		n.size = 1
		return
	}
	if data[depth] == '0' {
		n.left = insert_helper(n.left, data, depth+1)
	} else {
		n.right = insert_helper(n.right, data, depth+1)
	}
}

func (n *BinaryNode) update_sizes() {
	if n.left != nil {
		n.left.update_sizes()
		n.size += n.left.size
	}
	if n.right != nil {
		n.right.update_sizes()
		n.size += n.right.size
	}
}

func (n *BinaryNode) find_rating(greater bool) string {
	if n.left == nil && n.right == nil {
		return n.data
	}
	if n.left == nil {
		return n.right.find_rating(greater)
	}
	if n.right == nil {
		return n.left.find_rating(greater)
	}
	if n.left.size == n.right.size {
		if greater {
			return n.right.find_rating(greater)
		} else {
			return n.left.find_rating(greater)
		}
	}
	if n.left.size > n.right.size == greater {
		return n.left.find_rating(greater)
	}

	return n.right.find_rating(greater)
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v;%v\n", ch, node.data, node.size)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	tree := &BinaryTree{}

	num_lines := 0
	for scanner.Scan() {
		line := scanner.Text()
		tree.insert(line, 0)
		num_lines++
	}

	tree.root.update_sizes()
	print(os.Stdout, tree.root, 1, 'M')
	// Now find the path towards greater and path towards less
	//
	// Greater than rules:
	// go towards the side with greater size
	// tiebreak to the right
	//
	// Less than rules:
	// go towards the side with less size
	// tiebreak to the left
	//
	g, err := strconv.ParseInt(string(tree.root.find_rating(true)), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(string(tree.root.find_rating(false)), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g * e)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
