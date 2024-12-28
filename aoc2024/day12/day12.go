package day12

import (
	"aoc2024/common"
	"fmt"
	"slices"
)

var SHOW = false

type V [2]int

func (v *V) X() int {
	return v[0]
}

func (v *V) Y() int {
	return v[1]
}

func (v *V) Left() V {
	return V{v.X() - 1, v.Y()}
}

func (v *V) Right() V {
	return V{v.X() + 1, v.Y()}
}

func (v *V) Top() V {
	return V{v.X(), v.Y() - 1}
}

func (v *V) Bottom() V {
	return V{v.X(), v.Y() + 1}
}

func (v *V) TopLeft() V {
	return V{v.X() - 1, v.Y() - 1}
}

func (v *V) TopRight() V {
	return V{v.X() + 1, v.Y() - 1}
}

func (v *V) BottomLeft() V {
	return V{v.X() - 1, v.Y() + 1}
}

func (v *V) BottomRight() V {
	return V{v.X() + 1, v.Y() + 1}
}

type PlotProperties struct {
	Nodes       []V
	Perimeter   int
	Sides       int
	TopLeft     V
	BottomRight V
}

type Plots struct {
	plots []rune
	xs    int
	ys    int
}

func NewPlots() Plots {
	return Plots{make([]rune, 0), 0, 0}
}

func (p *Plots) Load(lines []string) {
	p.xs = len(lines[0])
	p.ys = len(lines)
	p.plots = make([]rune, p.xs*p.ys)
	for y := range p.ys {
		for x := range p.xs {
			offs := p.xs*y + x
			c := lines[y][x]
			p.plots[offs] = rune(c)
		}
	}
}

func (p *Plots) MapPlot(pos V) PlotProperties {
	perimeter := 0
	sides := 0
	topLeft := pos
	bottomRight := pos
	nodes := append(make([]V, 0), pos)
	nn := append(make([]V, 0), pos)
	r := p.Get(pos)
	for len(nn) > 0 {
		pp := nn[0]
		for _, n := range [4]V{pp.Left(), pp.Right(), pp.Bottom(), pp.Top()} {
			if slices.Contains(nodes, n) {
				continue
			} else if !(n.X() >= 0 && n.X() < p.xs && n.Y() >= 0 && n.Y() < p.ys) {
				perimeter += 1
				continue
			} else if p.Get(n) == r {
				nodes = append(nodes, n)
				nn = append(nn, n)
				if n[0] < topLeft[0] {
					topLeft[0] = n[0]
				}
				if n[1] < topLeft[1] {
					topLeft[1] = n[1]
				}
				if n[0] > bottomRight[0] {
					bottomRight[0] = n[0]
				}
				if n[1] > bottomRight[1] {
					bottomRight[1] = n[1]
				}
			} else {
				perimeter += 1
			}
		}
		nn = nn[1:]
	}
	return PlotProperties{nodes, perimeter, sides, topLeft, bottomRight}
}

func (p *Plots) MapPlot2(pos V) PlotProperties {
	pp := p.MapPlot(pos)
	// count changes in min/max in each col/row
	maxY := 0
	maxX := 0
	minX := 0
	minY := 0
	for i, n := range pp.Nodes {
		x, y := n.X(), n.Y()
		if i == 0 {
			maxX = x
			maxY = y
			minX = x
			minY = y
		} else {
			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
		}
	}

	pp.Sides = 0

	rowMin, rowMax := -1, -1
	for y := minY; y <= maxY; y++ {
		row := make([]int, 0)
		for _, n := range pp.Nodes {
			if n.Y() == y {
				row = append(row, n.X())
			}
		}
		currMin := slices.Min(row)
		currMax := slices.Max(row)
		if rowMax != currMax {
			rowMax = currMax
			pp.Sides++
		}
		if rowMin != currMin {
			rowMin = currMin
			pp.Sides++
		}
	}

	colMin, colMax := -1, -1
	for x := minX; x <= maxX; x++ {
		col := make([]int, 0)
		for _, n := range pp.Nodes {
			if n.X() == x {
				col = append(col, n.Y())
			}
		}
		currMin := slices.Min(col)
		currMax := slices.Max(col)
		if colMax != currMax {
			colMax = currMax
			pp.Sides++
		}
		if colMin != currMin {
			colMin = currMin
			pp.Sides++
		}
	}

	return pp
}

type Edge struct {
	start V
	end   V
	side  bool
}

func (p *Plots) MapPlot3(pos V) PlotProperties {
	pp := p.MapPlot(pos)
	pp.Sides = 0
	a := pp.TopLeft
	b := pp.BottomRight
	xes := make([]Edge, 0)
	for x := a[0]; x <= b[0]+1; x++ {
		var e *Edge = nil
		for y := a[1]; y <= b[1]; y++ {
			v := V{x, y}
			isPlot := slices.Contains(pp.Nodes, v) // precalc a map to make this faster
			isLeftPlot := slices.Contains(pp.Nodes, v.Left())
			if isPlot && !isLeftPlot || !isPlot && isLeftPlot {
				if e == nil {
					start := V{x, y}
					e = &Edge{start, start, isPlot}
				} else if e.side != isPlot {
					xes = append(xes, *e)
					e = nil
				} else {
					e.end = V{x, y}
				}
			} else { // if !isPlot || (isPlot && isLeftPlot)
				if e != nil {
					xes = append(xes, *e)
					// fmt.Printf("XEdge %v\n", e)
					e = nil
				}
			}
		}
		if e != nil {
			xes = append(xes, *e)
			// fmt.Printf("XEdge %v\n", e)
			e = nil
		}
	}
	yes := make([]Edge, 0)
	for y := a[1]; y <= b[1]+1; y++ {
		var e *Edge = nil
		for x := a[0]; x <= b[0]; x++ {
			v := V{x, y}
			isPlot := slices.Contains(pp.Nodes, v) // precalc a map to make this faster
			isTopPlot := slices.Contains(pp.Nodes, v.Top())
			if isPlot && !isTopPlot || !isPlot && isTopPlot {
				if e == nil {
					start := V{x, y}
					e = &Edge{start, start, isPlot}
				} else if e.side != isPlot {
					yes = append(yes, *e)
					e = nil
				} else {
					e.end = V{x, y}
				}
			} else {
				if e != nil {
					yes = append(yes, *e)
					// fmt.Printf("YEdge %v\n", e)
					e = nil
				}
			}
		}
		if e != nil {
			yes = append(yes, *e)
			// fmt.Printf("YEdge %v\n", e)
			e = nil
		}
	}
	pp.Sides += len(xes) + len(yes)
	if SHOW {
		fmt.Printf("xes=%v\n", xes)
		fmt.Printf("yes=%v\n", yes)
		p.ShowPlot(pp)
		fmt.Printf(
			"%s sides=%v area=%v price=%v\n",
			string(p.Get(pp.Nodes[0])),
			pp.Sides,
			len(pp.Nodes),
			pp.Sides*len(pp.Nodes),
		)
		// np := NewPlots()
		// np.xs = p.xs
		// np.ys = p.ys
		// np.plots = make([]rune, p.xs*p.ys)
		// for y := range p.ys {
		// 	for x := range p.xs {
		// 		np.plots[y*np.xs+x] = '.'
		// 	}
		// }
		// for i, e := range xes {
		// 	x := e.start.X()
		// 	r := rune('a' + i)
		// 	for y := e.start.Y(); y <= e.end.Y(); y++ {
		// 		np.plots[y*np.xs+x] = r
		// 	}
		// }
		// for i, e := range yes {
		// 	y := e.start.Y()
		// 	r := rune('a' + i)
		// 	for x := e.start.X(); x <= e.end.X(); x++ {
		// 		np.plots[y*np.xs+x] = r
		// 	}
		// }
		// np.ShowPlots()
	}
	return pp
}

func (p *Plots) MapPlot4(pos V) PlotProperties {
	pp := p.MapPlot(pos)
	pp.Sides = 0
	corners := 0
	// this := p.Get(pos)
	same := func(n V) bool {
		return slices.Contains(pp.Nodes, n)
	}
	for _, n := range pp.Nodes {
		if same(n.Left()) && same(n.Top()) && !same(n.TopLeft()) || !same(n.Left()) && !same(n.Top()) {
			corners += 1
		}
		if same(n.Right()) && same(n.Top()) && !same(n.TopRight()) || !same(n.Right()) && !same(n.Top()) {
			corners += 1
		}
		if same(n.Left()) && same(n.Bottom()) && !same(n.BottomLeft()) || !same(n.Left()) && !same(n.Bottom()) {
			corners += 1
		}
		if same(n.Right()) && same(n.Bottom()) && !same(n.BottomRight()) || !same(n.Right()) && !same(n.Bottom()) {
			corners += 1
		}
	}
	pp.Sides = corners
	// if SHOW {
	// 	p.ShowPlot(pp)
	// 	fmt.Printf(
	// 		"%s sides=%v area=%v price=%v\n",
	// 		string(p.Get(pp.Nodes[0])),
	// 		pp.Sides,
	// 		len(pp.Nodes),
	// 		pp.Sides*len(pp.Nodes),
	// 	)
	// }
	return pp
}

func (p *Plots) Get(pos V) rune {
	return p.plots[pos.X()+pos.Y()*p.xs]
}

func (p *Plots) ShowPlot(pp PlotProperties) {
	if !SHOW {
		return
	}
	// fmt.Printf("%v:\n", pp)
	for y := range p.ys {
		for x := range p.xs {
			if slices.Contains(pp.Nodes, V{x, y}) {
				fmt.Printf("%s", string(p.Get(V{x, y})))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (p *Plots) ShowPlots() {
	if !SHOW {
		return
	}
	// fmt.Printf("%v:\n", pp)
	for y := range p.ys {
		for x := range p.xs {
			fmt.Printf("%s", string(p.Get(V{x, y})))
		}
		fmt.Println()
	}
}

func solve1(lines []string) int {
	sum := 0

	p := NewPlots()
	p.Load(lines)

	visited := make([]V, 0)
	for x := range p.xs {
		for y := range p.ys {
			start := V{x, y}
			if !slices.Contains(visited, start) {
				pp := p.MapPlot2(start)
				visited = append(visited, pp.Nodes...)
				sum += pp.Perimeter * len(pp.Nodes)
				// p.ShowPlot(pp)
			}
		}
	}

	return sum
}

func solve2(lines []string) int {
	sum := 0

	p := NewPlots()
	p.Load(lines)
	p.ShowPlots()

	visited := make([]V, 0)
	for x := range p.xs {
		for y := range p.ys {
			start := V{x, y}
			if !slices.Contains(visited, start) {
				pp := p.MapPlot4(start)
				visited = append(visited, pp.Nodes...)
				sum += pp.Sides * len(pp.Nodes)
				// p.ShowPlot(pp)
			}
		}
	}

	return sum
}

func countSides(lines []string, start V) int {
	p := NewPlots()
	p.Load(lines)
	pp := p.MapPlot4(start)
	return pp.Sides
}

func countSides3(lines []string, start V) int {
	p := NewPlots()
	p.Load(lines)
	pp := p.MapPlot3(start)
	return pp.Sides
}

func compare34(lines []string) int {
	sum := 0

	p := NewPlots()
	p.Load(lines)
	p.ShowPlots()

	visited := make([]V, 0)
	for x := range p.xs {
		for y := range p.ys {
			start := V{x, y}
			if !slices.Contains(visited, start) {
				pp4 := p.MapPlot4(start)
				pp3 := p.MapPlot3(start)
				visited = append(visited, pp4.Nodes...)
				sum += pp4.Sides * len(pp4.Nodes)
				if pp4.Sides != pp3.Sides {
					fmt.Println()
					fmt.Printf("sides3=%v sides4=%v\n", pp3.Sides, pp4.Sides)
					p.ShowPlot(pp4)
					fmt.Scanf("\n")
				}
			}
		}
	}

	return sum
}

func Main() {
	fmt.Printf("Day 12\n=====\n")

	// SHOW = true
	testinput, err := common.LoadFileLines("day12/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1A, expect  140: %d\n", solve1(testinput))
		fmt.Printf("Test 2A, expect   80: %d\n", solve2(testinput))
	}

	SHOW = false
	testinput2, err := common.LoadFileLines("day12/testinput2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1B, expect  772: %d\n", solve1(testinput2))
		fmt.Printf("Test 2B, expect  436: %d\n", solve2(testinput2))
	}

	testinput3, err := common.LoadFileLines("day12/testinput3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1C, expect 1930: %d\n", solve1(testinput3))
		fmt.Printf("Test 2C, expect 1206: %d\n", solve2(testinput3))
	}

	testinput4, err := common.LoadFileLines("day12/testinput4")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 2D, expect  368: %d\n", solve2(testinput4))
	}

	testinput5, err := common.LoadFileLines("day12/testinput5")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 2E, expect  236: %d\n", solve2(testinput5))
	}

	sol2extract1, err := common.LoadFileLines("day12/sol2extract1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 2F, expect  40: %d\n", countSides(sol2extract1, V{8, 1}))
	}

	T1Lines := make([]string, 0)
	T1Lines = append(T1Lines, "********")
	T1Lines = append(T1Lines, "***AAA**")
	T1Lines = append(T1Lines, "*AAA*A**")
	T1Lines = append(T1Lines, "*****AA*")
	T1Lines = append(T1Lines, "********")
	// SHOW = true
	fmt.Printf("Test T1, expect  604: %d\n", solve2(T1Lines))
	// SHOW = false

	lines, err := common.LoadFileLines("day12/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		s2 := solve2(lines)
		fmt.Printf("Solution 2: %d\n", s2) // WRONG: 836289
		// fmt.Printf("Compare 34:\n")
		// SHOW = true
		// compare34(lines)
		// SHOW = false
	}

}
