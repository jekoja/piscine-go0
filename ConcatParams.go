package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// calculate total length for make()
	totalLen := 0
	for _, s := range args {
		totalLen += len(s)
	}
	totalLen += len(args) - 1 // for '\n' between strings

	// create a byte slice
	result := make([]byte, 0, totalLen)

	// concatenate with newline
	for i, s := range args {
		for _, ch := range s {
			result = append(result, byte(ch))
		}
		if i < len(args)-1 {
			result = append(result, '\n')
		}
	}

	return string(result)
}
