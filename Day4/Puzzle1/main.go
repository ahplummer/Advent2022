package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elf struct {
	SectionDefinition string
	Sections          []int
}

func (e *Elf) GenerateSections() {
	parsed := strings.Split(e.SectionDefinition, "-")
	if len(parsed) == 2 {
		i, _ := strconv.Atoi(parsed[0])
		j, _ := strconv.Atoi(parsed[1])
		for idx := i; idx <= j; idx++ {
			e.Sections = append(e.Sections, idx)
		}
	}
}

type ElfPair struct {
	Elf1, Elf2 Elf
}

func (p *ElfPair) DoesFullyContain() bool {
	e1Start := p.Elf1.Sections[0]
	e1Stop := p.Elf1.Sections[len(p.Elf1.Sections)-1]
	e2Start := p.Elf2.Sections[0]
	e2Stop := p.Elf2.Sections[len(p.Elf2.Sections)-1]
	if e2Start >= e1Start && e2Stop <= e1Stop {
		return true
	} else if e1Start >= e2Start && e1Stop <= e2Stop {
		return true
	}
	return false
}

func main() {
	fmt.Println("Day 4, Puzzle 1")
	file, err := os.Open("/home/allen/source/advent/Day4/PuzzleInput.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var elfPairs []ElfPair
	for scanner.Scan() {
		line := scanner.Text()
		elves := strings.Split(line, ",")
		if len(elves) == 2 {
			var elf1 Elf
			elf1.SectionDefinition = elves[0]
			elf1.GenerateSections()
			var elf2 Elf
			elf2.SectionDefinition = elves[1]
			elf2.GenerateSections()
			var elfPair ElfPair
			elfPair.Elf1 = elf1
			elfPair.Elf2 = elf2
			elfPairs = append(elfPairs, elfPair)
		}
	}
	full := 0
	for _, v := range elfPairs {
		//fmt.Printf("Elf Pair: %s\n", v)
		if v.DoesFullyContain() {
			fmt.Printf("Elf 1 and Elf 2 have some full containment: \n %s\n %s \n", v.Elf1, v.Elf2)
			full++
		}
	}
	fmt.Printf("%d pairs are fully contained.\n", full)
}
