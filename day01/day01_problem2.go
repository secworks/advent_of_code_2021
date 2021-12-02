//======================================================================
//
// day01_problem2.go
// -----------------
// Advent of Code 2021.
// Day 01, problem 2.
//======================================================================

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func slicesum(s []int) int {
	var sum = 0
	for _, e := range s {
		sum += e
	}
	return sum
}


func main() {
	b, err := ioutil.ReadFile("./day01_input_problem1.txt")
	if err != nil {
		fmt.Print(err)
	}

	// Get a list of strings from the input.
	// Remove the extra empty element.
	my_str := string(b)
	var lines = strings.Split(my_str,"\n")
	var trunc_lines = lines[:len(lines) - 1]

	// Convert to list of ints.
	var values []int
	var value int
	for _, e := range trunc_lines {
		value, err = strconv.Atoi(e)
		values = append(values, value)
	}

	var cnt int = 0
	var slice1 []int
	var slice2 []int
	var sum1 int
	var sum2 int
	for i := 0 ; i < len(values) - 3 ; i++ {
		slice1 = values[i : i + 3]
		sum1 = slice1[0] + slice1[1] + slice1[2]
		slice2 = values[i + 1 : i + 4]
		sum2 = slice2[0] + slice2[1] + slice2[2]
		fmt.Println("i:", i, "slice1:", slice1, "sum1:", sum1, "slice2:", slice2, "sum2:", sum2)
		if sum1 < sum2 {
			cnt += 1
		}
	}

	fmt.Println("Solution to AoC 2021, day 01, problem 2")
	fmt.Println("---------------------------------------")
	fmt.Println("Number of increasing values: ", cnt)
	fmt.Println("")
}

//======================================================================
//======================================================================
