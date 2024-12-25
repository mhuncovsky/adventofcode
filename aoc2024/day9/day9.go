package day9

import (
	"aoc2024/common"
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

var SHOW = false

type fsEntry struct {
	offset int
	length int
}

func NewEntry(offset, length int) fsEntry {
	return fsEntry{offset, length}
}

type FileSystem struct {
	data  []int
	files []fsEntry
	frees []fsEntry
}

func NewFileSystem() *FileSystem {
	return &FileSystem{make([]int, 0), make([]fsEntry, 0), make([]fsEntry, 0)}
}

func (fs *FileSystem) Load(text string) {
	id := 0
	isFile := true
	for _, r := range text {
		if !unicode.IsDigit(r) {
			continue
		}
		blocks, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		if isFile {
			if blocks > 0 {
				fs.files = append(fs.files, NewEntry(len(fs.data), blocks))
			}
			for range blocks {
				fs.data = append(fs.data, id)
			}
			id++
		} else {
			if blocks > 0 {
				fs.frees = append(fs.frees, NewEntry(len(fs.data), blocks))
			}
			for range blocks {
				fs.data = append(fs.data, -1)
			}
		}
		isFile = !isFile
	}
}

func (fs *FileSystem) MoveOneFragment() {
	if !fs.HasFree() || len(fs.files) <= 0 {
		return
	}
	src := &fs.files[len(fs.files)-1]
	dst := &fs.frees[0]
	srcOfs := src.offset + src.length - 1
	dstOfs := dst.offset

	// fs.Show()
	// for i := range fs.data {
	// 	if i == srcOfs {
	// 		fmt.Print("s")
	// 	} else if i == dstOfs {
	// 		fmt.Print("d")
	// 	} else {
	// 		fmt.Print(" ")
	// 	}
	// }
	// fmt.Println()

	if srcOfs < dstOfs {
		fs.frees = make([]fsEntry, 0)
		return
	}

	fs.data[dstOfs] = fs.data[srcOfs]
	fs.data[srcOfs] = -1

	src.length--
	if src.length <= 0 {
		fs.files = fs.files[:len(fs.files)-1]
	}

	dst.length--
	dst.offset++
	if dst.length <= 0 {
		fs.frees = fs.frees[1:]
	}

}

func (fs *FileSystem) MoveOneFile() bool {
	if !fs.HasFree() || len(fs.files) <= 0 {
		return false
	}

	var src *fsEntry = nil
	var dst *fsEntry = nil
	dstIdx := 0

	for {
		if len(fs.files) <= 0 {
			return false
		}

		src = &fs.files[len(fs.files)-1]

		for i := range len(fs.frees) {
			maybeDst := &fs.frees[i]
			if maybeDst.length >= src.length && src.offset > maybeDst.offset {
				dst = maybeDst
				dstIdx = i
				break
			}
		}

		fs.files = fs.files[:len(fs.files)-1]

		if dst != nil {
			break
		}

	}

	srcOfs := src.offset
	dstOfs := dst.offset

	if SHOW {
		for i := range fs.data {
			if i >= srcOfs && i < srcOfs+src.length {
				fmt.Print("s")
			} else if i >= dstOfs && i < dstOfs+src.length {
				fmt.Print("d")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := range src.length {
		fs.data[dstOfs+i] = fs.data[srcOfs+i]
		fs.data[srcOfs+i] = -1
	}

	dst.length -= src.length
	dst.offset += src.length
	if dst.length <= 0 {
		fs.frees = slices.Delete(fs.frees, dstIdx, dstIdx+1)
	}

	return true
}

func (fs *FileSystem) Show() {
	if !SHOW {
		return
	}
	for _, v := range fs.data {
		if v >= 0 {
			fmt.Printf("%v", strconv.Itoa(v%10))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func (fs *FileSystem) HasFree() bool {
	return len(fs.frees) > 0
}

func (fs *FileSystem) Checksum() int {
	sum := 0
	for i, v := range fs.data {
		if v > 0 {
			sum += i * v
		}
	}
	return sum
}

func solve1(text string) int {
	fs := NewFileSystem()
	fs.Load(text)
	fs.Show()
	for fs.HasFree() {
		fs.MoveOneFragment()
		fs.Show()
	}

	return fs.Checksum()
}

func solve2(text string) int {
	fs := NewFileSystem()
	fs.Load(text)
	fs.Show()

	for fs.HasFree() {
		if !fs.MoveOneFile() {
			break
		}
		fs.Show()
	}

	return fs.Checksum()
}

func Main() {
	fmt.Printf("Day 9\n=====\n")

	SHOW = false

	text, err := common.LoadFileText("day9/testinput")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Test 1, expect 1928: %d\n", solve1(text))
		fmt.Printf("Test 2, expect 2858: %d\n", solve2(text))
	}

	SHOW = false

	text, err = common.LoadFileText("day9/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Solution 1: %d\n", solve1(text))
		fmt.Printf("Solution 2: %d\n", solve2(text))
	}
}
