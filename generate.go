package main

import (
	"fmt"
	"strings"
)

func Generate(input string, bannerFile string) (string, error) {

	lines, err := LoadBanner(bannerFile)
	if err != nil {
		return "", fmt.Errorf("error reading banner: %w", err)
	}

	parts := Split(input)

	var sb strings.Builder

	for _, part := range parts {

		if part == "" {
			sb.WriteString("\n")
			continue
		}

		sb.WriteString(Render(part, lines))
	}

	return sb.String(), nil
}