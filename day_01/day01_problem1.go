//======================================================================
//
// day91_pronlem1.go
// -----------------
// Advent of Code 2021.
// Day 01, problem 1.
//======================================================================

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func main() {
	b, err := ioutil.ReadFile("./day01_input_problem1.txt")
	if err != nil {
		fmt.Print(err)
	}

	// Get a list of strings from the input.
	my_str := string(b)
	var lines = strings.Split(my_str,"\n")
	// Remove the extra empty element
	var trunc_lines = lines[:len(lines) - 1]

	// Process the values.
	var evalue int
	var prev int = 100000
	var cnt int = 0

	for _, elem := range trunc_lines {
		evalue, err = strconv.Atoi(elem)
		if evalue > prev {
			cnt += 1
		}
		prev = evalue
	}

	fmt.Println("Solution to AoC 2021, day 01, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println("Number of increasing values: ", cnt)
	fmt.Println("")
}

//======================================================================
//======================================================================
