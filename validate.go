package main

import "strings"

func Validate(args []string) (string, string, bool) {

	if len(args) != 1 && len(args) != 2 {
		return "", "", false
	}

	input := args[0]
	bannerFile := "banners/standard.txt"
if len(args) == 2 {
    switch args[1] {
    case "standard", "shadow", "thinkertoy":
        bannerFile = "banners/" + args[1] + ".txt"
    default:
        return "", "", false
    }
}

	clean := strings.ReplaceAll(input, `\n`, "")

	for _, ch := range clean {
		if ch < 32 || ch > 126 {
			return "", "", false
		}
	}

	return input, bannerFile, true
}