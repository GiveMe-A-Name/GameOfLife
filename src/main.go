package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	newUniverse := make([][]bool, 0, height)
	for i := 0; i < height; i++ {
		newUniverse = append(newUniverse, make([]bool, width))
	}
	return newUniverse
}

func (universe Universe) Show() {
	fmt.Print("\x1bc")
	for _, row := range universe {
		for _, unit := range row {
			if unit {
				print("*")
			} else {
				print(" ")
			}
		}
		fmt.Println("")
	}
}

func (universe Universe) Seed() {
	for i := 0; i < width*height*0.25; i++ {
		var x, y = rand.Intn(height), rand.Intn(width)
		universe[x][y] = true
	}
}

func (universe Universe) Alive(x, y int) bool {
	x = (x + height) % height
	y = (y + width) % width
	return universe[x][y]
}

func (universe Universe) Neighbors(x, y int) int {
	neighbors := 0
	directions := [3]int{0, 1, -1}
	for _, incremental_x := range directions {
		for _, incremental_y := range directions {
			if !(incremental_x == 0 && incremental_y == 0) {
				neighbors += Btoi(universe.Alive(x+incremental_x, y+incremental_y))
			}
		}
	}
	return neighbors
}

func (universe Universe) Next(x, y int) bool {
	neighbors := universe.Neighbors(x, y)
	if universe[x][y] {
		switch {
		case neighbors < 2:
			return false
		case neighbors > 3:
			return false
		default:
			return true
		}
	} else {
		switch {
		case neighbors == 3:
			return true
		default:
			return false
		}
	}
}
func Step(a Universe) {
	b := NewUniverse()
	for i := range a {
		copy(b[i], a[i])
	}
	for i, row := range a {
		for j := range row {
			b[i][j] = a.Next(i, j)
		}
	}
	copy(a, b)
	a.Show()
}

func main() {
	universe := NewUniverse()
	universe.Seed()

	for {
		Step(universe)
		time.Sleep(500 * time.Millisecond)
	}
}

func Btoi(flag bool) int {
	if flag {
		return 1
	}
	return 0
}
