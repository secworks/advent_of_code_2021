//======================================================================
//
// day06_problem2.go
// -----------------
// Advent of Code 2021.
// Day 06, problem 2.
//======================================================================

package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
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


func num_fish(fishes [9]int) int {
	var num int = 0
	for i := 0 ; i < len(fishes) ; i++ {
		num += fishes[i]
	}
	return num
}


func next_day(prev_day [9]int) [9]int {
	var next_day [9]int
	for i := 0 ; i < (len(prev_day) - 1) ; i++ {
		next_day[i] = prev_day[(i + 1)]
		next_day[6] = prev_day[0] + prev_day[7]
		next_day[8] = prev_day[0]
	}
	return next_day
}


func main() {
	var days int = 256
	var fishlist [9]int

	my_input := get_input()
	for i := 0 ; i < len(my_input) ; i++ {
		num, _ := strconv.Atoi(my_input[i])
		fishlist[num] += 1
	}

	for d := 1 ; d <= days ; d++ {
		fishlist = next_day(fishlist)
	}

	fmt.Println("Solution to AoC 2021, day 06, problem 2")
	fmt.Println("---------------------------------------")
	fmt.Println("Number of lanternfish after", days, "days:", num_fish(fishlist))
	fmt.Println("")
}

//======================================================================
//======================================================================
