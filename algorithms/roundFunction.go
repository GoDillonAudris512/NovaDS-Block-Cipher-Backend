package algorithms

import (
	"block-cipher/constants"
)

func roundFunction(input []int, roundKey []int) []int {
	//****************** PERMUTATION 1 *******************//
	// Expansion permutation from 64-bit input to 96-bit input
	// This permutation operates in bytes

	permutation1Result := permutation1(input)

	//****************** SUBSTITUTION 1 *******************//
	substitution1Result := substitution1(permutation1Result, roundKey)

	//****************** PERMUTATION 2 *******************//
	permutation2Result := permutation2(substitution1Result, roundKey)

	//****************** SUBSTITUTION 2 *******************//
	substitution2Result := substitution2(permutation2Result)

	//****************** PERMUTATION 3 *******************//
	// Compression permutation from 96-bit input to 64-bit input
	// This permutation operates in bits
	permutation3Result := permutation3(substitution2Result)

	return permutation3Result
}

func generateSBox(key []int) [][][]int {
	// generateSBox generates the S-Box using the provided key.
	// It constructs the S-Box by filling it with values derived from the key
	// using cyclic shifts and index manipulation.

	sBox := make([][]int, 64)
	for i := range sBox {
		sBox[i] = make([]int, 6)
	}

	// Populate the first 21 rows of the S-Box
	// with values derived from the key
	index := 0
	for i := 0; i < 21; i++ {
		for j := 0; j < 6; j++ {
			sBox[i][j] = key[index]
			index++
		}
	}

	// Continue populating the S-Box using cyclic shifts
	// for the next 21 and 21 rows respectively
	key = CyclicShiftLeft(key, 7)
	index = 0
	for i := 21; i < 42; i++ {
		for j := 0; j < 6; j++ {
			sBox[i][j] = key[index]
			index++
		}
	}

	// Continue populating the S-Box using cyclic shifts
	// for the next 21 and 21 rows respectively
	key = CyclicShiftLeft(key, 7)
	index = 0
	for i := 42; i < 63; i++ {
		for j := 0; j < 6; j++ {
			sBox[i][j] = key[index]
			index++
		}
	}

	// Fill the last row of the S-Box with the remaining key bits
	key = CyclicShiftLeft(key, 7)
	for j := 0; j < 6; j++ {
		sBox[63][j] = key[j]
	}

	// Rearrange the S-Box into a 4x16x6 structure
	fsBox := make([][][]int, 4)
	for i := range fsBox {
		fsBox[i] = make([][]int, 16)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 16; j++ {
			fsBox[i][j] = sBox[i*16+j]
		}
	}

	return fsBox
}

func permutation1(input []int) []int {
	// Initialize an empty variable to hold 96-bit permutation result
	permutation1Result := []int{}

	// Permute the input using P-box 1. Each element in P-box 1 is a pointer to index of byte in input
	for i := 0; i < len(constants.PBox1); i++ {
		// Get the start and end index of a certain byte in input
		start := constants.PBox1[i] * 8
		end := start + 8

		// Put the byte in the permutation result
		permutation1Result = append(permutation1Result, input[start:end]...)
	}

	return permutation1Result
}

func substitution1(input []int, roundKey []int) []int {
	// Substitution 1 performs the substitution operation for a given input
	// using a round key. It utilizes the generated S-Box to perform substitution.

	// Trim the round key and perform a left cyclic shift
	trimmedKey := roundKey[1 : len(roundKey)-1]
	trimmedKey = CyclicShiftLeft(trimmedKey, 7)

	// Generate the S-Box using the trimmed key
	sBox := generateSBox(trimmedKey)

	// Convert the input into bit blocks
	inputBlocks := BinaryArrayToBitBlocks(input, 6)

	// Perform substitution for each input block using the generated S-Box
	for i, block := range inputBlocks {
		row := binaryArrayToInt([]int{block[1], block[4]})
		col := binaryArrayToInt([]int{block[0], block[2], block[3], block[5]})
		substituteValue := sBox[row][col]

		// Update the input block with the substituted value
		inputBlocks[i] = substituteValue
	}

	// Merge the substituted input blocks into a single array
	output := MergeBlockArrays(inputBlocks)

	return output
}

func permutation2(subsResult []int, roundKey []int) []int {
	// Permutation 2 performs permutation on the result of substitution operation
	// using a round key. It utilizes the permutation table (PBox2) for permutation.

	// Extract trimmed keys from the round key
	var trimmedKeys [][]int
	blockKey := BinaryArrayToBitBlocks(roundKey, 8)

	for _, block := range blockKey {
		trimmedKey := append(block[1:6], block[7])
		trimmedKeys = append(trimmedKeys, trimmedKey)
	}

	// XOR the result of substitution with trimmed keys
	res := XORBitArray(subsResult, MergeBlockArrays(trimmedKeys))

	// Permute the bits according to the PBox2 permutation table
	permutatedBits := intArrayPermutation(res, constants.PBox2)

	return permutatedBits
}

func substitution2(permutation2Result []int) []int {
	// Substitution 2 performs the substitution operation on the result of permutation
	// using a predefined S-Box for Rijndael. It utilizes the provided constants.SBoxRijndael
	// for substitution.

	// Convert the permutation result into bit blocks
	bitBlocks := BinaryArrayToBitBlocks(permutation2Result, 8)

	// Perform substitution for each bit block using the Rijndael S-Box
	for i, block := range bitBlocks {
		col := binaryArrayToInt(block[:4])
		row := binaryArrayToInt(block[4:])

		substitutedValue := constants.SBoxRijndael[row][col]
		bitBlocks[i] = intToBinaryArray(substitutedValue)[56:64]
	}

	// Merge the substituted bit blocks into a single array
	substitutedResult := MergeBlockArrays(bitBlocks)

	return substitutedResult
}

func permutation3(substitution2Result []int) []int {
	// Return the permutation of the result of substitution 2 using P-box 3
	//Each element in P-box 3 is a pointer to index of bit in input
	return intArrayPermutation(substitution2Result, constants.PBox3)
}
