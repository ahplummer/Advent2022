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
	fmt.Println("Starting Day 1, Puzzle 1")
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
	elfIndex1 := -1
	for i, v := range elves {
		fmt.Printf("Elf %d has %d items, totalling %d calories.\n", i+1, len(v.items), v.TotalAmount())
		if v.TotalAmount() >= highestCalories1 {
			highestCalories1 = v.TotalAmount()
			elfIndex1 = i
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("Elf %d has highest calories: %d\n", elfIndex1+1, highestCalories1)
}
