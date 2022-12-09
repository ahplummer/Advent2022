package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WorkingDirectory() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	return mydir
}
func CountVisibleLeft(forest [][]int, i, j int) int {
	result := 0
	value := forest[i][j]
	lastTallestVisible := value
	for x := j - 1; x >= 0; x-- {
		if forest[i][x] >= lastTallestVisible {
			result++
			break
		} else if forest[i][x] < value {
			result++
			lastTallestVisible = value
		}
	}
	return result
}
func CountVisibleUp(forest [][]int, i, j int) int {
	result := 0
	value := forest[i][j]
	lastTallestVisible := value
	for x := i - 1; x >= 0; x-- {
		if forest[x][j] >= lastTallestVisible {
			result++
			break
		} else if forest[x][j] < value {
			result++
			lastTallestVisible = value
		}
	}
	return result
}
func CountVisibleRight(forest [][]int, i, j, colCount int) int {
	result := 0
	value := forest[i][j]
	lastTallestVisible := value
	for x := j + 1; x < colCount; x++ {
		if forest[i][x] >= lastTallestVisible {
			result++
			break
		} else if forest[i][x] < value {
			result++
			lastTallestVisible = value
		}
	}
	return result
}
func CountVisibleDown(forest [][]int, i, j, rowCount int) int {
	result := 0
	value := forest[i][j]
	lastTallestVisible := value
	for x := i + 1; x < rowCount; x++ {
		if forest[x][j] >= lastTallestVisible {
			result++
			break
		} else if forest[x][j] < value {
			result++
			lastTallestVisible = value
		}
	}
	return result
}
func main() {
	fmt.Println("Starting Day 8, Puzzle 2")
	file, err := os.Open(fmt.Sprintf("%s/Day8/PuzzleInput.txt", WorkingDirectory()))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var forest [][]int
	totalRows := 0
	totalColumns := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		totalRows++
		line := scanner.Text()
		totalColumns = 0
		var fRow []int
		for _, v := range line {
			vInt, err := strconv.Atoi(string(v))
			if err == nil {
				totalColumns++
				fRow = append(fRow, vInt)
			}
		}
		forest = append(forest, fRow)
	}
	//it's assumed that this is a cube (same number of rows as columns)
	outerEdges := (2 * totalRows) + (2 * totalColumns) - 4
	totalTrees := totalRows * totalColumns
	topVisibleTrees := 0
	topVisibleTreesI := -1
	topVisibleTreesJ := -1
	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalColumns; j++ {
			viewedTrees := -1
			leftVisible := CountVisibleLeft(forest, i, j)
			upVisible := CountVisibleUp(forest, i, j)
			rightVisible := CountVisibleRight(forest, i, j, totalColumns)
			bottomVisible := CountVisibleDown(forest, i, j, totalRows)

			viewedTrees = leftVisible * rightVisible * upVisible * bottomVisible
			if topVisibleTrees < viewedTrees {
				topVisibleTrees = viewedTrees
				topVisibleTreesI = i
				topVisibleTreesJ = j
			}
			fmt.Printf("Row %d, Column %d, has value %d; left: %d; up: %d; right: %d; bottom: %d; score: %d\n", i, j, forest[i][j], leftVisible, upVisible, rightVisible, bottomVisible, viewedTrees)
		}
		fmt.Println("---------")
	}
	fmt.Printf("Outer edges: %d, total trees: %d\n", outerEdges, totalTrees)
	fmt.Printf("Top tree is position %d, %d, with value %d, view score of %d", topVisibleTreesI, topVisibleTreesJ, forest[topVisibleTreesI][topVisibleTreesJ], topVisibleTrees)
}
