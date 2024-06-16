package reroaded

func BinCheker(s string) bool {
	bin := "-01"

	if s == "" {
		return false
	}

	for i, char := range s {

		if i != 0 && char == '-' {
			PrintWarning(s + " is not a bin Base")
			PrintEX("10 (bin)")
			return false
		}

		if !contains(bin, char) {
			PrintWarning(s + " is not a bin Base")
			PrintEX("10 (bin)")
			return false
		}
	}
	return true
}

func HexCheker(s string) bool {
	hex := "-0123456789abcdef"

	if s == "" {
		return false
	}

	for i, char := range s {
		
		if i != 0 && char == '-' {
			PrintWarning(s + " is not a hex Base")
			PrintEX("10 (bin)")
			return false
		}

		if !contains(hex, char) {
			PrintWarning(s + " is not a hex Base")
			PrintEX("7D (hex)")
			return false
		}
	}
	return true
}

func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}
