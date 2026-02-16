package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tetromino struct {
	coords [][]int
	width  int
	height int
	id     rune
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	tetrominoes, err := parseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	board := solve(tetrominoes)
	if board != nil {
		printBoard(board)
	} else {
		fmt.Println("ERROR")
	}
}

func parseFile(filename string) ([]Tetromino, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tetrominoes []Tetromino
	scanner := bufio.NewScanner(file)
	var block []string
	id := 'A'
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if len(block) == 4 {
				t, err := parseBlock(block, rune(id))
				if err != nil {
					return nil, err
				}
				tetrominoes, id = append(tetrominoes, t), id+1
				block = nil
			} else if len(block) > 0 {
				return nil, fmt.Errorf("err")
			}
			continue
		}
		if len(line) != 4 {
			return nil, fmt.Errorf("err")
		}
		block = append(block, line)
	}
	if len(block) == 4 {
		t, err := parseBlock(block, rune(id))
		if err != nil {
			return nil, err
		}
		tetrominoes = append(tetrominoes, t)
	} else if len(block) > 0 {
		return nil, fmt.Errorf("err")
	}
	if err := scanner.Err(); err != nil || len(tetrominoes) == 0 {
		return nil, fmt.Errorf("err")
	}
	return tetrominoes, nil
}

func parseBlock(block []string, id rune) (Tetromino, error) {
	var coords [][]int
	for r, line := range block {
		for c, char := range line {
			if char == '#' {
				coords = append(coords, []int{r, c})
			} else if char != '.' {
				return Tetromino{}, fmt.Errorf("err")
			}
		}
	}
	if len(coords) != 4 || !isConnected(coords) {
		return Tetromino{}, fmt.Errorf("err")
	}
	minR, minC := 4, 4
	maxR, maxC := 0, 0
	for _, coord := range coords {
		if coord[0] < minR {
			minR = coord[0]
		}
		if coord[1] < minC {
			minC = coord[1]
		}
	}
	for i := range coords {
		coords[i][0] -= minR
		coords[i][1] -= minC
		if coords[i][0] > maxR {
			maxR = coords[i][0]
		}
		if coords[i][1] > maxC {
			maxC = coords[i][1]
		}
	}
	return Tetromino{coords: coords, width: maxC + 1, height: maxR + 1, id: id}, nil
}

func isConnected(coords [][]int) bool {
	connections := 0
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if abs(coords[i][0]-coords[j][0])+abs(coords[i][1]-coords[j][1]) == 1 {
				connections++
			}
		}
	}
	return connections >= 3
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(pieces []Tetromino) [][]rune {
	size := 1
	for size*size < len(pieces)*4 {
		size++
	}
	for {
		board := make([][]rune, size)
		for i := range board {
			board[i] = make([]rune, size)
		}
		if backtrack(board, pieces, 0) {
			return board
		}
		size++
		if size > 16 {
			return nil
		}
	}
}

func backtrack(board [][]rune, pieces []Tetromino, index int) bool {
	if index == len(pieces) {
		return true
	}
	p := pieces[index]
	size := len(board)
	for r := 0; r <= size-p.height; r++ {
		for c := 0; c <= size-p.width; c++ {
			if canPlace(board, p, r, c) {
				place(board, p, r, c)
				if backtrack(board, pieces, index+1) {
					return true
				}
				remove(board, p, r, c)
			}
		}
	}
	return false
}

func canPlace(board [][]rune, p Tetromino, r, c int) bool {
	for _, coord := range p.coords {
		if board[r+coord[0]][c+coord[1]] != 0 {
			return false
		}
	}
	return true
}

func place(board [][]rune, p Tetromino, r, c int) {
	for _, coord := range p.coords {
		board[r+coord[0]][c+coord[1]] = p.id
	}
}

func remove(board [][]rune, p Tetromino, r, c int) {
	for _, coord := range p.coords {
		board[r+coord[0]][c+coord[1]] = 0
	}
}

func printBoard(board [][]rune) {
	if board == nil {
		return
	}
	for _, row := range board {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", cell)
			}
		}
		fmt.Println()
	}
}
