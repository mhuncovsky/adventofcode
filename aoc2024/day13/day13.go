package day13

import (
	"aoc2024/common"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const SHOW = false

func AB(x0, x1, y0, y1, px, py int64) (int64, int64, error) {
	b := (x0*py - y0*px) / (x0*y1 - x1*y0)
	a := (px - b*x1) / x0
	if a*x0+b*x1 == px && a*y0+b*y1 == py {
		return a, b, nil
	}
	return 0, 0, errors.New("no solution")
}

func price(game string) int64 {
	x0, x1, y0, y1, px, py := parseGame(game)
	if a, b, err := AB(x0, x1, y0, y1, px, py); err == nil {
		return 3*a + b
	}
	return 0
}

func parseGame(game string) (int64, int64, int64, int64, int64, int64) {
	lines := strings.Split(game, "\n")
	if len(lines) != 3 {
		panic(fmt.Sprintf("unrecoverable error: cannot parse game, lines=%#v", lines))
	}

	ss := strings.SplitN(lines[0][12:], ", Y+", 2)
	x0, err := strconv.ParseInt(ss[0], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse x0, game=%#v", game))
	}
	y0, err := strconv.ParseInt(ss[1], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse y0, game=%#v", game))
	}

	ss = strings.SplitN(lines[1][12:], ", Y+", 2)
	x1, err := strconv.ParseInt(ss[0], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse x1, game=%#v", game))
	}
	y1, err := strconv.ParseInt(ss[1], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse y1, game=%#v", game))
	}

	ss = strings.SplitN(lines[2][9:], ", Y=", 2)
	px, err := strconv.ParseInt(ss[0], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse px, game=%#v", game))
	}
	py, err := strconv.ParseInt(ss[1], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse py, game=%#v", game))
	}

	return x0, x1, y0, y1, px, py
}

func solve1(text string) int64 {
	total := int64(0)
	text = strings.TrimRight(text, "\r\n ")
	games := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n\n")
	for _, game := range games {
		total += price(game)
	}
	return total
}

func solve2(text string) int64 {
	total := int64(0)
	o := int64(10000000000000)
	text = strings.TrimRight(text, "\r\n ")
	games := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n\n")
	for _, game := range games {
		x0, x1, y0, y1, px, py := parseGame(game)
		if a, b, err := AB(x0, x1, y0, y1, px+o, py+o); err == nil {
			total += 3*a + b
		}
	}
	return total
}

func Main() {
	fmt.Printf("Day 13\n=====\n")

	text, err := common.LoadFileText("day13/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 480: %d\n", solve1(text))
		fmt.Printf("Test 2, expect ???: %d\n", solve2(text))
	}

	text, err = common.LoadFileText("day13/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(text))
		fmt.Printf("Solution 2: %d\n", solve2(text))
	}
}
