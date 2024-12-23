package day3

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"unicode"
)

type Memory struct {
	pos  int
	data *string
}

type Instruction int

const (
	InstructionMul Instruction = iota
	InstructionDo
	InstructionDoNot
	InstructionUnknown
)

func InstructionName(inst Instruction) string {
	switch inst {
	case InstructionMul:
		return "InstructionMul"
	case InstructionDo:
		return "InstructionDo"
	case InstructionDoNot:
		return "InstructionDoNot"
	case InstructionUnknown:
		return "InstructionUnknown"
	default:
		return "InstructionUnknown"
	}
}

func (m *Memory) readMul() bool {
	for {
		if m.data == nil || m.pos >= len(*m.data)-4 || m.pos < 0 {
			break
		}
		if (*m.data)[m.pos:m.pos+4] == "mul(" {
			m.pos += 4
			return true
		}
		m.pos++
	}
	return false
}

func (m *Memory) readInstruction() (bool, Instruction) {
	for {
		if m.data == nil || m.pos < 0 || m.pos >= len(*m.data) {
			break
		}

		if m.pos <= len(*m.data)-7 {
			if (*m.data)[m.pos:m.pos+7] == "don't()" {
				m.pos += 7
				return true, InstructionDoNot
			}
		}

		if m.pos <= len(*m.data)-4 {
			s := (*m.data)[m.pos : m.pos+4]
			if s == "mul(" {
				m.pos += 4
				return true, InstructionMul
			} else if s == "do()" {
				m.pos += 4
				return true, InstructionDo
			}
		}

		m.pos++
	}

	return false, InstructionUnknown
}

func (m *Memory) readDigit(end byte) (bool, int) {
	if m.data == nil || m.pos >= len(*m.data) || m.pos < 0 {
		return false, 0
	}

	digit := make([]byte, 0)
	for _ = range 3 {
		if m.data == nil || m.pos >= len(*m.data) || m.pos < 0 {
			return false, 0
		}

		c := (*m.data)[m.pos]
		if unicode.IsDigit(rune(c)) {
			digit = append(digit, c)
		} else {
			break
		}

		m.pos++
	}

	c := (*m.data)[m.pos]
	if c == end {
		if len(digit) == 0 || len(digit) > 3 {
			return false, 0
		}

		v, err := strconv.ParseInt(string(digit), 10, 64)
		if err != nil {
			panic(err)
		}
		m.pos++
		return true, int(v)
	}

	return false, 0
}

func solve1(text string) int {
	sum := 0
	m := Memory{pos: 0, data: &text}
	for {
		ok := m.readMul()
		if !ok {
			break
		}

		ok, a := m.readDigit(',')
		if !ok {
			continue
		}
		// fmt.Printf("a: %d\n", a)

		ok, b := m.readDigit(')')
		if !ok {
			continue
		}
		// fmt.Printf("b: %d\n", b)
		sum += a * b
	}
	return sum
}

func solve2(text string) int {
	do := true
	sum := 0
	m := Memory{pos: 0, data: &text}
	for {
		ok, inst := m.readInstruction()
		if !ok {
			break
		}
		// fmt.Printf("%v, pos=%d\n", InstructionName(inst), m.pos)
		switch inst {
		case InstructionMul:
			ok, a := m.readDigit(',')
			if !ok {
				continue
			}
			ok, b := m.readDigit(')')
			if !ok {
				continue
			}
			if do {
				sum += a * b
			}
		case InstructionDo:
			do = true
		case InstructionDoNot:
			do = false
		}
	}
	return sum
}

func Main() {
	fmt.Printf("Day 3\n=====\n")

	text_test, err := common.LoadFileText("day3/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 161: %d\n", solve1(text_test))
	}

	text_test2, err := common.LoadFileText("day3/testinput2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 2, expect  48: %d\n", solve2(text_test2))
	}

	text, err := common.LoadFileText("day3/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(text))
		fmt.Printf("Solution 2: %d\n", solve2(text))
	}
}
