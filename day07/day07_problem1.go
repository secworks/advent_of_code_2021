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
	"strconv"
	"sort"
)


func get_input() []string {
	var binstr []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		binstr = append(binstr, s...)
	}
	return binstr
}


func main() {
	var crablist []int
	var fuel, minfuel, minpos int

	my_input := get_input()
	for i := 0 ; i < len(my_input) ; i++ {
		num, _ := strconv.Atoi(my_input[i])
		crablist = append(crablist, num)
	}
	sort.Ints(crablist)

	minpos = 0
	minfuel = 100000000000
	for _,pos := range crablist {
		fuel = 0
		for _,crabpos := range crablist {
			if crabpos > pos {
				fuel += (crabpos - pos)
			}
			if crabpos < pos {
				fuel += (pos - crabpos)
			}
		}
		if fuel < minfuel {
			minfuel = fuel
			minpos = pos
		}
	}


	fmt.Println("Solution to AoC 2021, day 07, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println("Minimal fuel", minfuel, "at position", minpos)
}

//======================================================================
//======================================================================
