package reroaded

import (
	"strings"
)

func HundleA(s []string) string {
	res := ""

	for index, word := range s {
		// hundle a e i o ....
		if index >= 0 && word == "a" && index != len(s)-1 {
			nextWord := strings.ToLower(s[index+1])
			if strings.HasPrefix(strings.ToLower(nextWord), "a") || strings.HasPrefix(strings.ToLower(nextWord), "e") || strings.HasPrefix(strings.ToLower(nextWord), "i") || strings.HasPrefix(strings.ToLower(nextWord), "o") || strings.HasPrefix(strings.ToLower(nextWord), "u") || strings.ToLower(nextWord) == "h" {
				res += "an "
			} else {
				res += word + " "
			}
		} else if index >= 0 && word == "A" && index != len(s)-1 {
			nextWord := strings.ToLower(s[index+1])
			if strings.HasPrefix(strings.ToLower(nextWord), "a") || strings.HasPrefix(strings.ToLower(nextWord), "e") || strings.HasPrefix(strings.ToLower(nextWord), "i") || strings.HasPrefix(strings.ToLower(nextWord), "o") || strings.HasPrefix(strings.ToLower(nextWord), "u") || strings.ToLower(nextWord) == "h" {
				res += "An "
			} else {
				res += word + " "
			}
		} else if index >= 0 && word == "'A" && index != len(s)-1 {
			nextWord := strings.ToLower(s[index+1])
			if strings.HasPrefix(strings.ToLower(nextWord), "a") || strings.HasPrefix(strings.ToLower(nextWord), "e") || strings.HasPrefix(strings.ToLower(nextWord), "i") || strings.HasPrefix(strings.ToLower(nextWord), "o") || strings.HasPrefix(strings.ToLower(nextWord), "u") || strings.ToLower(nextWord) == "h" {
				res += "'An "
			} else {
				res += word + " "
			}
		} else if index >= 0 && word == "'a" && index != len(s)-1 {
			nextWord := strings.ToLower(s[index+1])
			if strings.HasPrefix(strings.ToLower(nextWord), "a") || strings.HasPrefix(strings.ToLower(nextWord), "e") || strings.HasPrefix(strings.ToLower(nextWord), "i") || strings.HasPrefix(strings.ToLower(nextWord), "o") || strings.HasPrefix(strings.ToLower(nextWord), "u") || strings.ToLower(nextWord) == "h" {
				res += "'an "
			} else {
				res += word + " "
			}
		} else {
			res += word + " "
		}
	}

	return strings.TrimSpace(strings.ReplaceAll(res, " \n ", "\n"))
}
