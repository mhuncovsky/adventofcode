package day7

import (
	"aoc2024/common"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const SHOW = false

type queue struct {
	lock sync.Mutex
	q    []int
}

func (q *queue) Enqueue(v int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.q = append(q.q, v)
}

func (q *queue) Dequeue() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	v := q.q[0]
	q.q = q.q[1:]
	return v
}

func (q *queue) Empty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.q) <= 0
}

func (q *queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.q)
}

func NewQueue() *queue {
	return &queue{sync.Mutex{}, make([]int, 0)}
}

type Entry struct {
	result  int
	numbers []int
}

func NewEntry(result int, numbers []int) Entry {
	return Entry{result, numbers}
}

func splitInput(lines []string) []Entry {
	entries := make([]Entry, len(lines))

	for i, line := range lines {
		splitLine := strings.Split(line, ": ")
		if len(splitLine) != 2 {
			panic(fmt.Sprintf("Expected 2 parts after splitting a line on ': ', got %v", len(splitLine)))
		}

		r, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}

		ns := make([]int, 0)
		for _, strNum := range strings.Split(splitLine[1], " ") {
			if n, err := strconv.Atoi(strNum); err == nil {
				ns = append(ns, n)
			} else {
				panic(err)
			}
		}

		entries[i] = Entry{r, ns}
	}

	return entries
}

func isValid(e Entry) bool {
	q := NewQueue()
	ns := e.numbers
	for i := range len(e.numbers) {
		if i == 0 {
			q.Enqueue(ns[i])
			continue
		}
		for range q.Len() {
			v := q.Dequeue()
			q.Enqueue(v + ns[i])
			q.Enqueue(v * ns[i])
		}
	}
	for !q.Empty() {
		v := q.Dequeue()
		if v == e.result {
			return true
		}
	}
	return false
}

func mergeNumbers(a, b int) int {
	v, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic(fmt.Sprintf("Couldn't merge %d and %d, %v", a, b, err))
	}
	return v
}

func isValid2(e Entry) bool {
	q := NewQueue()
	ns := e.numbers
	for i := range len(e.numbers) {
		if i == 0 {
			q.Enqueue(ns[i])
			continue
		}
		for range q.Len() {
			v := q.Dequeue()
			q.Enqueue(v + ns[i])
			q.Enqueue(v * ns[i])
			q.Enqueue(mergeNumbers(v, ns[i]))
		}
	}
	for !q.Empty() {
		v := q.Dequeue()
		if v == e.result {
			return true
		}
	}
	return false
}

func parSolve2(lines []string) int {
	sum := 0
	lock := sync.Mutex{}
	entries := splitInput(lines)
	var wg sync.WaitGroup
	for _, e := range entries {
		wg.Add(1)
		go func(e Entry) {
			if isValid2(e) {
				lock.Lock()
				sum += e.result
				lock.Unlock()
			}
			wg.Done()
		}(e)
	}
	wg.Wait()
	return sum
}

func solve1(lines []string) int {
	sum := 0
	entries := splitInput(lines)
	for _, entry := range entries {
		if isValid(entry) {
			sum += entry.result
		}
	}
	return sum
}

func solve2(lines []string) int {
	// sum := 0
	// entries := splitInput(lines)
	// for _, entry := range entries {
	// 	if isValid2(entry) {
	// 		sum += entry.result
	// 	}
	// }
	// return sum
	return parSolve2(lines)
}

func Main() {
	fmt.Printf("Day 7\n=====\n")

	lines_test, err := common.LoadFileLines("day7/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect  3749: %d\n", solve1(lines_test))
		fmt.Printf("Test 2, expect 11387: %d\n", solve2(lines_test))
	}

	lines, err := common.LoadFileLines("day7/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(lines))
		fmt.Printf("Solution 2: %d\n", solve2(lines))
	}
}
