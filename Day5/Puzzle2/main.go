package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	IdentifierIndex int
	Identifier      int
	//lower number is lower in the stack. Like a classical CS stack, you can only pop off the top.
	Items []string
}

func (s *Stack) InsertAtBottom(letter string) {
	if len(strings.Trim(letter, " ")) == 0 {
		return
	}
	//need to shift things.
	var olditems []string
	for _, v := range s.Items {
		olditems = append(olditems, v)
	}
	s.Items = nil
	s.Items = append(s.Items, letter)
	for _, v := range olditems {
		s.Items = append(s.Items, v)
	}
}
func (s *Stack) PopAndShift(destStack *Stack) {
	item := s.Items[len(s.Items)-1]
	//copy it over
	destStack.Items = append(destStack.Items, item)
	//remove it
	s.Items = s.Items[:len(s.Items)-1]

}
func (s *Stack) ShiftMoves(destStack *Stack, moves int) {
	stackheight := len(s.Items)

	/*
		A - 3
		B - 2
		C - 1
		D - 0
		To move top 3, moves would be 3, with stackheight of 4
	*/
	if moves <= stackheight {
		//can move things
		for i := stackheight - moves; i < stackheight; i++ {
			item := s.Items[i]
			//copy it over
			destStack.Items = append(destStack.Items, item)
			s.Items[i] = "-"
		}
		//remove the "-'s"
		j := 0
		for _, v := range s.Items {
			if v != "-" {
				s.Items[j] = v
				j++
			}
		}
		s.Items = s.Items[:j]
	} else {
		log.Fatalln("UNEXPECTED STACK HEIGHT")
	}
}
func (s *Stack) ProcessCurrentStack(config string) {

	if len(config) > s.IdentifierIndex {
		item := config[s.IdentifierIndex : s.IdentifierIndex+1]
		//fmt.Printf("Processing %s, item is %s\n", config, item)
		s.InsertAtBottom(item)
	}
}
func PrintStacks(stacks []*Stack) {
	for _, v := range stacks {
		fmt.Printf("Stack: %s\n", v)
	}
}
func main() {
	fmt.Println("Day 5, Puzzle 1")
	file, err := os.Open("/home/allen/source/advent/Day5/PuzzleInput.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	var currentStacks []string
	var moveLines []string
	var stacks []*Stack
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		//fmt.Printf("Line: %s\n", line)
		if lineNumber <= 8 {
			//fmt.Println("stack config")
			currentStacks = append(currentStacks, line)
		} else if lineNumber == 9 {
			//fmt.Println("stack defin")
			parts := strings.Split(line, " ")
			if len(parts) > 0 {
				for _, v := range parts {
					si, err := strconv.Atoi(v)
					if err == nil {
						var stack Stack
						stack.Identifier = si
						stack.IdentifierIndex = strings.Index(line, v)
						stacks = append(stacks, &stack)
					}
				}
			}
		} else if len(line) > 4 {
			if line[0:4] == "move" {
				moveLines = append(moveLines, line)
			}
		}
	}
	for _, v := range stacks {
		//fmt.Printf("Stack: %d\n", v.Identifier)
		for _, w := range currentStacks {
			v.ProcessCurrentStack(w)
		}
	}
	PrintStacks(stacks)
	/* NOW, let's walk through the movements.*/
	for _, v := range moveLines {
		//move 3 from 8 to 9 -> this means move 3 things from column 8 to column 9, starting from top.
		parts := strings.Split(v, " ")
		if len(parts) == 6 {
			moves := 0
			source := 0
			dest := 0
			moves, _ = strconv.Atoi(parts[1])
			source, _ = strconv.Atoi(parts[3])
			dest, _ = strconv.Atoi(parts[5])
			//shifting "moves" number of boxes in the same order.
			stacks[source-1].ShiftMoves(stacks[dest-1], moves)
		}
	}
	fmt.Println("===============CHANGED:")
	PrintStacks(stacks)
}
