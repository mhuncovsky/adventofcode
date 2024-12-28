package day12

import (
	"testing"
)

func TestO(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".A.A.")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".....")
	expected := 8
	sides := countSides(lines, V{1, 1})
	if sides != expected {
		t.Fatalf(`sides=%v, expected %v`, sides, expected)
	}
}

func TestTL(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, "..AA.")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".....")
	expected := 6
	sides := countSides(lines, V{2, 1})
	if sides != expected {
		t.Fatalf(`sides=%v, expected %v`, sides, expected)
	}
}

func TestTLTR(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, "..A..")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".....")
	expected := 8
	sides := countSides(lines, V{2, 1})
	if sides != expected {
		t.Fatalf(`sides=%v, expected %v`, sides, expected)
	}
	expected += 4
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestPlus(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, "..A..")
	lines = append(lines, ".AAA.")
	lines = append(lines, "..A..")
	lines = append(lines, ".....")
	expected := 12
	sides := countSides(lines, V{2, 1})
	if sides != expected {
		t.Fatalf(`A sides=%v, expected %v`, sides, expected)
	}
	expected = 16
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestTLL(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".A...")
	lines = append(lines, ".....")
	lines = append(lines, ".....")
	expected := 6
	sides := countSides(lines, V{2, 1})
	if sides != expected {
		t.Fatalf(`sides=%v, expected %v`, sides, expected)
	}
	expected += 4
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestTRL(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, ".AAA.")
	lines = append(lines, "...A.")
	lines = append(lines, ".....")
	lines = append(lines, ".....")
	expected := 6
	sides := countSides(lines, V{2, 1})
	if sides != expected {
		t.Fatalf(`sides=%v, expected %v`, sides, expected)
	}
	expected += 4
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestBLL(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, ".A...")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".....")
	lines = append(lines, ".....")
	expected := 6
	sides := countSides(lines, V{1, 1})
	if sides != expected {
		t.Fatalf(`sides=%v, expected %v`, sides, expected)
	}
	expected += 4
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestBRL(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".....")
	lines = append(lines, "...A.")
	lines = append(lines, ".AAA.")
	lines = append(lines, ".....")
	lines = append(lines, ".....")
	expected := 6
	sides := countSides(lines, V{3, 1})
	if sides != expected {
		t.Fatalf(`sides=%v, expected %v`, sides, expected)
	}
	expected += 4
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestZigZag(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, "........")
	lines = append(lines, "...AAA..")
	lines = append(lines, ".AAA.A..")
	lines = append(lines, ".....AA.")
	lines = append(lines, "........")
	expected := 12
	sides := countSides(lines, V{3, 1})
	if sides != expected {
		t.Fatalf(`A sides=%v, expected %v`, sides, expected)
	}
	expected += 4
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestExtract(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, ".................")
	lines = append(lines, "........L........")
	lines = append(lines, "........L........")
	lines = append(lines, ".......LLLLL.....")
	lines = append(lines, "....LLLLLLL......")
	lines = append(lines, "...LLLLLLL.......")
	lines = append(lines, "..LLLLLLLLLL.....")
	lines = append(lines, "..LLLLLLLLLLLLLL.")
	lines = append(lines, ".LLLLLLLLLLLLL...")
	lines = append(lines, ".LLLLLLLLLLLLL...")
	lines = append(lines, ".LLLLLLLLLLLLL...")
	lines = append(lines, ".LLLLLLLLLLLLL...")
	lines = append(lines, "....LLLLLL.......")
	lines = append(lines, "....LLLLLL.......")
	lines = append(lines, ".....LL.LLL......")
	lines = append(lines, ".....L...........")
	lines = append(lines, ".................")
	expected := 40
	sides := countSides(lines, V{8, 1})
	if sides != expected {
		t.Fatalf(`L sides=%v, expected %v`, sides, expected)
	}
	expected += 4
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

func TestExtract2(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, "............")
	lines = append(lines, "...K...KKK..")
	lines = append(lines, "...KK..KKKK.")
	lines = append(lines, "..KKKK.KKKK.")
	lines = append(lines, "..KKKKKKKKK.")
	lines = append(lines, "..KKKKKKKKK.")
	lines = append(lines, "..KKKKKKKKK.")
	lines = append(lines, ".KKKKKKKKKK.")
	lines = append(lines, ".KKKKKKKKKK.")
	lines = append(lines, ".KKKKKKKKKK.")
	lines = append(lines, ".KKKKKKKKKK.")
	lines = append(lines, "..KKKKKKKKK.")
	lines = append(lines, "..KKKKKKKKK.")
	lines = append(lines, "..KKKK.K.K..")
	lines = append(lines, "...KKK.K.K..")
	lines = append(lines, "...KK.KK....")
	lines = append(lines, "...K........")
	lines = append(lines, "............")
	expected := 38
	sides := countSides(lines, V{3, 1})
	if sides != expected {
		t.Fatalf(`K sides=%v, expected %v`, sides, expected)
	}
	// expected += 4 // the 4 sides inside are not counted
	sides = countSides(lines, V{0, 0})
	if sides != expected {
		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
	}
}

// func TestExtract2Use3(t *testing.T) {
// 	lines := make([]string, 0)
// 	lines = append(lines, "............")
// 	lines = append(lines, "...K...KKK..")
// 	lines = append(lines, "...KK..KKKK.")
// 	lines = append(lines, "..KKKK.KKKK.")
// 	lines = append(lines, "..KKKKKKKKK.")
// 	lines = append(lines, "..KKKKKKKKK.")
// 	lines = append(lines, "..KKKKKKKKK.")
// 	lines = append(lines, ".KKKKKKKKKK.")
// 	lines = append(lines, ".KKKKKKKKKK.")
// 	lines = append(lines, ".KKKKKKKKKK.")
// 	lines = append(lines, ".KKKKKKKKKK.")
// 	lines = append(lines, "..KKKKKKKKK.")
// 	lines = append(lines, "..KKKKKKKKK.")
// 	lines = append(lines, "..KKKK.K.K..")
// 	lines = append(lines, "...KKK.K.K..")
// 	lines = append(lines, "...KK.KK....")
// 	lines = append(lines, "...K........")
// 	lines = append(lines, "............")
// 	expected := 38
// 	SHOW = true
// 	sides := countSides3(lines, V{3, 1})
// 	SHOW = false
// 	if sides != expected {
// 		t.Fatalf(`K sides=%v, expected %v`, sides, expected)
// 	}
// 	// expected += 4 // the 4 sides inside are not counted
// 	sides = countSides(lines, V{0, 0})
// 	if sides != expected {
// 		t.Fatalf(`. sides=%v, expected %v`, sides, expected)
// 	}
// }
