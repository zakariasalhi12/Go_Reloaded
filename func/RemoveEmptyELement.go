package reroaded

func RemoveEmptyElements(s []string) []string {

	new := []string{}

	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			new = append(new, s[i])
		}
	}

	return new
}
