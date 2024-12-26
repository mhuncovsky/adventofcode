package day10

import (
	"aoc2024/common"
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

const SHOW = false

type vec2 [2]int

type TopoMap struct {
	data   [][]int
	xs     int
	ys     int
	starts []vec2
}

func NewTopoMap(xs, ys int) *TopoMap {
	data := make([][]int, xs)
	starts := make([]vec2, 0)
	for x := range xs {
		data[x] = make([]int, ys)
	}
	return &TopoMap{data, xs, ys, starts}
}

func (m *TopoMap) Load(lines []string) {
	xs := len(lines[0])
	ys := len(lines)

	for y := range ys {
		for x := range xs {
			r := lines[y][x]
			if unicode.IsDigit(rune(r)) {
				v, err := strconv.Atoi(string(r))
				if err != nil {
					panic(err)
				}
				m.data[x][y] = v
				if v == 0 {
					m.starts = append(m.starts, vec2{x, y})
				}
			} else {
				m.data[x][y] = -1
			}
		}
	}
}

func (m *TopoMap) Show() {
	for y := range m.ys {
		for x := range m.xs {
			v := m.data[x][y]
			if v >= 0 && v <= 9 {
				fmt.Print(strconv.Itoa(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (m *TopoMap) ShowPath(nodes []vec2) {
	for y := range m.ys {
		for x := range m.xs {
			v := "."
			if slices.Contains(nodes, vec2{x, y}) {
				v = strconv.Itoa(m.data[x][y])
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func (m *TopoMap) Get(pos vec2) int {
	return m.data[pos[0]][pos[1]]
}

func (m *TopoMap) IterAround(pos vec2, exclude []vec2) []vec2 {
	res := make([]vec2, 0)
	x := pos[0]
	y := pos[1]

	if x-1 >= 0 {
		n := vec2{x - 1, y}
		if !slices.Contains(exclude, n) {
			res = append(res, n)
		}
	}

	if x+1 < m.xs {
		n := vec2{x + 1, y}
		if !slices.Contains(exclude, n) {
			res = append(res, n)
		}
	}

	if y-1 >= 0 {
		n := vec2{x, y - 1}
		if !slices.Contains(exclude, n) {
			res = append(res, n)
		}
	}

	if y+1 < m.ys {
		n := vec2{x, y + 1}
		if !slices.Contains(exclude, n) {
			res = append(res, n)
		}
	}

	return res
}

func (m *TopoMap) FindAround(pos vec2, value int, exclude []vec2) []vec2 {
	found := make([]vec2, 0)
	for _, p := range m.IterAround(pos, exclude) {
		if m.Get(p) == value {
			found = append(found, p)
		}
	}
	return found
}

func (m *TopoMap) CalcTrailScore(start vec2) int {
	if m.Get(start) != 0 {
		return 0
	}
	type item struct {
		n vec2
		v int
	}
	q := make([]item, 0)
	q = append(q, item{start, 1})
	visited := make([]vec2, 0)

	reached9 := make(map[vec2]struct{})

	for len(q) > 0 {
		x := q[0]
		// fmt.Printf("p=%v v=%v lf=%v\n", x.n, m.Get(x.n), x.v)
		if len(q) == 1 {
			q = make([]item, 0)
		} else {
			q = q[1:]
		}
		// visited = append(visited, x.n)
		ns := m.FindAround(x.n, x.v, visited)
		for _, n := range ns {
			if x.v < 9 {
				q = append(q, item{n, x.v + 1})
			} else {
				if _, ok := reached9[n]; !ok {
					reached9[n] = struct{}{}
				}
			}
		}
	}
	return len(reached9)
}

func (m *TopoMap) CalcTrailScore2(start vec2) int {
	score := 0
	if m.Get(start) != 0 {
		return score
	}
	type item struct {
		n vec2
		v int
	}
	q := make([]item, 0)
	q = append(q, item{start, 1})
	visited := make([]vec2, 0)

	reached9 := make(map[vec2]struct{})

	for len(q) > 0 {
		x := q[0]
		if len(q) == 1 {
			q = make([]item, 0)
		} else {
			q = q[1:]
		}
		ns := m.FindAround(x.n, x.v, visited)
		for _, n := range ns {
			if x.v < 9 {
				q = append(q, item{n, x.v + 1})
			} else {
				if _, ok := reached9[n]; !ok {
					score += 1
				}
			}
		}
	}
	return score
}

func solve1(lines []string) int {
	xs := len(lines[0])
	ys := len(lines)
	sum := 0

	m := NewTopoMap(xs, ys)
	m.Load(lines)
	// m.Show()
	// fmt.Println()

	// fmt.Printf("score=%v\n", m.CalcTrailScore(vec2{2, 0}))

	for _, start := range m.starts {
		score := m.CalcTrailScore(start)
		// fmt.Printf("%v score=%v\n", start, score)
		sum += score
	}

	// exclude := make([]vec2, 0)
	// nodes := m.IterAround(vec2{1, 1}, exclude)
	// for _, n := range nodes {
	// 	nn := make([]vec2, 1)
	// 	nn[0] = n
	// 	m.ShowPath(nn)
	// 	fmt.Println()
	// }
	// m.ShowPath(nodes)

	return sum
}

func solve2(lines []string) int {
	sum := 0

	xs := len(lines[0])
	ys := len(lines)

	m := NewTopoMap(xs, ys)
	m.Load(lines)

	for _, start := range m.starts {
		score := m.CalcTrailScore2(start)
		sum += score
	}
	return sum
}

func Main() {
	fmt.Printf("Day 10\n=====\n")

	lines_test, err := common.LoadFileLines("day10/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 36: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect 81: %d\n", solve2(lines_test))
	}

	lines, err := common.LoadFileLines("day10/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines))
	}
}
