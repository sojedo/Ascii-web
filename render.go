package main

import "strings"

func Render(part string, lines []string) string {
	var sb strings.Builder

	for row := 0; row < 8; row++ {

		for _, ch := range part {

			index := (int(ch)-32)*9 + 1

			sb.WriteString(lines[index+row])
		}

		sb.WriteString("\n")
	}

	return sb.String()
}