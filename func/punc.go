package reroaded

func Puncs(s []string) []string {
	fawasil := []string{".", ",", "!", "?", ":", ";"}
	s = RemoveEmptyElements(s)

	if checker(s) {
		return s
	}

	Mark := true

	for i := 0; i < len(s); i++ {
		if len(s[i]) == 0 {
			continue // Skip empty strings
		}

		if len(s[i]) > 1 && s[i][0] == '\'' {
			Mark = false
		} else if len(s[i]) > 1 && s[i][len(s[i])-1] == '\'' {
			Mark = true
		} else if len(s[i]) > 1 && s[i][0] == '\'' && s[i][len(s[i])-1] == '\'' {
			Mark = true
		}

		if Mark && s[i] == "'" && i >= 0 && i < len(s)-1 {
			s[i+1] = "'" + s[i+1]
			s[i] = ""
			Mark = false
		} else if !Mark && s[i] == "'" && i >= 0 && i > 0 {
			s[i-1] = s[i-1] + "'"
			s[i] = ""
			Mark = true
		}
	}

	// Additional handling for the first and last elements of s
	if len(s) > 0 && s[0] == "'" {
		if len(s) > 1 {
			s[1] = "'" + s[1]
		}
		s = s[1:]
	}

	if len(s) > 0 && s[len(s)-1] == "'" {
		if len(s) > 1 {
			s[len(s)-2] += "'"
		}
		s = s[:len(s)-1]
	}

	for i := 1; i < len(s); i++ {
		for j := 0; j < len(fawasil); j++ {
			if len(s[i]) > 0 && string(s[i][0]) == fawasil[j] {
				// Append the punctuation to the previous word
				if i > 0 {
					s[i-1] += string(s[i][0])
				}
				// Remove the punctuation from the current word
				if len(s[i]) > 1 {
					s[i] = s[i][1:]
					break
				} else {
					s = append(s[:i], s[i+1:]...)
					break
				}
			}
		}
	}

	// Handle single quote mark in the middle

	// Recursively process the slice until all punctuation marks are properly placed
	return Puncs(s)
}

func checker(s []string) bool {
	// Check for the single quote mark
	for i, word := range s {
		if i > 0 && word == "'" {
			return false
		}
	}

	fawasil := []string{".", ",", "!", "?", ":", ";"}

	// Check for punctuation marks
	for i, word := range s {
		for _, punc := range fawasil {
			if i > 0 && len(word) > 0 && string(word[0]) == punc {
				return false
			}
		}
	}

	return true
}
