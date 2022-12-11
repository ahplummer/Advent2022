package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WorkingDirectory() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	return mydir
}
func ProcessCycleGetSignalStrength(cycle int, sumcycle *int, xValue int) int {
	result := 0
	//sumcyclevalue := *sumcycle
	if cycle == 20 {
		//first cycle
		fmt.Println("---------20th cycle---------")
		*sumcycle = 60
		result = xValue * cycle
	} else if cycle == *sumcycle {
		*sumcycle += 40
		result = xValue * cycle
		fmt.Println("---------every 40th cycle---------")
	} else {
		//fmt.Printf("Cycle %d, xValue %d\n", cycle, xValue)
	}
	return result
}
func main() {
	fmt.Println("Starting Day 10, Puzzle 1")
	file, err := os.Open(fmt.Sprintf("%s/Day10/PuzzleInput.txt", WorkingDirectory()))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	xValue := 1
	cycle := 0
	sumcycle := 20
	sumSignalStrength := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("%s\n", line)
		if line[:4] == "noop" {
			cycle++
			signalStrength := ProcessCycleGetSignalStrength(cycle, &sumcycle, xValue)
			if signalStrength > 0 {
				sumSignalStrength += signalStrength
			}
			fmt.Printf("%s: Cycle %d, Value of X is %d, signalStrenth is %d\n", line, cycle, xValue, signalStrength)
		} else if line[:4] == "addx" {
			parts := strings.Split(line, " ")
			for i := 1; i <= 2; i++ {
				cycle++
				//fmt.Printf("execution %d; ", i)
				signalStrength := ProcessCycleGetSignalStrength(cycle, &sumcycle, xValue)
				if signalStrength > 0 {
					sumSignalStrength += signalStrength
				}
				fmt.Printf("%d: %s: Cycle %d, Value of X is %d, signalStrenth is %d\n", i, line, cycle, xValue, signalStrength)

				if i == 2 {
					v, _ := strconv.Atoi(parts[1])
					//fmt.Printf("now adding %d to X.", v)
					xValue += v
				}

			}
			//fmt.Printf("\n")
		}
		fmt.Printf("Sum of signal strength: %d\n", sumSignalStrength)
	}
}
