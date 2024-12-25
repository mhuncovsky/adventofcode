package day8

import (
	"aoc2024/common"
	"fmt"
)

const SHOW = false

type vec2 [2]int

func uniqueCombinations(vs []vec2) [][2]vec2 {
	c := make([][2]vec2, 0)
	for i := range len(vs) {
		for j := range len(vs) - i - 1 {
			c = append(c, [2]vec2{vs[i], vs[i+j+1]})
		}
	}
	return c
}

func solve1(lines []string) int {
	sum := 0

	sx := len(lines[0])
	sy := len(lines)

	antenas := make(map[rune][]vec2, 0)
	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				vv, ok := antenas[ch]
				if !ok {
					vv = make([]vec2, 0)
				}
				antenas[ch] = append(vv, vec2{x, y})
			}
		}
	}

	antinodesM := make(map[vec2]struct{})
	for _, v := range antenas {
		comb := uniqueCombinations(v)
		for _, c := range comb {
			a, b := c[0], c[1]
			dx := a[0] - b[0]
			dy := a[1] - b[1]
			antinodesM[vec2{a[0] + dx, a[1] + dy}] = struct{}{}
			antinodesM[vec2{b[0] - dx, b[1] - dy}] = struct{}{}
		}
	}
	antinodesK := make(map[vec2]struct{})
	for k := range antinodesM {
		if k[0] >= 0 && k[0] < sx && k[1] >= 0 && k[1] < sy {
			sum += 1
			antinodesK[k] = struct{}{}
		}
	}

	return sum
}

func solve2(lines []string) int {
	sum := 0

	sx := len(lines[0])
	sy := len(lines)

	antenas := make(map[rune][]vec2, 0)
	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				vv, ok := antenas[ch]
				if !ok {
					vv = make([]vec2, 0)
				}
				antenas[ch] = append(vv, vec2{x, y})
			}
		}
	}

	antinodes := make(map[vec2]struct{})
	for _, v := range antenas {
		comb := uniqueCombinations(v)
		for _, c := range comb {
			a, b := c[0], c[1]
			dx := a[0] - b[0]
			dy := a[1] - b[1]

			antinodes[a] = struct{}{}
			antinodes[b] = struct{}{}

			m := 1
			for {
				x := a[0] + dx*m
				y := a[1] + dy*m
				if x >= 0 && y >= 0 && x < sx && y < sy {
					antinodes[vec2{x, y}] = struct{}{}
					m++
				} else {
					break
				}
			}

			m = 1
			for {
				x := a[0] - dx*m
				y := a[1] - dy*m
				if x >= 0 && y >= 0 && x < sx && y < sy {
					antinodes[vec2{x, y}] = struct{}{}
					m++
				} else {
					break
				}
			}

		}
	}
	for range antinodes {
		sum += 1
	}

	return sum
}

func Main() {
	fmt.Printf("Day 8\n=====\n")

	lines_test, err := common.LoadFileLines("day8/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 14: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect 34: %d\n", solve2(lines_test))
	}

	lines, err := common.LoadFileLines("day8/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines))
	}
}
