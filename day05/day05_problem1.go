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
	// "strconv"
)


func get_input() []string {
	var rows []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\n")
		rows = append(rows, s[0])
	}
	return rows
}


func main() {
	my_input := get_input()

	fmt.Println("Solution to AoC 2021, day 05, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println(my_input)
	fmt.Println("")
}

//======================================================================
//======================================================================
