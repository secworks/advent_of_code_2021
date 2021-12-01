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

	my_str := string(b)
	var lines = strings.Split(my_str,"\n")
	var evalue int
	var prev int = 100000
	var cnt int = 0

	for i, elem := range lines {
		evalue, err = strconv.Atoi(elem)
		fmt.Println(i, "-> ", evalue)
		if evalue > prev {
			cnt += 1
		}
		prev = evalue
	}
	fmt.Println("Number of increasing values: ", cnt)
}
