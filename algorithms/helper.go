package algorithms

//****************** BLOCK ARRAY *******************//
// Function to divide the text bit array into blocks for block cipher processing
func CreateBlockArrays(bitArray []int) [][]int {
	var blockArrays [][]int
	var blockArray []int

	for i := 0; i < len(bitArray); i++ {
		blockArray = append(blockArray, bitArray[i])
		if (i+1)%128 == 0 {
			blockArrays = append(blockArrays, blockArray)
			blockArray = []int{}
		}
	}
	blockArrays = append(blockArrays, blockArray)

	lastArrayIdx := len(blockArrays) - 1
	if len(blockArrays[lastArrayIdx]) < 128 {
		for len(blockArrays[lastArrayIdx]) < 128 {
			blockArrays[lastArrayIdx] = append(blockArrays[lastArrayIdx], 0)
		}
	}

	return blockArrays
}

// Function to merge all block arrays into one single array
func MergeBlockArrays(blockArrays [][]int) []int {
	var bitArray []int

	for _, array := range blockArrays {
		bitArray = append(bitArray, array...)
	}

	return bitArray
}

//****************** BINARY ARRAY CONVERSION *******************//
// Turn a representation of binary array into array of bytes
func binaryArrayToBytes(binaryArray []int) []byte {
	numBytes := len(binaryArray) / 8
	bytes := make([]byte, numBytes)

	for i := 0; i < len(binaryArray); i++ {
		if binaryArray[i] == 1 {
			bytes[i/8] |= 1 << uint(7-(i%8))
		}
	}

	return bytes
}

// Turn an array of bytes into representation of binary array
func bytesToBinaryArray(bytes []byte) []int {
	binaryArray := make([]int, len(bytes)*8)

	for i, b := range bytes {
		for j := 0; j < 8; j++ {
			bit := (b >> uint(7-j)) & 0x01
			index := i*8 + j
			binaryArray[index] = int(bit)
		}
	}

	return binaryArray
}

// Turn a representation of binary array into its integer
func binaryArrayToInt(binaryArray []int) int {
	result := 0

	for _, bit := range binaryArray {
		result = (result << 1) | bit
	}

	return result
}

// Turn an integer into its representation of binary array
func intToBinaryArray(n int) []int {
	binaryArray := make([]int, 64)

	for i := 0; i < 64; i++ {
		binaryArray[i] = (n >> uint(64-1-i)) & 1
	}

	return binaryArray
}

// Turn binary array into bit blocks
func BinaryArrayToBitBlocks(binaryArray []int, blockSize int) [][]int {
    numBlocks := len(binaryArray) / blockSize
    bitBlocks := make([][]int, numBlocks)

    for i := 0; i < numBlocks; i++ {
        start := i * blockSize
        end := start + blockSize
        bitBlocks[i] = binaryArray[start:end]
    }

    return bitBlocks
}

// Do A Cyclic Shift Left on Bit Blocks
func CyclicShiftLeft(binaryArray []int, shiftAmount int) []int {
	arrayLen := len(binaryArray)
	shiftAmount %= arrayLen

	shiftedArray := make([]int, arrayLen)
	copy(shiftedArray, binaryArray)

	for i := 0; i < arrayLen; i++ {
		destIdx := (i - shiftAmount + arrayLen) % arrayLen
		shiftedArray[destIdx] = binaryArray[i]
	}

	return shiftedArray
}

// Do A Cyclic Shift Right on Bit Blocks
func CyclicShiftRight(binaryArray []int, shiftAmount int) []int {
	arrayLen := len(binaryArray)
	shiftAmount %= arrayLen

	shiftedArray := make([]int, arrayLen)
	copy(shiftedArray, binaryArray)

	for i := 0; i < arrayLen; i++ {
		destIdx := (i + shiftAmount) % arrayLen
		shiftedArray[destIdx] = binaryArray[i]
	}

	return shiftedArray
}

//****************** BINARY ARRAY OPERATION *******************//
// Function to do XOR operation between 2 arrays of bits with the same length
func XORBitArray(bitArrayA []int, bitArrayB []int) []int {
	result := []int{}

	for i := 0; i < len(bitArrayA); i++ {
		if bitArrayA[i] == bitArrayB[i] {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}
	}

	return result
}

// Permute an array of integer of n element using P-Box (P-box contains value of 0 to n)
func intArrayPermutation(array []int, pBox []int) []int {
	// Initialize an empty variable to hold permutation result
	permutationResult := []int{}

	// Loop through P-box. Each element in P-box is a pointer to index of element in array
	for i := 0; i < len(pBox); i++ {
		permutationResult = append(permutationResult, array[pBox[i]])
	}

	return permutationResult
}
