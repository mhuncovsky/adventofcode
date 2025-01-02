package day13

import (
	"aoc2024/common"
	"errors"
	"fmt"
)

const SHOW = false

func AB(x0, x1, y0, y1, px, py int) (int, int, error) {
	b := (x0*py - y0*px) / (x0*y1 - x1*y0)
	a := (px - b*x1) / x0
	if a*x0+b*x1 == px && a*y0+b*y1 == py {
		return a, b, nil
	}
	return 0, 0, errors.New("no solution")
}

func solve1(lines []string) int {
	return 0
}

func solve2(lines []string, n int) int {
	return 0
}

func Main() {
	fmt.Printf("Day 13\n=====\n")

	lines_test, err := common.LoadFileLines("day13/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect ???: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect ???: %d\n", solve2(lines_test, 25))
	}

	lines, err := common.LoadFileLines("day13/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines, 75))
	}
}
