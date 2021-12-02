//======================================================================
//
// day02_problem2.go
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


func get_moves(fname string) ([]move, error) {
	var moves []move
	var m move
	var cmd string
	var value int

	file, e := os.Open(fname)
	if e != nil {
		fmt.Println("Error is = ",e)
		return moves, e
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		cmd = s[0]
		value, _ = strconv.Atoi(s[1])
		m.command, m.value = cmd, value
		moves = append(moves, m)

	}

	defer file.Close()
	return moves, nil
}


func main() {
	my_moves, _ := get_moves("./day02_input.txt")

	var distance int
	var depth int
	var aim int

	for _, move := range my_moves {
		if move.command == "forward" {
			depth += (aim * move.value)
			distance += move.value
		}

		if move.command == "up" {
			aim -= move.value
		}

		if move.command == "down" {
			aim += move.value
		}

	}

	fmt.Println("Solution to AoC 2021, day 02, problem 2")
	fmt.Println("---------------------------------------")
	fmt.Println("Final distance:", distance)
	fmt.Println("Final depth:   ", depth)
	fmt.Println("Solution:      ", distance * depth)
	fmt.Println("")
}

//======================================================================
//======================================================================