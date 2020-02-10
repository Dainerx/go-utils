package main

import (
	"fmt"

	"github.com/Dainerx/go-utils/pkg/crypto"
)

func main() {
	var k = [10]byte{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}
	fmt.Printf("k:%b", k)

	var kp10 [10]byte
	copy(kp10[:], crypto.Permute(k, []int8{2, 4, 1, 6, 3, 9, 0, 8, 7, 5}, 10)[:])
	fmt.Printf("\nkp10:%b", kp10)

	var kp8 [8]byte
	copy(kp8[:], crypto.Permute(kp10, []int8{5, 2, 6, 3, 7, 4, 9, 8}, 8))
	fmt.Printf("\nk1:%b", kp8)

	// Split kp10 to get the left and right array of bytes
	left, right := crypto.SplitKey(kp10)
	fmt.Printf("\nleft:%b,right:%b", left, right)
	// Apply the shift on both left and right
	left, right = crypto.CircularShift(left, -3, 5), crypto.CircularShift(right, -3, 5)
	fmt.Printf("\nleft:%b,right:%b", left, right)
	// Concat left and right to get k2
	var k2 [10]byte
	copy(k2[:5], left[:])
	copy(k2[5:], right[:])
	fmt.Printf("\nk2:%b", k2)
	// Permutation 8 for k2
	var k2p8 [8]byte
	copy(k2p8[:], crypto.Permute(k2, []int8{5, 2, 6, 3, 7, 4, 9, 8}, 8)[:8])
	fmt.Printf("\nk2p8:%b", k2p8)
}
