package main

import (
	"math/rand"
	"log"
	"time"
)

const (
	LeftWall  = iota
	RightWall
	UpWall
	DownWall
)

type Block struct {
	walls   [4]bool
	Visited bool
}

type Location struct {
	row, col int
}

type Maze struct {
	rows, cols int
	blocks     [][]Block
}

func (m *Maze) visit(row, col int, wall int) {
	m.blocks[row][col].Visited = true
	m.blocks[row][col].walls[wall] = true
}

func (m *Maze) VisitTo(loc Location, wall int) Location {
	row := loc.row
	col := loc.col
	m.blocks[row][col].walls[wall] = false
	newRow := row
	newCol := col
	var newWall int
	switch wall {
	case LeftWall:
		newCol -= 1
		newWall = RightWall
	case RightWall:
		newCol += 1
		newWall = LeftWall
	case UpWall:
		newRow -= 1
		newWall = DownWall
	case DownWall:
		newRow += 1
		newWall = UpWall
	}
	m.visit(newRow, newCol, newWall)
	return Location{newRow, newCol}
}

func (m *Maze) getBlock(row, col int) *Block {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return nil
	}
	return &m.blocks[row][col]
}

func (m *Maze) GetNextBlockByWall(row, col, wall int) *Block {
	switch wall {
	case LeftWall:
		col -= 1
	case RightWall:
		col += 1
	case UpWall:
		row -= 1
	case DownWall:
		row += 1
	}
	return m.getBlock(row, col)
}

func (m *Maze) GetAvailableLocation(row, col int) []int {
	result := make([]int, 0, 4)
	for i := 0; i < 4; i ++ {
		block := m.GetNextBlockByWall(row, col, i)
		if block != nil && !block.Visited {
			result = append(result, i)
		}
	}
	return result
}

func newMaze(row, col int) *Maze {
	maze := &Maze{rows: row, cols: col}
	maze.blocks = make([][]Block, row, row)
	for i := 0; i < row; i ++ {
		maze.blocks[i] = make([]Block, col, col)
		for j := 0; j < col; j ++ {
			maze.blocks[i][j] = Block{
				walls: [4]bool{true, true, true, true},
			}
		}
	}
	maze.blocks[0][0].Visited = true
	return maze
}

func main() {
	maze := newMaze(10, 10)
	path := []Location{{0, 0}}
	rand.Seed(time.Now().Unix())
	count := 1
	for len(path) > 0 {
		curLocation := path[len(path)-1]
		availableWalls := maze.GetAvailableLocation(curLocation.row, curLocation.col)
		if len(availableWalls) == 0 {
			path = path[:len(path)-1]
			continue
		}

		// 选出下一个位置
		index := rand.Intn(len(availableWalls))
		wall := availableWalls[index]
		nextLocation := maze.VisitTo(curLocation, wall)
		path = append(path, nextLocation)
		log.Printf("%v -> %v", curLocation, nextLocation)
		count += 1
	}
	log.Println(count)
}
