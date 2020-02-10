package crypto

import "fmt"

func Permute(input []byte, positions []int8, n int) []byte {
	permutedInput := make([]byte, n)
	for i := 0; i < n; i++ {
		permutedInput[i] = input[positions[i]]
	}
	return permutedInput
}

func SplitKey(input []byte, n int) ([]byte, []byte) {
	left := make([]byte, n/2)
	right := make([]byte, n/2)
	copy(left[:], input[:n/2])
	copy(right[:], input[n/2:])
	return left, right
}

func CircularShift(input []byte, rotation int, n int) []byte {
	rotatedInput := make([]byte, n)
	rotation = rotation % n
	if rotation != 0 { // needs to rotate
		if rotation < 0 { // negative value just add size to solve.
			rotation += n
		}
		fmt.Println(rotation)
		for i := 0; i < n; i++ {
			newPos := (i + rotation) % (n)
			rotatedInput[newPos] = input[i]
		}
	}
	return rotatedInput
}

func Xor(input []byte, key []byte, n int) []byte {
	xorResult := make([]byte, n)
	for i := 0; i < n; i++ {
		xorResult[i] = input[i] ^ key[i]
	}
	return xorResult
}
