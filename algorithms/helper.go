package algorithms


// Function to divide the text bit array into blocks for block cipher processing
func CreateBlockArrays(bitArray []int) [][]int{
	var blockArrays [][]int
	var blockArray []int

	for i := 0; i < len(bitArray); i++ {
		blockArray = append(blockArray, bitArray[i])
		if (i + 1) % 128 == 0 {
			blockArrays = append(blockArrays, blockArray)
			blockArray = []int{}
		}
	}

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