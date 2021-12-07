//======================================================================
//
// day07_problem1.go
// -----------------
// Advent of Code 2021.
// Day 07, problem 1.
//======================================================================

package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)


func get_input() []string {
	var binstr []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\n")
		binstr = append(binstr, s[0])
	}
	return binstr
}


func main() {
	my_input := get_input()

	fmt.Println("Solution to AoC 2021, day 07, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println(my_input)
	fmt.Println("")
}

//======================================================================
//======================================================================
