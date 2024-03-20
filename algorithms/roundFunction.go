package algorithms

import (
	"block-cipher/constants"
)

func roundFunction(input []int, roundKey []int) []int {
	//****************** PERMUTATION 1 *******************//
	// Expansion permutation from 64-bit input to 96-bit input
	// This permutation operates in bytes

	// Initialize an empty variable to hold 96-bit permutation result
	permutation1Result := []int{}

	// Permute the input using P-box 1. Each element in P-box 1 is a pointer to index of byte in input
	for i := 0; i < len(constants.PBox1); i++ {
		// Get the start and end index of a certain byte in input
		start := constants.PBox1[i]*8
		end := start + 8

		// Put the byte in the permutation result
		permutation1Result = append(permutation1Result, input[start:end]...)
	}

	return []int{}
}
