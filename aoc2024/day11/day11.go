package day11

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
)

const SHOW = false

type Stones struct {
	stones []int
}

func NewStones() *Stones {
	return &Stones{make([]int, 0)}
}

func (s *Stones) Load(lines []string) {
	strStones := strings.Split(lines[0], " ")
	for _, ss := range strStones {
		v, err := strconv.Atoi(ss)
		if err != nil {
			panic(err)
		}
		s.stones = append(s.stones, v)
	}
}

func (s *Stones) Show() {
	if !SHOW {
		return
	}
	for _, n := range s.stones {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func (s *Stones) Blink() {
	ns := make([]int, 0)
	for _, stone := range s.stones {
		if stone == 0 {
			ns = append(ns, 1)
			continue
		}

		ss := strconv.Itoa(stone)
		if len(ss)%2 == 0 {
			ofs := len(ss) / 2
			a, err := strconv.Atoi(ss[:ofs])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(ss[ofs:])
			if err != nil {
				panic(err)
			}
			ns = append(ns, a, b)
			continue
		}
		ns = append(ns, stone*2024)
	}
	s.stones = ns
}

func (s *Stones) Count() int {
	return len(s.stones)
}

func Blink(stone int) []int {
	ns := make([]int, 0)
	if stone == 0 {
		ns = append(ns, 1)
		return ns
	}

	ss := strconv.Itoa(stone)
	if len(ss)%2 == 0 {
		ofs := len(ss) / 2
		a, err := strconv.Atoi(ss[:ofs])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(ss[ofs:])
		if err != nil {
			panic(err)
		}
		ns = append(ns, a, b)
		return ns
	}

	ns = append(ns, stone*2024)
	return ns
}

type BetterStoneBlinker struct {
	stones map[int]int
}

func NewBetterStoneBlinker() BetterStoneBlinker {
	return BetterStoneBlinker{make(map[int]int)}
}

func (sb *BetterStoneBlinker) Add(stone int) {
	sb.stones[stone] += 1
}

func (sb *BetterStoneBlinker) Blink() {
	stones := make(map[int]int, len(sb.stones))
	for stone, count := range sb.stones {
		stones[stone] = count
	}

	for k := range sb.stones {
		delete(sb.stones, k)
	}

	for stone, count := range stones {
		ns := Blink(stone)
		for _, s := range ns {
			sb.stones[s] += count
		}
	}
}

func (sb *BetterStoneBlinker) BlinkNCount(n int) int {
	for range n {
		sb.Blink()
	}
	return sb.Count()
}

func (sb *BetterStoneBlinker) Count() int {
	total := 0
	for _, count := range sb.stones {
		total += count
	}
	return total
}

func (sb *BetterStoneBlinker) Show() {
	stones := make([]int, 0)
	for stone, count := range sb.stones {
		for range count {
			stones = append(stones, stone)
		}
	}
	fmt.Printf("count=%v stones=[%v]\n", sb.Count(), stones)
}

func (sb *BetterStoneBlinker) Load(lines []string) {
	strStones := strings.Split(lines[0], " ")
	for _, ss := range strStones {
		v, err := strconv.Atoi(ss)
		if err != nil {
			panic(err)
		}
		sb.Add(v)
	}
}

func solve1(lines []string) int {
	s := NewStones()
	s.Load(lines)
	s.Show()
	for range 25 {
		s.Blink()
		s.Show()
	}
	return s.Count()
}

func solve2(lines []string, n int) int {
	sb := NewBetterStoneBlinker()
	sb.Load(lines)
	return sb.BlinkNCount(n)
}

func Main() {
	fmt.Printf("Day 11\n=====\n")

	lines_test, err := common.LoadFileLines("day11/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 55312: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect 55312: %d\n", solve2(lines_test, 25))
	}

	lines, err := common.LoadFileLines("day11/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines, 75))
	}
}
