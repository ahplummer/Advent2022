package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func letterExists(input string, toFind string) bool {
	result := false
	for i := 0; i < len(input); i++ {
		letter := input[i : i+1]
		if letter == toFind {
			result = true
			break
		}
	}
	return result

}
func letterValue(letter string) int {
	avail := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := 0
	for i := 0; i < len(avail); i++ {
		test := avail[i : i+1]
		if test == letter {
			result = i + 1
			break
		}
	}
	return result
}
func findCommon(input1, input2, input3 string) (string, int) {
	lresult := ""
	vresult := 0
	for i := 0; i < len(input1); i++ {
		letter := input1[i : i+1]
		fmt.Printf("Gonna look for %s in second string.\n", letter, input2)
		if letterExists(input2, letter) {
			fmt.Printf("--found common letter: %s\n", letter)
			fmt.Printf("Gonna look for %s in third string.\n", letter, input3)
			if letterExists(input3, letter) {
				fmt.Printf("---found common letter: %s\n", letter)
				lresult = letter
				vresult = letterValue(lresult)
				break
			}
		}
	}
	return lresult, vresult
}
func main() {
	fmt.Println("Puzzle 1 for Day 3")
	file, err := os.Open("/home/allen/source/advent/Day3/PuzzleInput.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	totvalue := 0
	for i := 0; i < len(lines); i = i + 3 {
		line1 := strings.Join(lines[i:i+1], "")
		line2 := strings.Join(lines[i+1:i+2], "")
		line3 := strings.Join(lines[i+2:i+3], "")
		l, v := findCommon(fmt.Sprintf("%s", line1), fmt.Sprintf("%s", line2), fmt.Sprintf("%s", line3))
		fmt.Printf("Line1: %s, Line2: %s, Line3:%s, common letter is: %s, value: %d\n", line1, line2, line3, l, v)
		totvalue += v
	}
	fmt.Printf("Total Value: %d\n", totvalue)
}
