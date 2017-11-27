package board

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Board struct {
	adj     map[int]map[int]bool
	runes   []rune
	n       int
	visited []bool
}

func (board *Board) Read(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		if s, err := strconv.Atoi(scanner.Text()); err == nil {
			board.n = s
		} else {
			log.Fatal(err)
		}
	}

	board.adj = make(map[int]map[int]bool)
	board.runes = make([]rune, board.n*board.n)
	board.visited = make([]bool, board.n*board.n)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for i := 0; i < board.n; i++ {
		scanner.Scan()
		line := []rune(scanner.Text())

		for j := 0; j < board.n; j++ {
			point := Point{x: i, y: j, n: board.n}
			board.runes[point.Vertex()] = line[j]
			for _, neighbor := range point.Neighbors() {
				board.addEdge(point.Vertex(), neighbor.Vertex())
			}
		}
	}
}

func (board *Board) addEdge(u int, v int) {
	if board.adj[u] == nil {
		board.adj[u] = make(map[int]bool)
	}

	board.adj[u][v] = true

	if board.adj[v] == nil {
		board.adj[v] = make(map[int]bool)
	}

	board.adj[v][u] = true
}

func (board *Board) lookup(trie *Trie, word []rune, u int) {
	newWord := append(word, board.runes[u])

	if !trie.StartWith(newWord) {
		return
	}
	if len(newWord) >= 3 && trie.Contains(newWord) {
		fmt.Println(string(newWord))
	}
	board.visited[u] = true
	for v := range board.adj[u] {
		if board.visited[v] {
			continue
		}
		board.lookup(trie, newWord, v)
	}
	board.visited[u] = false
}

func (board *Board) Search(words [][]rune) {
	trie := &Trie{}

	for _, word := range words {
		trie.AddString(word)
	}

	for _, m := range board.adj {
		for u := range m {
			board.lookup(trie, []rune(""), u)
		}
	}
}
