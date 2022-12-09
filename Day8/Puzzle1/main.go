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
func IsVisibleLeft(forest [][]int, i, j int) bool {
	result := false
	if j == 0 {
		result = true
	} else {
		value := forest[i][j]
		result = true
		//walk left
		for x := 0; x < j; x++ {
			if forest[i][x] >= value {
				result = false
				break
			}
		}
	}
	return result
}
func IsVisibleRight(forest [][]int, i, j, lenRow int) bool {
	result := false
	if j == lenRow-1 {
		result = true
	} else {
		value := forest[i][j]
		result = true
		//walk right
		for x := lenRow - 1; x > j; x-- {
			if forest[i][x] >= value {
				result = false
				break
			}
		}
	}
	return result
}
func IsVisibleTop(forest [][]int, i, j int) bool {
	result := false
	if i == 0 {
		result = true
	} else {
		value := forest[i][j]
		result = true
		//walk up
		for x := 0; x < i; x++ {
			if forest[x][j] >= value {
				result = false
				break
			}
		}
	}
	return result
}
func IsVisibleBottom(forest [][]int, i, j, lenColumn int) bool {
	result := false
	if i == lenColumn-1 {
		result = true
	} else {
		value := forest[i][j]
		result = true
		//walk bottom
		for x := lenColumn - 1; x > i; x-- {
			if forest[x][j] >= value {
				result = false
				break
			}
		}
	}
	return result
}
func main() {
	fmt.Println("Starting Day 8, Puzzle 1")
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
		//fmt.Printf("Line: %s\n", line)
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
	visibleTrees := 0
	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalColumns; j++ {
			//only doing inner part, skipping edges
			isVisibleLeft := IsVisibleLeft(forest, i, j)
			isVisibleRight := IsVisibleRight(forest, i, j, totalRows)
			isVisibleTop := IsVisibleTop(forest, i, j)
			isVisibleBottom := IsVisibleBottom(forest, i, j, totalColumns)
			if isVisibleLeft || isVisibleRight || isVisibleTop || isVisibleBottom {
				visibleTrees++
			}
			fmt.Printf("Row %d, Column %d, has value %d; left: %t, right: %t, top: %t, bottom: %t\n", i, j, forest[i][j], isVisibleLeft, isVisibleRight, isVisibleTop, isVisibleBottom)
		}
		fmt.Println("---------")
	}
	fmt.Printf("Outer edges: %d, total trees: %d, total visible trees: %d\n", outerEdges, totalTrees, visibleTrees)
}
