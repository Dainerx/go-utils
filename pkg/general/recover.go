package general

func NoReturn() {
	panic("return")
}

func ChangeReturnSquare(n int, changeReturn bool) (x int) {
	defer func() {
		switch changeReturn {
		case true:
			//return
			x = -1
		case false:
			// no return
		default:
			panic("corrupted vlaue of changeReturn")
		}
	}()

	x = n * n
	return x
}
