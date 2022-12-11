package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	I, J int
}
type Head struct {
	CurrentPosition Position
	LastDirection   string
	LastNumber      int
}

func (h *Head) StepLeft(tail *Tail) {
	h.CurrentPosition.I--
	fmt.Printf("-Head has stepped left, now at (%d,%d)\n", h.CurrentPosition.I, h.CurrentPosition.J)
}
func (h *Head) StepRight(tail *Tail) {
	h.CurrentPosition.I++
	fmt.Printf("-Head has stepped right, now at (%d,%d)\n", h.CurrentPosition.I, h.CurrentPosition.J)
}
func (h *Head) StepDown(tail *Tail) {
	h.CurrentPosition.J--
	fmt.Printf("-Head has stepped down, now at (%d,%d)\n", h.CurrentPosition.I, h.CurrentPosition.J)
}
func (h *Head) StepUp(tail *Tail) {
	h.CurrentPosition.J++
	fmt.Printf("-Head has stepped up, now at (%d,%d)\n", h.CurrentPosition.I, h.CurrentPosition.J)
}
func (h *Head) SetCurrentPosition(direction string, move int, tail *Tail) {
	h.LastDirection = direction
	h.LastNumber = move
	fmt.Printf("Head is currently at (%d,%d), setting position %s, by %d.\n", h.CurrentPosition.I, h.CurrentPosition.J, direction, move)
	if direction == "L" {
		shouldI := h.CurrentPosition.I - move
		for x := h.CurrentPosition.I; x > shouldI; x-- {
			h.StepLeft(tail)
			h.MoveTail(tail)
		}
	} else if direction == "R" {
		shouldI := h.CurrentPosition.I + move
		for x := h.CurrentPosition.I; x < shouldI; x++ {
			h.StepRight(tail)
			h.MoveTail(tail)
		}
	} else if direction == "U" {
		shouldJ := h.CurrentPosition.J + move
		for x := h.CurrentPosition.J; x < shouldJ; x++ {
			h.StepUp(tail)
			h.MoveTail(tail)
		}
	} else if direction == "D" {
		shouldJ := h.CurrentPosition.J - move
		for x := h.CurrentPosition.J; x > shouldJ; x-- {
			h.StepDown(tail)
			h.MoveTail(tail)
		}
	} else {
		log.Fatalln("Bad things happened, friend.")
	}

}
func (h *Head) MoveTail(tail *Tail) {

	fmt.Printf("--Head: current position: (%d, %d); ", h.CurrentPosition.I, h.CurrentPosition.J)
	lastPosI := tail.CurrentPosition.I
	lastPosJ := tail.CurrentPosition.J
	movedTail := false

	//let's do diagonals first.
	if h.CurrentPosition.I-lastPosI >= 1 && h.CurrentPosition.J-lastPosJ >= 2 { // && h.LastDirection == "U" { //H(4,2) T(3,0)
		//good
		fmt.Printf("DIAGONAL, need to to up and to the rightt! ")
		tail.SetCurrentPosition(lastPosI+1, lastPosJ+1)
		movedTail = true
	} else if h.CurrentPosition.I-lastPosI >= 2 && h.CurrentPosition.J-lastPosJ >= 1 { // && h.LastDirection == "R" {
		//untested
		fmt.Printf("DIAGONAL, need to to up and to the right! ")
		tail.SetCurrentPosition(lastPosI+1, lastPosJ+1)
		movedTail = true

	} else if h.CurrentPosition.I-lastPosI <= -2 && h.CurrentPosition.J-lastPosJ >= 1 { //&& h.LastDirection == "L" { //H(2,4) T(4,3)
		//good
		fmt.Printf("DIAGONAL, need to go up and to the leftt! ")
		tail.SetCurrentPosition(lastPosI-1, lastPosJ+1)
		movedTail = true

	} else if h.CurrentPosition.I-lastPosI <= -1 && h.CurrentPosition.J-lastPosJ >= 2 { //&& h.LastDirection == "U" { //
		//untested -----FIXED THIS ONE.
		fmt.Printf("DIAGONAL, need to go up and to the left! ")
		tail.SetCurrentPosition(lastPosI-1, lastPosJ+1)
		movedTail = true

	} else if h.CurrentPosition.I-lastPosI >= 2 && h.CurrentPosition.J-lastPosJ <= -1 { //&& h.LastDirection == "R" { //H(4, 3) T(2, 4)
		//good
		fmt.Printf("DIAGONAL, need to go down and to the rightt! ")
		tail.SetCurrentPosition(lastPosI+1, lastPosJ-1)
		movedTail = true

	} else if h.CurrentPosition.I-lastPosI >= 1 && h.CurrentPosition.J-lastPosJ <= -2 { //&& h.LastDirection == "D" { //
		//untested
		fmt.Printf("DIAGONAL, need to go down and to the right! ")
		tail.SetCurrentPosition(lastPosI+1, lastPosJ-1)
		movedTail = true

	} else if h.CurrentPosition.I-lastPosI <= -2 && h.CurrentPosition.J-lastPosJ <= -1 { //&& h.LastDirection == "L" { //H(2, 2) T(4, 3)
		//good
		fmt.Printf("DIAGONAL, need to go down and to the leftt! ")
		tail.SetCurrentPosition(lastPosI-1, lastPosJ-1)
		movedTail = true
	} else if h.CurrentPosition.I-lastPosI <= -1 && h.CurrentPosition.J-lastPosJ <= -2 { //&& h.LastDirection == "D" { //H(-400, 87) T(-399, 89)
		//untested
		fmt.Printf("DIAGONAL, need to go down and to the left! ")
		tail.SetCurrentPosition(lastPosI-1, lastPosJ-1)
		movedTail = true
		//SIDE TO SIDE NOW
	} else if h.CurrentPosition.I-lastPosI >= 2 {
		//good
		fmt.Printf("NEED TO MOVE RIGHT! ")
		tail.SetCurrentPosition(lastPosI+1, lastPosJ)
		movedTail = true
	} else if h.CurrentPosition.J-lastPosJ >= 2 {
		//good
		fmt.Printf("NEED TO MOVE UP! ")
		tail.SetCurrentPosition(lastPosI, lastPosJ+1)
		movedTail = true
	} else if h.CurrentPosition.I-lastPosI <= -2 {
		//good
		fmt.Printf("NEED TO MOVE LEFT! ")
		tail.SetCurrentPosition(lastPosI-1, lastPosJ)
		movedTail = true
	} else if h.CurrentPosition.J-lastPosJ <= -2 {
		//untested
		fmt.Printf("NEED TO MOVE DOWN! ")
		tail.SetCurrentPosition(lastPosI, lastPosJ-1)
		movedTail = true
	}

	if movedTail {
		fmt.Printf("Tail: before move: (%d, %d), after move: (%d,%d)\n", lastPosI, lastPosJ, tail.CurrentPosition.I, tail.CurrentPosition.J)
	} else {
		fmt.Printf("Tail: didn't move: (%d, %d)\n", tail.CurrentPosition.I, tail.CurrentPosition.J)
	}
}

type Tail struct {
	CurrentPosition  Position
	VisitedPositions []Position
}

func (t *Tail) SetCurrentPosition(i, j int) {
	t.CurrentPosition.I = i
	t.CurrentPosition.J = j
	if !t.HasVisitedPosition(i, j) {
		var lastPos Position
		lastPos.I = i
		lastPos.J = j
		t.VisitedPositions = append(t.VisitedPositions, lastPos)
	}
}
func (t *Tail) HasVisitedPosition(i, j int) bool {
	for _, v := range t.VisitedPositions {
		if v.I == i && v.J == j {
			return true
		}
	}
	return false
}

func WorkingDirectory() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	return mydir
}
func main() {
	fmt.Println("Starting Day 9, Puzzle 1")
	file, err := os.Open(fmt.Sprintf("%s/Day9/PuzzleInput.txt", WorkingDirectory()))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var tail Tail
	tail.SetCurrentPosition(0, 0)
	var head Head
	head.CurrentPosition.I = 0
	head.CurrentPosition.J = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			//fmt.Printf("Line: %s\n", line)
			move, err := strconv.Atoi(parts[1])
			if err == nil {
				head.SetCurrentPosition(parts[0], move, &tail)
			}

		}
	}
	fmt.Printf("Tail has visited %d unique positions at least once.\n", len(tail.VisitedPositions))
}
