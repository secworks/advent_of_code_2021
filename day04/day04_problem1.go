//======================================================================
//
// day04_problem1.go
// -----------------
// Advent of Code 2021.
// Day 04, problem 1.
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
		s := strings.Split(scanner.Text(), "\n")
		binstr = append(binstr, s[0])
	}
	return binstr
}


func main() {
	var ones int;
	var zeros int;
	var num string
	var ch string
	var gamma []string
	var epsilon []string
	var gamma_num int64
	var epsilon_num int64

	my_input := get_input()

	for i := 0 ; i < len(my_input[0]) ; i++ {
		ones = 0
		zeros = 0
		for j := 0 ; j < len(my_input) ; j++ {
			num = my_input[j]
			ch = string(num[i])
			if ch == "1" {
				ones += 1
			} else {
				zeros += 1
			}
		}
		if ones > zeros {
			gamma = append(gamma, string("1"))
			epsilon = append(epsilon, string("0"))
		} else {
			gamma = append(gamma, string("0"))
			epsilon = append(epsilon, string("1"))
		}
	}

	gamma_num, _ = strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	epsilon_num, _ = strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)

	fmt.Println("Solution to AoC 2021, day 04, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println("gamma:  ", gamma,   "=", gamma_num)
	fmt.Println("epslion:", epsilon, "=", epsilon_num)
	fmt.Println("Answer: ", gamma_num * epsilon_num)
	fmt.Println("")
}

//======================================================================
//======================================================================
