package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	items []int
}

func (e Elf) TotalAmount() int {
	total := 0
	for _, v := range e.items {
		total += v
	}
	return total
}
func main() {
	fmt.Println("Starting Day2, Puzzle2")
	file, err := os.Open("/home/allen/source/advent/Day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var elves []Elf
	scanner := bufio.NewScanner(file)
	var items []int
	for scanner.Scan() {
		calories, err := strconv.Atoi(scanner.Text())
		if err == nil {
			fmt.Printf("Calories: %d\n", calories)
			items = append(items, calories)
		} else {
			var elf Elf
			elf.items = items
			elves = append(elves, elf)
			fmt.Printf("Adding elf to list of elves: %s\n", elf)
			items = make([]int, 0)
		}
	}
	//get the last one
	var elf Elf
	elf.items = items
	elves = append(elves, elf)
	fmt.Printf("Adding last elf to list of elves: %s\n", elf)

	highestCalories1 := 0
	highestCalories2 := 0
	highestCalories3 := 0
	elfIndex1 := -1
	elfIndex2 := -1
	elfIndex3 := -1
	for i, v := range elves {
		fmt.Printf("Elf %d has %d items, totalling %d calories.\n", i+1, len(v.items), v.TotalAmount())
		if v.TotalAmount() >= highestCalories1 {
			fmt.Printf("REPLACING FIRST PLACE, shuffling first to second and second to third....was %d, now %d\n", highestCalories1, v.TotalAmount())
			highestCalories3 = highestCalories2
			highestCalories2 = highestCalories1
			highestCalories1 = v.TotalAmount()
			elfIndex1 = i
		} else if v.TotalAmount() >= highestCalories2 {
			fmt.Printf("REPLACING SECOND PLACE....shuffling second to thirdwas %d, now %d\n", highestCalories2, v.TotalAmount())
			highestCalories3 = highestCalories2
			highestCalories2 = v.TotalAmount()
			elfIndex2 = i
		} else if v.TotalAmount() >= highestCalories3 {
			fmt.Printf("REPLACING THIRD PLACE....was %d, now %d\n", highestCalories3, v.TotalAmount())
			highestCalories3 = v.TotalAmount()
			elfIndex3 = i
		}
	}
	fmt.Printf("Elf %d has highest calories: %d\n", elfIndex1+1, highestCalories1)
	fmt.Printf("Elf %d has second highest calories: %d\n", elfIndex2+1, highestCalories2)
	fmt.Printf("Elf %d has third highest calories: %d\n", elfIndex3+1, highestCalories3)
	fmt.Printf("Total of all three: %d\n", highestCalories1+highestCalories2+highestCalories3)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
