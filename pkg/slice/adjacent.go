package slice

func elimnateAdjDuplicate(s []string) []string {
	n := len(s)
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:]) //costy
			n--
		}
	}
	return s[:n]
}

func elimnateAdjDuplicate1(s []string) []string {
	c := 0
	for _, str := range s {
		if s[c] == str {
			continue
		}
		c++
		s[c] = str
	}
	return s[:c+1]
}
