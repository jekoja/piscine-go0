package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert from baseFrom to decimal
	baseFromLen := len(baseFrom)
	value := 0

	for _, digit := range nbr {
		value = value*baseFromLen + indexInBase(baseFrom, byte(digit))
	}

	// Step 2: Convert from decimal to baseTo
	if value == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	var result []byte

	for value > 0 {
		remainder := value % baseToLen
		result = append([]byte{baseTo[remainder]}, result...)
		value /= baseToLen
	}

	return string(result)
}

// Helper function to get index of a character in base string
func indexInBase(base string, c byte) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1 // shouldn't happen with valid input
}
