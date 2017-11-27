package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ethylamide/boggle/board"
)

func main() {
	var words [][]rune

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words = append(words, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	b := &board.Board{}
	b.Read("input.txt")

	b.Search(words)
}
