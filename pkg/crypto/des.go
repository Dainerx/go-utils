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

func permute(input []byte, positions []int8, n int) []byte {
	permutedInput := make([]byte, n)
	for i := 0; i < n; i++ {
		permutedInput[i] = input[positions[i]]
	}
	return permutedInput
}

func splitKey(input []byte, n int) ([]byte, []byte) {
	left := make([]byte, n/2)
	right := make([]byte, n/2)
	copy(left[:], input[:n/2])
	copy(right[:], input[n/2:])
	return left, right
}

func circularShift(input []byte, rotation int, n int) []byte {
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

func xor(input []byte, key []byte, n int) []byte {
	xorResult := make([]byte, n)
	for i := 0; i < n; i++ {
		xorResult[i] = input[i] ^ key[i]
	}
	return xorResult
}

func getTwoBytesFromS0(row, column int8) []byte {
	twoBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(twoBytes, S0[row][column])
	return twoBytes
}

func getTwoBytesFromS1(row, column int8) []byte {
	twoBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(twoBytes, S1[row][column])
	return twoBytes
}

func byteToInt(sliceByte []byte) int {
	result := 0
	for i, bit := range sliceByte {
		sum := 0
		if bit == 1 {
			sum = 1
			c := i
			for c > 0 {
				sum = sum << 1
				c--
			}
		}
		result += sum
	}

	return result
}

func EncryptEightBitsText(plainText []byte, k []byte) []byte {
	// Input: plain text and a key.
	// 8 bits plain text.
	fmt.Printf("Plain text: %b\n", plainText)
	// 10 bits key.
	fmt.Printf("Key:%b\n\n", k)

	fmt.Println("Generating K1 and K2...")
	// K1 Generation
	seedK1 := k
	// Generate kp10 = P10(k)
	var kp10 = permute(seedK1, []int8{2, 4, 1, 6, 3, 9, 0, 8, 7, 5}, 10)
	fmt.Printf("P10(k) = %b\n", kp10)
	// Split kp10
	kp10Left, kp10Right := splitKey(kp10, 10)
	// Shift both left and right
	kp10Left, kp10Right = circularShift(kp10Left, -1, 5), circularShift(kp10Right, -1, 5)
	// Concat both to generate seedK2 and move to the final step for K1.
	seedK2 := append(kp10Left, kp10Right...)
	fmt.Printf("Shift(P10(k)) = %b\n", seedK2)
	// Generate k1 through final permutation P8
	k1 := permute(seedK2, []int8{5, 2, 6, 3, 7, 4, 9, 8}, 8)
	fmt.Printf("K1 = P8(Shift(P10(k))) = %b\n", k1)

	// K2 Generation
	// Split the seed
	seedK2Left, seedK2Right := splitKey(seedK2, 10)
	// Shift the seed
	seedK2Shifted := append(circularShift(seedK2Left, -2, 5),
		circularShift(seedK2Right, -2, 5)...)
	// Generate k2 through one final permutation P8
	k2 := permute(seedK2Shifted, []int8{5, 2, 6, 3, 7, 4, 9, 8}, 8)
	fmt.Printf("K2 = P8(Shift(Shift(P10(k)))) = %b\n\n", k2)

	// Run IP on the plain text
	plainTextIP := permute(plainText, []int8{1, 5, 2, 0, 3, 7, 4, 6}, 8)
	fmt.Printf("IP(text) = %b\n", plainTextIP)

	// Generate L and R that are the input for the fk function
	l, r := splitKey(plainTextIP, 8)
	fmt.Println("Running fk on K1:")
	// We will use R as input for expantion and permutation
	r1 := permute(r, []int8{3, 0, 1, 2, 1, 2, 3, 0}, 8)
	fmt.Printf("R1 = E/P(R) = %b\n", r1)
	r1XorK1 := xor(r1, k1, 8)
	fmt.Printf("R1 xor K1 = %b\n", r1XorK1)
	// S boxes section
	x1, y1, x2, y2 := []byte{r1XorK1[3], r1XorK1[0]}, []byte{r1XorK1[2], r1XorK1[1]},
		[]byte{r1XorK1[7], r1XorK1[4]}, []byte{r1XorK1[6], r1XorK1[5]}
	rowS0 := byteToInt(x1)
	columnS0 := byteToInt(y1)
	rowS1 := byteToInt(x2)
	columnS1 := byteToInt(y2)
	// Fetching results from the S boxes
	resultS0 := getTwoBytesFromS0(int8(rowS0), int8(columnS0)) // 2 bits
	resultS1 := getTwoBytesFromS1(int8(rowS1), int8(columnS1)) // 2 bits
	fmt.Printf("Result 2 bits from S0/S1 = %b / %b\n", resultS0, resultS1)
	outputFRK1 := permute(append(resultS0, resultS1...),
		[]int8{1, 3, 2, 0}, 4)
	fmt.Printf("F(R,K1) = %b\n\n", outputFRK1)

	fmt.Println("Running fk on K2:")
	// We will use R as input for expantion and permutation
	l1 := permute(l, []int8{3, 0, 1, 2, 1, 2, 3, 0}, 8)
	fmt.Printf("L1 = E/P(L) = %b\n", l1)
	l1XorK2 := xor(l1, k2, 8)
	fmt.Printf("L1 xor K2 = %b\n", l1XorK2)
	// S boxes section
	lx1, ly1, lx2, ly2 := []byte{l1XorK2[3], l1XorK2[0]}, []byte{l1XorK2[2], l1XorK2[1]},
		[]byte{l1XorK2[7], l1XorK2[4]}, []byte{l1XorK2[6], l1XorK2[5]}
	lrowS0 := byteToInt(lx1)
	lcolumnS0 := byteToInt(ly1)
	lrowS1 := byteToInt(lx2)
	lcolumnS1 := byteToInt(ly2)
	fmt.Printf("%d,%d  %d,%d\n\n", lrowS0, lcolumnS0, lrowS1, lcolumnS1)
	// Fetching results from the S boxes
	lresultS0 := getTwoBytesFromS0(int8(lrowS0), int8(lcolumnS0)) // 2 bits
	lresultS1 := getTwoBytesFromS1(int8(lrowS1), int8(lcolumnS1)) // 2 bits
	fmt.Printf("Result 2 bits from S0/S1 = %b / %b\n", lresultS0, lresultS1)
	outputFLK2 := permute(append(lresultS0, lresultS1...),
		[]int8{1, 3, 2, 0}, 4)
	fmt.Printf("F(L,K2) = %b\n", outputFLK2)

	// sum up the two outputs and apply IP-1
	output := append(outputFLK2, outputFRK1...)
	finalOutput := permute(output, []int8{3, 0, 2, 4, 6, 1, 7, 5}, 8)
	return finalOutput
}
