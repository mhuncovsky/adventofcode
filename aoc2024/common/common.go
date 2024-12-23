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
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

func LoadFileText(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
