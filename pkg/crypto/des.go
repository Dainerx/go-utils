package crypto

import (
	"encoding/binary"
	"fmt"
)

var S0 = [][]uint16{
	{1, 0, 3, 2},
	{3, 2, 1, 0},
	{0, 2, 1, 3},
	{3, 1, 3, 2},
}
var S1 = [][]uint16{
	{0, 1, 2, 3},
	{2, 0, 1, 3},
	{3, 0, 1, 0},
	{2, 1, 0, 3},
}

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

func GetTwoBytesFromS0(row, column int8) []byte {
	twoBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(twoBytes, S0[row][column])
	return twoBytes
}

func GetTwoBytesFromS1(row, column int8) []byte {
	fmt.Printf("%d,%d\n", row, column)
	twoBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(twoBytes, S1[row][column])
	return twoBytes
}
