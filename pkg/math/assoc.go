package math

func Max(val int, vals ...int) int {
	max := val
	for _, v := range vals {
		if max <= v {
			max = v
		}
	}
	return max
}

func Min(val int, vals ...int) int {
	min := val
	for _, v := range vals {
		if min >= v {
			min = v
		}
	}
	return min
}
