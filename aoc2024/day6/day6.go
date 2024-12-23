package day6

import (
	"aoc2024/common"
	"bufio"
	"fmt"
	"os"
	"slices"
)

const SHOW = false

type VecRC [2]int

var (
	Up    = VecRC{-1, 0}
	Down  = VecRC{1, 0}
	Left  = VecRC{0, -1}
	Right = VecRC{0, 1}
)

func NewVecRC(row, col int) VecRC {
	return VecRC{row, col}
}

func (v VecRC) row() int {
	return v[0]
}

func (v VecRC) col() int {
	return v[1]
}

func (a VecRC) add(b VecRC) VecRC {
	return NewVecRC(a[0]+b[0], a[1]+b[1])
}

func (v VecRC) rot90() VecRC {
	return NewVecRC(v.col(), -v.row())
}

type Predictor struct {
	tiles          [][]byte
	visited        []VecRC
	visitedWithDir map[[2]VecRC]struct{}
	startPos       VecRC
	pos            VecRC
	dir            VecRC
	finished       bool
	loop           bool
}

func NewPredictor(lines []string) *Predictor {
	tiles := make([][]byte, 0)
	visited := make([]VecRC, 0)
	visitedWithDir := make(map[[2]VecRC]struct{}, 0)
	pos := NewVecRC(0, 0)
	dir := Up
	for r, line := range lines {
		tiles = append(tiles, []byte(line))
		for c, b := range line {
			switch b {
			case '^':
				pos = NewVecRC(r, c)
				dir = Up
			case 'v':
				pos = NewVecRC(r, c)
				dir = Down
			case '<':
				pos = NewVecRC(r, c)
				dir = Left
			case '>':
				pos = NewVecRC(r, c)
				dir = Right
			}
		}
	}
	visited = append(visited, pos)
	visitedWithDir[[2]VecRC{pos, dir}] = struct{}{}
	return &Predictor{
		tiles:          tiles,
		visited:        visited,
		visitedWithDir: visitedWithDir,
		startPos:       pos,
		pos:            pos,
		dir:            dir,
		finished:       false,
		loop:           false,
	}
}

func (p *Predictor) rows() int {
	return len(p.tiles)
}

func (p *Predictor) cols() int {
	return len(p.tiles[0])
}

func (p *Predictor) value_rc(row, col int) byte {
	return p.tiles[row][col]
}

func (p *Predictor) value(pos VecRC) byte {
	return p.tiles[pos.row()][pos.col()]
}

func (p *Predictor) setValue(pos VecRC, v byte) {
	p.tiles[pos.row()][pos.col()] = v
}

func (p *Predictor) updateDir() {
	switch p.dir {
	case Up:
		p.setValue(p.pos, '^')
	case Down:
		p.setValue(p.pos, 'v')
	case Left:
		p.setValue(p.pos, '<')
	case Right:
		p.setValue(p.pos, '>')
	}
}

func (p *Predictor) step() {
	if !p.finished {
		newPos := p.pos.add(p.dir)
		if newPos.row() < 0 || newPos.col() < 0 || newPos.row() >= p.rows() || newPos.col() >= p.cols() {
			p.finished = true
			p.setValue(p.pos, 'X')
		} else if p.value(newPos) == '#' || p.value(newPos) == 'O' {
			p.dir = p.dir.rot90()
			p.updateDir()
		} else {
			p.setValue(p.pos, 'X')
			p.pos = newPos
			p.updateDir()
			p.visited = append(p.visited, newPos)
			p.visitedWithDir[[2]VecRC{newPos, p.dir}] = struct{}{}
		}
	}
}

func (p *Predictor) run() int {
	for !p.finished {
		p.step()
	}
	vis := make([]VecRC, 0)
	for _, v := range p.visited {
		if !slices.Contains(vis, v) {
			vis = append(vis, v)
		}
	}
	p.visited = vis
	return len(vis)
}

func (p *Predictor) show() {
	fmt.Println()
	for r := range len(p.tiles) {
		fmt.Println(string(p.tiles[r]))
	}
}

func input(prompt string) string {
	fmt.Printf("%s", prompt)
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return result
}

func (p *Predictor) isLoop() bool {
	for !p.finished {
		p.step()

		newPos := p.pos.add(p.dir)
		pd := [2]VecRC{newPos, p.dir}
		if _, visited := p.visitedWithDir[pd]; visited {
			return true
		}
	}
	return false
}

func solve1(lines []string) int {
	p := NewPredictor(lines)
	return p.run()
}

func isLoop(lines []string, pos VecRC, c chan bool) {
	pp := NewPredictor(lines)
	pp.setValue(pos, 'O')
	c <- pp.isLoop()
}

// Super slow, there has to be a better way...
func solve2(lines []string) int {
	sum := 0
	p := NewPredictor(lines)
	startPos := p.pos
	p.run()
	visited := p.visited
	is := slices.Index(visited, startPos)
	if is != -1 {
		visited = slices.Delete(visited, is, is+1)
	}
	// for _, pos := range visited {
	// 	pp := NewPredictor(lines)
	// 	pp.setValue(pos, 'O')
	// 	if pp.isLoop() {
	// 		sum += 1
	// 		// fmt.Printf("N=%v, pos=%v", i, pos)
	// 		// input(" > ")
	// 		// pp.show()
	// 	}
	// }
	c := make(chan bool)
	for _, pos := range visited {
		go isLoop(lines, pos, c)
	}
	for range visited {
		if <-c {
			sum += 1
		}
	}
	return sum
}

func Main() {
	fmt.Printf("Day 6\n=====\n")

	lines_test, err := common.LoadFileLines("day6/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 41: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect  6: %d\n", solve2(lines_test))
	}

	lines, err := common.LoadFileLines("day6/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines))
	}
}
