package algorithms

import (
	"crypto/md5"
	"math/rand"
	"slices"
)

// Generate 16 round keys for round function from external key
func GenerateRoundKeys(externalKey []int) [][]int {
	// Hash the external key using md5 hash
	bytes := binaryArrayToBytes(externalKey)
	hashResult := md5.Sum(bytes)
	hashedKey := bytesToBinaryArray(hashResult[:])

	// Initialize an empty variable to hold all round keys
	roundKeys := [][]int{}
	roundKey := hashedKey

	// Initialize a pseudo-random number generator using 64-LSB of hashed key
	source := rand.NewSource(int64(binaryArrayToInt(hashedKey[64:128])))
	rng := rand.New(source)

	// Initialize a pseudo-random P-box using reversed 64-MSB of hashed key
	msbHashedKey := hashedKey[0:64]
	slices.Reverse(msbHashedKey)
	pBoxRK := GeneratePBoxRK(binaryArrayToInt(msbHashedKey))

	// Loop for 16 times to generate 16 round keys
	for i := 0; i < 16; i++ {
		// Generate a random number
		random := rng.Int63()

		// XOR the 64-LSB of current round key with the random number
		randomBinaryArray := intToBinaryArray(int(random))
		xorResult := XORBitArray(roundKey[64:128], randomBinaryArray)

		// Permute the 64-MSB of current round key using P-box
		permutationResult := intArrayPermutation(roundKey[0:64], pBoxRK)

		// Make the new round key: XOR result in left side, permutation result in right side
		roundKey = append(xorResult, permutationResult...)

		// Add the new round key
		roundKeys = append(roundKeys, roundKey)
	}

	return roundKeys
}

// Generate a pseudo-random P-box using a number as a seed
func GeneratePBoxRK(seed int) []int {
	// Initialize a pseudo-random generator using the seed provided
	source := rand.NewSource(int64(seed))
	rng := rand.New(source)
	
	// Initialize the P-box
	pBoxRK := []int{}

	// Loop for 64 times, generate a random number from 0 to 63 and put the number in P-box
	for i := 0; i < 64; i++ {
		random := rng.Int63n(64)
		if !slices.Contains(pBoxRK, int(random)) {
			pBoxRK = append(pBoxRK, int(random))
		} else {
			pBoxRK = append(pBoxRK, 64)
		}
	}

	// Add the remaining number from 0 to 63 that hasn't been put in the P-box
	intNotInPBox := 0
	i := 0
	for i < 64 {
		if pBoxRK[i] == 64 {
			for slices.Contains(pBoxRK, intNotInPBox) {
				intNotInPBox++
			}
			pBoxRK[i] = intNotInPBox
			intNotInPBox++
		}
	}

	return pBoxRK
}