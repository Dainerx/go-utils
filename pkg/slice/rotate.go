package slice

func rotate(x []int, r int) {
	// If r is negative means left rotating
	// Then rotate on the right by len(x) + r
	if r < 0 {
		r = len(x) + (r % len(x))
	}
	y := make([]int, len(x))
	copy(y, x)
	for i := range y {
		x[(i+r)%len(x)] = y[i]
	}
}
