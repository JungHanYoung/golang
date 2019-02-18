package main

import "fmt"

const N int = 8
const IMAGE_COLOR int = 1
const ALREADY_COUNTED int = 2

type Point struct {
	x int
	y int
}

func countCells(point Point, grid [][8]int) int {
	x := point.x
	y := point.y
	if x < 0 || x >= N || y < 0 || y >= N {
		return 0
	} else if grid[x][y] != IMAGE_COLOR {
		return 0
	}

	grid[x][y] = ALREADY_COUNTED
	return 1 + countCells(Point{x: x - 1, y: y - 1}, grid) + countCells(Point{x: x, y: y - 1}, grid) + countCells(Point{x: x + 1, y: y - 1}, grid) + countCells(Point{x: x - 1, y: y}, grid) + countCells(Point{x: x, y: y}, grid) + countCells(Point{x: x + 1, y: y}, grid) + countCells(Point{x: x - 1, y: y + 1}, grid) + countCells(Point{x: x, y: y + 1}, grid) + countCells(Point{x: x + 1, y: y + 1}, grid)

}

func main() {

	grid := [][N]int{
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 1, 0, 0, 1, 0, 0},
		{1, 1, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0},
		{1, 0, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 0, 1, 1, 1},
	}

	point := Point{x: 5, y: 3}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if j == point.x && i == point.y {
				fmt.Print("*")
			} else if grid[i][j] == IMAGE_COLOR {
				fmt.Print("o")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println(len(grid))

	result := countCells(point, grid)

	fmt.Println("result: ", result)

}
