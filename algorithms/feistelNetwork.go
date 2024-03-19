package algorithms

// Implement the encryption method of a feistel network
func feistelNetworkEncrypt(left []int, right []int, roundKey []int) ([]int, []int) {
	// Find the result of applying round function and round key to the right side of block
	result := roundFunction(right, roundKey)

	// Return as follow,
	// Left side	: the original right side
	// Right side 	: XOR between the original left side with the result of round function
	return right, XORBitArray(left, result)
}

// Implement the decryption method of a feistel network
func feistelNetworkDecrypt(left []int, right []int, roundKey []int) ([]int, []int) {
	// Find the result of applying round function and round key to the left side of block
	result := roundFunction(left, roundKey)

	// Return as follow,
	// Left side	: XOR between the original right side with the result of round function
	// Right side 	: the original left side
	return XORBitArray(right, result), left
}
