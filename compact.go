package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	var compacted []string

	for _, v := range slice {
		if v != "" {
			compacted = append(compacted, v)
		}
	}

	*ptr = compacted
	return len(compacted)
}
