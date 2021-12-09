//======================================================================
// day05_problem1.go
//======================================================================

package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	// "reflect"
	"strconv"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}


func get_input() []string {
	var rows []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\n")
		rows = append(rows, s[0])
	}
	return rows
}


func get_lines(input []string) []line {
	var my_lines []line
	var l line

	//	fmt.Println("type of input:", reflect.TypeOf(input))
	//	fmt.Println("len(input):", len(input))

	for _, line := range(input) {
		l.x1, _ = strconv.Atoi(string(line[0]))
		l.y1, _ = strconv.Atoi(string(line[2]))
		l.x2, _ = strconv.Atoi(string(line[7]))
		l.y2, _ = strconv.Atoi(string(line[9]))
		my_lines = append(my_lines, l)
	}
	return my_lines
}

func main() {
	my_input := get_input()
	my_lines := get_lines(my_input)

	fmt.Println("Solution to AoC 2021, day 05, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println("my_lines:", my_lines)
	fmt.Println("")
}

//======================================================================
//======================================================================
