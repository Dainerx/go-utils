package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Dainerx/go-utils/pkg/crypto"
)

func main() {
	// Input: plain text and a key.
	// 8 bits plain text.
	var plainText = []byte{1, 0, 0, 1, 0, 0, 0, 1}
	fmt.Printf("Plain text: %b\n", plainText)
	// 10 bits key.
	var k = []byte{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}
	fmt.Printf("Key:%b\n\n", k)

	fmt.Println("Generating K1 and K2...")
	// K1 Generation
	seedK1 := k
	// Generate kp10 = P10(k)
	var kp10 = crypto.Permute(seedK1, []int8{2, 4, 1, 6, 3, 9, 0, 8, 7, 5}, 10)
	fmt.Printf("P10(k) = %b\n", kp10)
	// Split kp10
	kp10Left, kp10Right := crypto.SplitKey(kp10, 10)
	// Shift both left and right
	kp10Left, kp10Right = crypto.CircularShift(kp10Left, -1, 5), crypto.CircularShift(kp10Right, -1, 5)
	// Concat both to generate seedK2 and move to the final step for K1.
	seedK2 := append(kp10Left, kp10Right...)
	fmt.Printf("Shift(P10(k)) = %b\n", seedK2)
	// Generate k1 through final permutation P8
	k1 := crypto.Permute(seedK2, []int8{5, 2, 6, 3, 7, 4, 9, 8}, 8)
	fmt.Printf("K1 = P8(Shift(P10(k))) = %b\n", k1)

	// K2 Generation
	// Split the seed
	seedK2Left, seedK2Right := crypto.SplitKey(seedK2, 10)
	// Shift the seed
	seedK2Shifted := append(crypto.CircularShift(seedK2Left, -2, 5),
		crypto.CircularShift(seedK2Right, -2, 5)...)
	// Generate k2 through one final permutation P8
	k2 := crypto.Permute(seedK2Shifted, []int8{5, 2, 6, 3, 7, 4, 9, 8}, 8)
	fmt.Printf("K2 = P8(Shift(Shift(P10(k)))) = %b\n\n", k2)

	// Run IP on the plain text
	plainTextIP := crypto.Permute(plainText, []int8{1, 5, 2, 0, 3, 7, 4, 6}, 8)
	fmt.Printf("IP(text) = %b\n", plainTextIP)

	// Generate L and R that are the input for the fk function
	_, r := crypto.SplitKey(plainTextIP, 8)
	fmt.Println("Running fk on K1:")
	// We will use R as input for expantion and permutation
	r1 := crypto.Permute(r, []int8{3, 0, 1, 2, 1, 2, 3, 0}, 8)
	fmt.Printf("R1 = E/P(R) = %b\n", r1)
	r1XorK1 := crypto.Xor(r1, k1, 8)
	fmt.Printf("R1 xor K1 = %b\n", r1XorK1)
	// S boxes bugging section
	//x, y := []byte{r1XorK1[3], r1XorK1[0]}, []byte{r1XorK1[2], r1XorK1[1]}
	//fmt.Printf("x,b: %b,%b\n", x, y)
	rowS0, _ := binary.ReadVarint(bytes.NewBuffer([]byte{r1XorK1[3], r1XorK1[0]}))
	columnS0, _ := binary.ReadVarint(bytes.NewBuffer([]byte{r1XorK1[2], r1XorK1[1]}))
	rowS1, _ := binary.ReadVarint(bytes.NewBuffer([]byte{r1XorK1[7], r1XorK1[4]}))
	columnS1, _ := binary.ReadVarint(bytes.NewBuffer([]byte{r1XorK1[6], r1XorK1[5]}))
	// Fetching results from the S boxes
	resultS0 := crypto.GetTwoBytesFromS0(int8(rowS0), int8(columnS0)) // 2 bits
	resultS1 := crypto.GetTwoBytesFromS1(int8(rowS1), int8(columnS1)) // 2 bits
	fmt.Printf("Result 2 bits from S0/S1 = %b / %b\n", resultS0, resultS1)
	outputFRK1 := crypto.Permute(append(resultS0, resultS1...),
		[]int8{1, 3, 2, 0}, 4)
	fmt.Printf("F(R,K1) = %b\n", outputFRK1)
}
