package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Play struct {
	Choice     string
	RoundScore int
}

func (p Play) PlayScore() int {
	if p.Choice == "A" || p.Choice == "X" {
		return 1
	} else if p.Choice == "B" || p.Choice == "Y" {
		return 2
	} else if p.Choice == "C" || p.Choice == "Z" {
		return 3
	} else {
		log.Fatalln("Bad choice")
		return -1
	}
	return -1
}

type Match struct {
	PlayerA, PlayerB *Play
}

func (m Match) ComputeScores() {
	/* Roll through B's plays, and these are redefined as:
	X is lose
	Y is draw
	Z is win
	need to calculate what _would_ be required as B's move to achieve the indicated results.
	*/
	//ROCK PAPER SCISSORS
	switch m.PlayerB.Choice {
	case "X": // PlayerB is to lose
		if m.PlayerA.Choice == "A" { //choice is rock
			m.PlayerB.Choice = "Z" //scissors is loser to rock
		} else if m.PlayerA.Choice == "B" { //choice is paper
			m.PlayerB.Choice = "X" //rock is loser to paper
		} else { //choice is C which is scissors
			m.PlayerB.Choice = "Y" //paper is the loser to scissors
		}
		m.PlayerA.RoundScore = 6 //win
	case "Y": // Draw
		if m.PlayerA.Choice == "A" { //choice is rock
			m.PlayerB.Choice = "X" //rock is the draw
		} else if m.PlayerA.Choice == "B" { //choice is paper
			m.PlayerB.Choice = "Y" //paper is the draw
		} else {
			m.PlayerB.Choice = "Z" //scissors is the draw
		}
		m.PlayerA.RoundScore = 3 //draw
		m.PlayerB.RoundScore = 3 //draw
	case "Z": // PlayerB to win
		if m.PlayerA.Choice == "A" { //choice is rock
			m.PlayerB.Choice = "Y" //paper is the winner to rock
		} else if m.PlayerA.Choice == "B" { //choice is paper
			m.PlayerB.Choice = "Z" //scissors is the winner to paper
		} else {
			m.PlayerB.Choice = "X" //rock is the winner to scissors
		}
		m.PlayerB.RoundScore = 6 //win
	default:
		log.Fatalln("BAD CHOICES")
	}
	m.PlayerA.RoundScore += m.PlayerA.PlayScore()
	m.PlayerB.RoundScore += m.PlayerB.PlayScore()
}

func main() {
	fmt.Println("Welcome to day 2")
	file, err := os.Open("/home/allen/source/advent/Day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matches []Match
	for scanner.Scan() {
		line := scanner.Text()
		plays := strings.Split(line, " ")
		var playerA, playerB Play
		playerA.Choice = plays[0]
		playerB.Choice = plays[1]
		var match Match
		match.PlayerB = &playerB
		match.PlayerA = &playerA
		match.ComputeScores()
		matches = append(matches, match)
	}
	PlayerBTotal := 0
	for i, v := range matches {
		pBTotal := v.PlayerB.RoundScore

		fmt.Printf("Match %d is %s\n", i+1, v)
		fmt.Printf("----PlayerA chose %s, which is worth %d points; total is %d\n", v.PlayerA.Choice, v.PlayerA.PlayScore(), v.PlayerA.RoundScore)
		fmt.Printf("----PlayerB chose %s, which is worth %d points; total is %d\n", v.PlayerB.Choice, v.PlayerB.PlayScore(), v.PlayerB.RoundScore)
		if v.PlayerA.RoundScore > v.PlayerB.RoundScore {
			fmt.Printf("--------PlayerA wins over PlayerB, with %d vs %d points.\n", v.PlayerA.RoundScore, v.PlayerB.RoundScore)
		} else if v.PlayerA.RoundScore < v.PlayerB.RoundScore {
			fmt.Printf("--------PlayerB wins over PlayerA, with %d vs %d points. Total Points: %d \n", v.PlayerB.RoundScore, v.PlayerA.RoundScore, pBTotal)
		} else {
			fmt.Printf("--------DRAW!!! PlayerA has %d and PlayerB has %d points.\n", v.PlayerA.RoundScore, v.PlayerB.RoundScore)
		}
		PlayerBTotal += pBTotal
	}
	fmt.Printf("PlayerB's final total: %d\n", PlayerBTotal)
}
