package common

import (
	"bufio"
	"os"
)

func LoadFileLines(path string) ([]string, error) {

	lines := make([]string, 0)

	f, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
