package reroaded

func NumberFlag(s string) bool {
	numers := "-0123456789)"
	Closed := false
	for _, char := range s {

		if !contains(numers, char) {
			PrintWarning("The flag should receive a number")
			return false
		}

		if char == '-' {
			PrintWarning("The flag should receive a Positive number")
		}

		if char == ')' {
			Closed = true
		}

		if Closed {
			return true
		}
	}
	PrintWarning("The flag should receive a number")
	return false
}
