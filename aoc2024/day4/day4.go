package day4

import (
	"aoc2024/common"
	"fmt"
	"slices"
)

const SHOW = false

type Pos struct {
	r, c int
}

func NewPos(r, c int) Pos {
	return Pos{r, c}
}

type Puzzle struct {
	lines []string
}

func NewPuzzle(lines []string) *Puzzle {
	return &Puzzle{lines: lines}
}

func (pz *Puzzle) char(pos Pos) byte {
	return pz.lines[pos.r][pos.c]
}

func (pz *Puzzle) showPaths(ps []Pos) {
	fmt.Println()
	for r := range len(pz.lines) {
		lineBytes := make([]byte, len(pz.lines[r]))
		for c := range len(pz.lines[0]) {
			pos := NewPos(r, c)
			if slices.Contains(ps, pos) {
				lineBytes = append(lineBytes, pz.char(pos))
			} else {
				lineBytes = append(lineBytes, '.')
			}
		}
		fmt.Println(string(lineBytes))
	}
}

func (pz *Puzzle) show() {
	fmt.Println()
	for r := range len(pz.lines) {
		fmt.Println(pz.lines[r])
	}
}

func (pz *Puzzle) findAround(pos Pos, c byte, exclude []Pos) []Pos {
	result := make([]Pos, 0)
	for _, p := range pz.iterAround(pos, exclude) {
		if pz.char(p) == c {
			result = append(result, p)
		}
	}
	return result
}

func (pz *Puzzle) iterAround(pos Pos, exclude []Pos) []Pos {
	res := make([]Pos, 0)
	for rr := range 3 {
		for cc := range 3 {
			r := pos.r + rr - 1
			c := pos.c + cc - 1
			if r >= len(pz.lines) || r < 0 || c >= len(pz.lines[0]) || c < 0 {
				continue
			}
			n := NewPos(r, c)
			if n != pos && !slices.Contains(exclude, n) {
				res = append(res, n)
			}
		}
	}
	return res
}

func findAll(lines []string, r, c int) int {
	sum := 0
	pz := NewPuzzle(lines)
	pos := NewPos(r, c)
	visited := make([]Pos, 0)
	visited = append(visited, pos)

	if pz.char(pos) != 'X' {
		return 0
	}

	ms := pz.findAround(pos, 'M', visited)
	for _, m := range ms {
		as := pz.findAround(m, 'A', visited)
		for _, a := range as {
			ss := pz.findAround(a, 'S', visited)
			// for _, s := range ss {
			// 	sum += 1
			// 	pz.showPaths([]Pos{pos, m, a, s})
			// }
			for range ss {
				sum += 1
			}
		}
	}

	return sum
}

func horizontal(lines []string, c, r int) int {
	if r >= len(lines) || c+4 > len(lines[0]) {
		return 0
	}
	s := lines[r][c : c+4]
	if s == "XMAS" || s == "SAMX" {
		if SHOW {
			pz := NewPuzzle(lines)
			pz.showPaths([]Pos{
				NewPos(r, c),
				NewPos(r, c+1),
				NewPos(r, c+2),
				NewPos(r, c+3),
			})
		}
		return 1
	}
	return 0
}

func diagonal_right(lines []string, c, r int) int {
	if r+4 > len(lines) || c+4 > len(lines[0]) {
		return 0
	}
	a := make([]byte, 4)
	for i := range 4 {
		a[i] = lines[r+i][c+i]
	}
	s := string(a)
	if s == "XMAS" || s == "SAMX" {
		if SHOW {
			pz := NewPuzzle(lines)
			pz.showPaths([]Pos{
				NewPos(r, c),
				NewPos(r+1, c+1),
				NewPos(r+2, c+2),
				NewPos(r+3, c+3),
			})
		}
		return 1
	}
	return 0
}

func diagonal_left(lines []string, c, r int) int {
	if r+4 > len(lines) || c-3 < 0 {
		return 0
	}
	a := make([]byte, 4)
	for i := range 4 {
		a[i] = lines[r+i][c-i]
	}
	s := string(a)
	if s == "XMAS" || s == "SAMX" {
		if SHOW {
			pz := NewPuzzle(lines)
			pz.showPaths([]Pos{
				NewPos(r, c),
				NewPos(r+1, c-1),
				NewPos(r+2, c-2),
				NewPos(r+3, c-3),
			})
		}
		return 1
	}
	return 0
}

func vertical(lines []string, c, r int) int {
	if r+4 > len(lines) || c > len(lines[0]) {
		return 0
	}
	a := make([]byte, 4)
	for i := range 4 {
		a[i] = lines[r+i][c]
	}
	s := string(a)
	if s == "XMAS" || s == "SAMX" {
		if SHOW {
			pz := NewPuzzle(lines)
			pz.showPaths([]Pos{
				NewPos(r, c),
				NewPos(r+1, c),
				NewPos(r+2, c),
				NewPos(r+3, c),
			})
		}
		return 1
	}
	return 0
}

func solve1(lines []string) int {
	sum := 0

	// pz := NewPuzzle(lines)
	// pz.show()

	for r := range len(lines[0]) {
		for c := range len(lines) {
			sum += horizontal(lines, c, r)
			sum += diagonal_right(lines, c, r)
			sum += diagonal_left(lines, c, r)
			sum += vertical(lines, c, r)
		}
	}

	return sum
}

func (pz *Puzzle) xmas(r, c int) int {
	if r-1 < 0 || c-1 < 0 || r+1 >= len(pz.lines) || c+1 >= len(pz.lines[0]) {
		return 0
	}

	sum := 0
	a := [3]byte{'M', 'A', 'S'}
	b := [3]byte{'S', 'A', 'M'}

	if pz.char(NewPos(r, c)) != 'A' {
		return sum
	}

	left := [3]byte{
		pz.char(NewPos(r-1, c-1)),
		pz.char(NewPos(r, c)),
		pz.char(NewPos(r+1, c+1)),
	}

	right := [3]byte{
		pz.char(NewPos(r-1, c+1)),
		pz.char(NewPos(r, c)),
		pz.char(NewPos(r+1, c-1)),
	}

	if (left == a || left == b) && (right == a || right == b) {
		if SHOW {
			pp := [5]Pos{
				NewPos(r-1, c-1),
				NewPos(r, c),
				NewPos(r+1, c+1),
				NewPos(r-1, c+1),
				NewPos(r+1, c-1),
			}
			pz.showPaths(pp[:5])
		}
		sum += 1
	}

	return sum
}

func solve2(lines []string) int {
	sum := 0
	pz := NewPuzzle(lines)

	for r := range len(lines[0]) {
		for c := range len(lines) {
			sum += pz.xmas(r, c)
		}
	}

	return sum
}

func Main() {
	fmt.Printf("Day 4\n=====\n")

	lines_test, err := common.LoadFileLines("day4/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 18: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect  9: %d\n", solve2(lines_test))
	}

	lines, err := common.LoadFileLines("day4/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines))
	}
}
