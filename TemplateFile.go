package main

import (
	"bufio"
	"fmt"
	"os"
)

func WorkingDirectory() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	return mydir
}
func main() {
	fmt.Println("Starting Day X, Puzzle Y")
	file, err := os.Open(fmt.Sprintf("%s/DayX/PuzzleInput.txt", WorkingDirectory()))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line: %s\n", line)
	}
}
