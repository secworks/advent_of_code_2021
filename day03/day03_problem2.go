//======================================================================
// day03_problem2.go
//======================================================================

package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	//	"reflect"
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


func get_match_list(pos int, prefer string, old_list []string) []string {
	var new_list []string

	for _, num := range(old_list) {
		if string(num[pos]) == prefer {
			new_list = append(new_list, num)
		}
	}
	return new_list
}


func get_most_common_bit(pos int, prefer string, num_list []string) string {
	var ones, zeros int
	var ch string

	for _, num := range (num_list) {
		ch = string(num[pos])
		if ch == "1" {
			ones += 1
		} else {
			zeros += 1
		}
	}

	if ones == zeros {
		return prefer
	} else if ones > zeros {
		return "1"
	} else {
		return "0"
	}
}


func get_least_common_bit(pos int, prefer string, num_list []string) string {
	var ones, zeros int
	var ch string

	for _, num := range (num_list) {
		ch = string(num[pos])
		if ch == "1" {
			ones += 1
		} else {
			zeros += 1
		}
	}

	if ones == zeros {
		return prefer
	} else if ones < zeros {
		return "1"
	} else {
		return "0"
	}
}


func main() {
	my_input := get_input()

	var oxygen_list []string
	oxygen_list = append(oxygen_list, my_input...)

	var bit_index int
	var most_common string
	for len(oxygen_list) > 1 {
		most_common = get_most_common_bit(bit_index, "1", oxygen_list)
		oxygen_list = get_match_list(bit_index, most_common, oxygen_list)
		bit_index += 1
	}

	var co2_list []string
	co2_list = append(co2_list, my_input...)
	var least_common string
	bit_index = 0
	for len(co2_list) > 1 {
		least_common = get_least_common_bit(bit_index, "0", co2_list)
		co2_list = get_match_list(bit_index, least_common, co2_list)
		bit_index += 1
	}

	var oxygen_num, co2_num int64
	oxygen_num, _ = strconv.ParseInt(oxygen_list[0], 2, 32)
	co2_num, _ = strconv.ParseInt(co2_list[0], 2, 32)

	fmt.Println("Solution to AoC 2021, day 03, problem 2")
	fmt.Println("---------------------------------------")
	fmt.Println("oxygen_number", oxygen_list, "=", oxygen_num)
	fmt.Println("co2_number   ", co2_list, "=", co2_num)
	fmt.Println("Life support rating", oxygen_num * co2_num)
}

//======================================================================
//======================================================================
