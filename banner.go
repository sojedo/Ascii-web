package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) ([]string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("empty banner file")
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n")


	if len(lines) < 855 {
		return nil, fmt.Errorf("invalid banner file format")
	}

	return lines, nil
}
