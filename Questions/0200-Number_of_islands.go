package leetcode

import (
	"time"

	elapsedTime "BenWolfaardt/Practice_Coding_Challenges/tree/02-LeetCode-Go/Utilities"
)

// Bredth First Search - is much slower with the timer for the grid size in tests
func numIslandsBFS(grid [][]byte) int {

	defer elapsedTime.Timer(time.Now())

	var result int
	n, m := len(grid), len(grid[0])

	if n == 0 || m == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != '1' {
				continue
			}

			var queue = [][]int{{i, j}}

			for len(queue) > 0 {
				row, column := queue[0][0], queue[0][1]
				queue = queue[1:]

				if row+1 < n && grid[row+1][column] == '1' {
					queue = append(queue, []int{row + 1, column})
				}
				if row-1 >= 0 && grid[row-1][column] == '1' {
					queue = append(queue, []int{row - 1, column})
				}
				if column+1 < m && grid[row][column+1] == '1' {
					queue = append(queue, []int{row, column + 1})
				}
				if column-1 >= 0 && grid[row][column-1] == '1' {
					queue = append(queue, []int{row, column - 1})
				}
				grid[row][column] = 0
			}
			result++
		}
	}
	return result
}

// Depth First Search - is much quicker with the timer for the grid size in tests
func numIslandsDFS(grid [][]byte) int {

	defer elapsedTime.Timer(time.Now())

	if len(grid) == 0 {
		return 0
	}

	nRow, nCol := len(grid), len(grid[0])
	count := 1

	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if grid[i][j] == '1' {
				// For assigning island value
				count++
				fill(grid, i, j, nRow, nCol, count)
			}
		}
	}

	return count - 1 // Started with 1
}

func fill(grid [][]byte, x, y, X, Y, val int) {
	if x < 0 || y < 0 || x == X || y == Y {
		return
	}
	if grid[x][y] == '0' || grid[x][y] > '1' {
		return
	}
	grid[x][y] = byte(val) + '0'
	fill(grid, x-1, y, X, Y, val)
	fill(grid, x, y-1, X, Y, val)
	fill(grid, x+1, y, X, Y, val)
	fill(grid, x, y+1, X, Y, val)
}

// func fiddlingWithPopulationOfSliceOfSlices() {
// 	var m, n int // A matrix of m x n size
// 	fmt.Scan(&m, &n)

// 	type matrix [][]byte // A slice of byte slices.

// 	grid := matrix{
// 		[]byte("1"), []byte("1"), []byte("1"), []byte("1"), []byte("0"),
// 		[]byte("1"), []byte("1"), []byte("0"), []byte("1"), []byte("0"),
// 		[]byte("1"), []byte("1"), []byte("1"), []byte("1"), []byte("0"),
// 		[]byte("0"), []byte("0"), []byte("0"), []byte("0"), []byte("0")}

// 	grid = make([][]byte, m)
// 	for i := range grid {
// 		grid[i] = make([]byte, n)
// 		for j := range grid[i] {
// 			grid[i][j] = 0
// 		}
// 	}
// 	grid[0][0], grid[0][1], grid[0][2], grid[0][3] = 1, 1, 1, 1
// 	grid[1][0], grid[1][1], grid[1][3] = 1, 1, 1
// 	grid[2][0], grid[2][1], grid[2][2], grid[2][3] = 1, 1, 1, 1
// }