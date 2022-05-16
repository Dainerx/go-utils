package strings

func JoinVariant(sep string, s ...string) string {
	var result string
	for i, ss := range s {
		if i == len(s)-1 {
			result += ss
		} else {
			result += ss + sep
		}
	}

	return result
}
