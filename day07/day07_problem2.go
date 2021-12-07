//======================================================================
//
// day07_problem2.go
// -----------------
// Advent of Code 2021.
// Day 07, problem 2.
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


func get_fuel(start int, stop int) int {
	var fuel int

	if start == stop {
		fuel = 0
	}

	if start > stop {
		fuel = int((start - stop) * (((start - stop) + 1) / 2))
	}

	if stop > start {
		fuel = int((stop - start) * (((stop - start) + 1) / 2))
	}
	//	fmt.Println("start:", start, "stop:", stop, "fuel:", fuel)
	return fuel
}


func main() {
	var crablist []int
	var fuel, minfuel, minpos int
	var i int

	my_input := get_input()
	for i := 0 ; i < len(my_input) ; i++ {
		num, _ := strconv.Atoi(my_input[i])
		crablist = append(crablist, num)
	}
	sort.Ints(crablist)
	fmt.Println(crablist)

	minpos = 0
	minfuel = 10000000000000000
	for i = crablist[0] ; i <= crablist[len(crablist) - 1] ; i++ {
		fuel = 0
		for _,crabpos := range crablist {
			fuel += get_fuel(crabpos, i)
		}
		if fuel < minfuel {
			minfuel = fuel
			minpos = i
		}
	}


	fmt.Println("Solution to AoC 2021, day 07, problem 2")
	fmt.Println("---------------------------------------")
	fmt.Println("Minimal fuel", minfuel, "at position", minpos)
}

//======================================================================
//======================================================================
