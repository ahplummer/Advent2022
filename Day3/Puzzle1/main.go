package main

import (
	"bufio"
	"fmt"
	"os"
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
func findCommon(input string) (string, int) {
	totallen := len(input)
	firsthalf := input[0:(totallen / 2)]
	secondhalf := input[(totallen / 2):]
	lresult := ""
	vresult := 0
	for i := 0; i < len(firsthalf); i++ {
		letter := input[i : i+1]
		//fmt.Printf("Gonna look for %s in second string.\n", letter, secondhalf)
		if letterExists(secondhalf, letter) {
			//fmt.Printf("FOUND COMMON LETTER: %s\n", letter)
			lresult = letter
			vresult = letterValue(lresult)
			break
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
	totalValue := 0
	for scanner.Scan() {
		line := scanner.Text()
		totallen := len(line)
		firsthalf := line[0:(totallen / 2)]
		secondhalf := line[(totallen / 2):]
		l, v := findCommon(line)
		totalValue += v
		fmt.Printf("Line: %s, with first half: %s, second half: %s; common letter: %s, with value: %d\n", line, firsthalf, secondhalf, l, v)

	}
	fmt.Printf("Total value of common letter: %d\n", totalValue)
}
