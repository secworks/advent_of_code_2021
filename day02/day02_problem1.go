//======================================================================
//
// day02_problem1.go
// -----------------
// Advent of Code 2021.
// Day 02, problem 1.
//======================================================================

package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
)


type move struct {
	command string
	value int
}


func get_moves() []move {
	var moves []move
	var m move
	var cmd string
	var value int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		cmd = s[0]
		value, _ = strconv.Atoi(s[1])
		m.command, m.value = cmd, value
		moves = append(moves, m)
	}

	return moves
}


func main() {
	my_moves := get_moves()

	var distance int
	var depth int

	for _, move := range my_moves {
		if move.command == "forward" {
			distance += move.value
		}

		if move.command == "up" {
			depth -= move.value
		}

		if move.command == "down" {
			depth += move.value
		}
	}

	fmt.Println("Solution to AoC 2021, day 02, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println("Final distance:", distance)
	fmt.Println("Final depth:   ", depth)
	fmt.Println("Solution:      ", distance * depth)
	fmt.Println("")
}

//======================================================================
//======================================================================
