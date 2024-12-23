package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve1(a []int, b []int) int {
	slices.Sort(a)
	slices.Sort(b)
	sum := 0
	for i := range len(a) {
		if a[i] > b[i] {
			sum += a[i] - b[i]
		} else if b[i] > a[i] {
			sum += b[i] - a[i]
		}
	}
	return sum
}

func solve2(a, b []int) int {
	m := make(map[int]int)
	for _, n := range b {
		v, ok := m[n]
		if ok {
			m[n] = v + 1
		} else {
			m[n] = 1
		}
	}
	sum := 0
	for _, n := range a {
		if v, ok := m[n]; ok {
			sum += n * v
		}
	}
	return sum
}

func loadNumbers(path string) ([]int, []int, error) {

	aa := make([]int, 0)
	bb := make([]int, 0)

	f, err := os.Open(path)
	if err != nil {
		return aa, bb, err
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		parts := strings.SplitN(line, "   ", 2)
		x, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return aa, bb, err
		}
		y, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return aa, bb, err
		}
		aa = append(aa, int(x))
		bb = append(bb, int(y))
		// fmt.Printf("%s\n", parts)
	}

	return aa, bb, nil
}

func Main() {
	fmt.Println("Day 1")
	fmt.Println("=====")

	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	fmt.Printf("Test 1, expect 11: %d\n", solve1(a, b))
	fmt.Printf("Test 2, expect 31: %d\n", solve2(a, b))

	aa, bb, err := loadNumbers("day1/input")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution 1: %d\n", solve1(aa, bb))
	fmt.Printf("Solution 2: %d\n", solve2(aa, bb))

}
