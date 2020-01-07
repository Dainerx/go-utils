package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountTableLoop(x uint64) int {
	popcount := 0
	for i := 0; i < 8; i++ {
		popcount += int(pc[byte(x>>(i*8))])
	}
	return popcount
}

func PopCountShiftValue(x uint64) int {
	popcount := 0
	for i := 0; i < 64; i++ {
		if x%2 == 1 {
			popcount++
		}
		x >>= 1
	}
	return popcount
}

func PopCountClearRightmost(x uint64) int {
	i := 0
	for i = 0; x != 0; i++ {
		x &= x - 1
	}
	return i
}
