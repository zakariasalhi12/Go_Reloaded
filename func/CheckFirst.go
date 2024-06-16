package reroaded

func FirstElementChecker(s []string) []string {
	for len(s) > 0 && (s[0] == "(hex)" || s[0] == "(bin)" || s[0] == "(cap)" || s[0] == "(up)" || s[0] == "(low)") {
		PrintWarning("Cannot use a flag as the first element")
		s = s[1:]
	}

	for len(s) >= 2 && (s[0] == "(cap," || s[0] == "(up," || s[0] == "(low,") {
		if s[1][len(s[1])-1] == ')' {
		PrintWarning("Cannot use a flag as the first element")
			s = s[2:]
		} else {
			break
		}
	}
	return s
}
