package day2

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
)

func lineToNumbers(line string) []int {
	ns := make([]int, 0)
	for _, sn := range strings.Split(line, " ") {
		n64, err := strconv.ParseInt(sn, 10, 64)
		if err != nil {
			panic(err)
		}
		ns = append(ns, int(n64))
	}
	return ns
}

func linesToNumbers(lines []string) [][]int {
	numbers := make([][]int, 0)
	for _, line := range lines {
		numbers = append(numbers, lineToNumbers(line))
	}
	return numbers
}

func isSafe(numbers []int) bool {
	asc := false
	safe := true
	prev := 0

	for i, n := range numbers {

		if i == 0 {
			prev = n
			continue
		}

		if n == prev {
			safe = false
			break
		}

		if n > prev {
			if (n - prev) > 3 {
				safe = false
				break
			}
		} else if (prev - n) > 3 {
			safe = false
			break
		}

		if i == 1 {
			if n > prev {
				asc = true
			}
			prev = n
			continue
		}

		if i > 1 {
			if (n < prev && asc) || (n > prev && !asc) {
				safe = false
				break
			}
		}
		prev = n
	}
	return safe
}

func solve1(lines []string) int {
	safeCount := 0

	reports := linesToNumbers(lines)
	for _, numbers := range reports {
		if isSafe(numbers) {
			safeCount += 1
		}
	}
	return safeCount
}

func solve2(lines []string) int {
	safeCount := 0

	reports := linesToNumbers(lines)
	for _, numbers := range reports {
		if isSafe(numbers) {
			safeCount += 1
		} else {
			for i := range len(numbers) {
				newNumbers := make([]int, 0)
				if i <= len(numbers)-1 {
					newNumbers = append(newNumbers, numbers[:i]...)
					newNumbers = append(newNumbers, numbers[i+1:]...)
				} else {
					newNumbers = append(newNumbers, numbers[:i]...)
				}
				if isSafe(newNumbers) {
					safeCount += 1
					break
				}
			}
		}
	}
	return safeCount
}

func Main() {
	fmt.Printf("Day 2\n=====\n")

	testLines, err := common.LoadFileLines("day2/testinput")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Test 1, expect 2: %d\n", solve1(testLines))
	fmt.Printf("Test 2, expect 4: %d\n", solve2(testLines))

	inputLines, err := common.LoadFileLines("day2/input")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution 1: %d\n", solve1(inputLines))
	fmt.Printf("Solution 2: %d\n", solve2(inputLines))
}
