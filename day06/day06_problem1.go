//======================================================================
//
// day06_problem1.go
// -----------------
// Advent of Code 2021.
// Day 06, problem 1.
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


func next_day(prev_fish []int) []int {
	var fish int
	var old_fish []int
	var new_fish []int

	for i := 0 ; i < len(prev_fish) ; i++ {
		fish = prev_fish [i]
		if fish == 0 {
			old_fish = append(old_fish, 6)
			new_fish = append(new_fish, 8)
		} else {
			fish -= 1
			old_fish = append(old_fish, fish)
		}
	}
	return append(old_fish, new_fish...)
}


func main() {
	var days int = 80
	var fishlist []int

	my_input := get_input()
	for i := 0 ; i < len(my_input) ; i++ {
		num, _ := strconv.Atoi(my_input[i])
		fishlist = append(fishlist, num)
	}

	for d := 1 ; d <= days ; d++ {
		fishlist = next_day(fishlist)
	}

	fmt.Println("Solution to AoC 2021, day 06, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println("Number of lanternfish after", days, "days:", len(fishlist))
	fmt.Println("")
}

//======================================================================
//======================================================================
