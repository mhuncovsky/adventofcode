package day5

import (
	"aoc2024/common"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const SHOW = false

type Ruler struct {
	rules []string
}

func NewRuler(rules []string) *Ruler {
	r := &Ruler{rules: rules}
	slices.Sort(r.rules)
	return r
}

func (r *Ruler) ok(a, b string) bool {
	_, found := slices.BinarySearch(r.rules, fmt.Sprintf("%s|%s", b, a))
	// fmt.Printf("%s|%s: ok=%v\n", b, a, !found)
	return !found
}

func solve1(lines []string) int {
	rules := make([]string, 0)
	print := make([]string, 0)
	orderDone := false
	for _, line := range lines {
		if !orderDone {
			if line == "" {
				orderDone = true
				continue
			}
			rules = append(rules, line)
		} else {
			print = append(print, line)
		}
	}

	ruler := NewRuler(rules)

	sum := 0
	for _, line := range print {
		allOk := true
		ps := strings.Split(line, ",")
		if len(ps) <= 1 {
			continue
		}
		for i := range len(ps) {
			if i == 0 {
				continue
			}
			if !ruler.ok(ps[i-1], ps[i]) {
				allOk = false
			}
		}
		if allOk {
			v, err := strconv.ParseInt(ps[len(ps)/2], 10, 64)
			if err != nil {
				panic(err)
			}
			sum += int(v)
			if SHOW {
				fmt.Printf("%s, ok=true, n=%d\n", line, v)
			}
		} else if SHOW {
			fmt.Printf("%s, ok=false\n", line)
		}
	}

	return sum
}

func solve2(lines []string) int {
	rules := make([]string, 0)
	print := make([]string, 0)
	orderDone := false
	for _, line := range lines {
		if !orderDone {
			if line == "" {
				orderDone = true
				continue
			}
			rules = append(rules, line)
		} else {
			print = append(print, line)
		}
	}

	ruler := NewRuler(rules)

	sum := 0
	badPrints := make([]string, 0)
	for _, line := range print {
		ps := strings.Split(line, ",")
		if len(ps) <= 1 {
			continue
		}
		for i := range len(ps) {
			if i == 0 {
				continue
			}
			if !ruler.ok(ps[i-1], ps[i]) {
				badPrints = append(badPrints, line)
				break
			}
		}
	}
	for _, line := range badPrints {
		ps := strings.Split(line, ",")
		if len(ps) <= 1 {
			continue
		}
		for {
			allOk := true
			for i := range len(ps) {
				if i == 0 {
					continue
				}
				if !ruler.ok(ps[i-1], ps[i]) {
					temp := ps[i-1]
					ps[i-1] = ps[i]
					ps[i] = temp
					allOk = false
				}
			}
			if allOk {
				break
			}
		}

		v, err := strconv.ParseInt(ps[len(ps)/2], 10, 64)
		if err != nil {
			panic(err)
		}
		sum += int(v)
	}

	return sum
}

func Main() {
	fmt.Printf("Day 5\n=====\n")

	lines_test, err := common.LoadFileLines("day5/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 143: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect 123: %d\n", solve2(lines_test))
	}

	lines, err := common.LoadFileLines("day5/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines))
	}
}
