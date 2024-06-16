package reroaded

import (
	"strconv"
	"strings"
)

func Reloaded(s string) string {
	hundlesplit := strings.Split(s, " ")
	word2 := HundleA(hundlesplit)
	word2 = strings.ReplaceAll(word2, "\n", " \n ")
	splited := strings.Split(word2, " ")
	splited = FirstElementChecker(splited)
	splited = RemoveEmptyElements(splited)
	for index := 0; index < len(splited); index++ {
		// hex
		if splited[index] == "(hex)" {
			if HexCheker(strings.ToLower(splited[index-1])) {
				splited[index-1] = PrintNbrBase(AtoiBase(string(strings.ToLower(splited[index-1])), "0123456789abcdef"), "0123456789")
			}
			splited = append(splited[:index], splited[index+1:]...)
			index--
			// bin
		} else if splited[index] == "(bin)" {
			if BinCheker(strings.ToLower(splited[index-1])) {
				splited[index-1] = PrintNbrBase(AtoiBase(splited[index-1], "01"), "0123456789")
			}
			splited = append(splited[:index], splited[index+1:]...)
			index--
			// up
		} else if splited[index] == "(up)" {
			splited[index-1] = strings.ToUpper(splited[index-1])
			splited = append(splited[:index], splited[index+1:]...)
			index--
			// low
		} else if splited[index] == "(low)" {
			splited[index-1] = strings.ToLower(splited[index-1])
			splited = append(splited[:index], splited[index+1:]...)
			index--
			// cap
		} else if splited[index] == "(cap)" {
			splited[index-1] = Capitalize(splited[index-1])
			splited = append(splited[:index], splited[index+1:]...)
			index--
			// cap with number
		} else if splited[index] == "(up," && len(splited) > 2 {
			if NumberFlag(splited[index+1]) {
				fullnumber := splited[index+1]
				number, _ := strconv.Atoi(fullnumber[:len(fullnumber)-1])
				for j := 1; j <= number; j++ {
					splited[index-j] = strings.ToUpper(splited[index-j])
					if index-j == 0 {
						break
					}
				}
				splited = append(splited[:index], splited[index+2:]...)
				index--
			}

			// low with number
		} else if splited[index] == "(low," && len(splited) > 2 {
			if NumberFlag(splited[index+1]) {
				fullnumber := splited[index+1]
				number, _ := strconv.Atoi(fullnumber[:len(fullnumber)-1])
				for j := 1; j <= number; j++ {
					splited[index-j] = strings.ToLower(splited[index-j])
					if index-j == 0 {
						break
					}
				}
				splited = append(splited[:index], splited[index+2:]...)
				index--
			}
			// cap with number
		} else if splited[index] == "(cap," && len(splited) > 2 {
			if NumberFlag(splited[index+1]) {
				fullnumber := splited[index+1]
				number, _ := strconv.Atoi(fullnumber[:len(fullnumber)-1])
				for j := 1; j <= number; j++ {
					splited[index-j] = Capitalize(splited[index-j])
					if index-j == 0 {
						break
					}
				}
				splited = append(splited[:index], splited[index+2:]...)
				index--
			}
		}
	}

	b := Puncs(splited)

	return HundleA(b)
}
