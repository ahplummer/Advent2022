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
func StartMarker(input string) int {
	result := 0

	var buffSlice [14]rune
	for i, v := range input {
		if SlideIn(&buffSlice, v) {
			result = i + 1
			break
		}
	}
	return result
}
func Duplicates(buffSlice *[14]rune) bool {
	result := false
	for i := 0; i < 14; i++ {
		for j := 0; j < 14; j++ {
			if i != j {
				if buffSlice[i] != 0 && buffSlice[i] == buffSlice[j] {
					result = true
					break
				} else if buffSlice[i] == 0 {
					result = true
					break
				}
			}
		}
	}
	return result
}
func SlideIn(buffSlice *[14]rune, r rune) bool {
	//buffer should only have 14 items at any given time.
	//If any repeat, need to shove off the side.
	result := false
	loading := false
	for i := 0; i < 14; i++ {
		if buffSlice[i] == 0 {
			buffSlice[i] = r
			loading = true
			break
		}
	}
	if !loading {
		//need to shift things
		var newSlice [13]rune
		for i := 1; i < len(buffSlice); i++ {
			newSlice[i-1] = buffSlice[i]
		}
		for i := 0; i < len(newSlice); i++ {
			buffSlice[i] = newSlice[i]
		}
		buffSlice[13] = r
	}
	hasDupes := Duplicates(buffSlice)
	if !hasDupes {
		result = true
	}
	return result
}
func main() {
	fmt.Println("Starting Day X, Puzzle Y")
	file, err := os.Open(fmt.Sprintf("%s/Day6/PuzzleInput.txt", WorkingDirectory()))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lastline := ""
	for scanner.Scan() {
		lastline = scanner.Text()
		break
	}
	/*
	   mjqjpqmgbljsphdztnvjfqwrcgsmlb: first marker after character 19
	   bvwbjplbgvbhsrlpgdmjqwftvncz: first marker after character 23
	   nppdvjthqldpwncqszvftbrmjlhg: first marker after character 23
	   nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: first marker after character 29
	   zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: first marker after character 26
	*/
	fmt.Printf("Testing input %s: %d\n", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", StartMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	fmt.Printf("Testing input %s: %d\n", "bvwbjplbgvbhsrlpgdmjqwftvncz", StartMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	fmt.Printf("Testing input %s: %d\n", "nppdvjthqldpwncqszvftbrmjlhg", StartMarker("nppdvjthqldpwncqszvftbrmjlhg"))
	fmt.Printf("Testing input %s: %d\n", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", StartMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	fmt.Printf("Testing input %s: %d\n", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", StartMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))

	//if tested well, do actual:
	fmt.Printf("Real Input %s: %d\n", lastline, StartMarker(lastline))

}
