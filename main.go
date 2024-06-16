package main

import (
	"fmt"
	"os"
	"strings"

	reroaded "reroaded/func"
)

func main() {
	if len(os.Args) != 3 {
		reroaded.PrintEror("You Need 2 args to use this tool")
		reroaded.PrintEX("sample.txt resa.txt")
		return
	}

	args := os.Args[1:]

	text, _ := os.ReadFile(args[0])

	// Hundle file error

	if len(args[0]) >= 4 && strings.ToLower(args[0][len(args[0])-4:len(args[0])]) != ".txt" {
		reroaded.PrintWarning("You Can Only Edit on txt files")
		reroaded.PrintEX("sample.txt resa.txt")
		return
	}

	if len(args[0]) >= 4 && strings.ToLower(args[1][len(args[1])-4:len(args[1])]) != ".txt" {
		reroaded.PrintWarning("You Can Only Edit on txt files")
		reroaded.PrintEX("sample.txt resa.txt")
		return
	}

	resault := reroaded.Reloaded(string(text))

	finalres := os.WriteFile(args[1], []byte(resault), 0o777)
	if finalres != nil {
		fmt.Printf("Error :%s", finalres)
		fmt.Println()
		return
	}
}
