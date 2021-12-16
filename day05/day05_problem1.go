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
		points := strings.Split(line, " -> ")
		p1 :=strings.Split(points[0], ",")
		p2 :=strings.Split(points[1], ",")

		// fmt.Println("line:", i, "points:", points, "p1:", p1, "p2:", p2)

		l.x1, _ = strconv.Atoi(string(p1[0]))
		l.y1, _ = strconv.Atoi(string(p1[1]))
		l.x2, _ = strconv.Atoi(string(p2[0]))
		l.y2, _ = strconv.Atoi(string(p2[1]))

		// fmt.Println("line:", i, "x1:", l.x1, "y1:", l.y1, "x2:", l.x2, "y2:", l.y2)

		my_lines = append(my_lines, l)
	}
	return my_lines
}


func get_dimensions(lines []line) (int, int, int, int) {
	var xmin = 1000000
	var xmax = 0
	var ymin = 1000000
	var ymax = 0
	for _, line := range lines {
		if line.x1 < xmin {
			xmin = line.x1
		}

		if line.x1 > xmax {
			xmax = line.x1
		}

		if line.x2 < xmin {
			xmin = line.x2
		}

		if line.x2 > xmax {
			xmax = line.x2
		}

		if line.y1 < ymin {
			ymin = line.y1
		}

		if line.y1 > ymax {
			ymax = line.y1
		}

		if line.y2 < ymin {
			ymin = line.y2
		}

		if line.y2 > ymax {
			ymax = line.y2
		}
	}
	return xmin, xmax, ymin, ymax
}


func print_area(area [][]int) {
	fmt.Println("Printing the area:")
	fmt.Println("------------------")

	for _, row := range(area) {
		fmt.Println(row)
	}
	fmt.Println("")
}


func create_area(xmax, ymax int) [][]int {
	area := make([][]int, (ymax + 1))
	for i := range area {
		area[i] = make([]int, (xmax + 1))
	}
	return area
}


func draw_rectline(line line, area [][]int) {
	var i int
	var start, stop int

	if line.x1 == line.x2 {
		if line.y1 > line.y2 {
			start = line.y2
			stop = line.y1
		} else {
			start = line.y1
			stop = line.y2
		}
		for i = start ; i <= stop ; i++ {
			area[i][line.x1] += 1
		}
	} else {
		if line.x1 > line.x2 {
			start = line.x2
			stop = line.x1
		} else {
			start = line.x1
			stop = line.x2
		}
		for i = start ; i <= stop ; i++ {
			area[line.y1][i] += 1
		}
	}
}


func draw_rectlines(lines []line, area [][]int) {
	for _, line := range lines {
		if (line.x1 == line.x2) || (line.y1 == line.y2) {
			draw_rectline(line, area)
		}
	}
}


func get_multiple_crossings(area [][]int, xdim int, ydim int) int {
	var num_points int
	var x, y int

	for y = 0 ; y <= ydim ; y++ {
		for x = 0 ; x <= xdim ; x++ {
			if area[y][x] > 1 {
				num_points += 1
			}
		}
	}

	return num_points
}


func main() {
	my_input := get_input()
	my_lines := get_lines(my_input)
	xmin, xmax, ymin, ymax := get_dimensions(my_lines)

	my_area := create_area(xmax, ymax)
	draw_rectlines(my_lines, my_area)

	num_points := get_multiple_crossings(my_area, xmax, ymax)

	fmt.Println("Solution to AoC 2021, day 05, problem 1")
	fmt.Println("---------------------------------------")
	fmt.Println("xmin", xmin, "xmax", xmax)
	fmt.Println("ymin", ymin, "ymax", ymax)
	fmt.Println("Number of points with multipe overlap:", num_points)
	fmt.Println("")
}


//======================================================================
//======================================================================
