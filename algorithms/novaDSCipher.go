package algorithms

// Implement the encryption function of NovaDS Cipher
func NovaDSEncrypt(blockArray []int, externalKey []int) []int {
	//****************** ROUND KEYS *******************//
	// Generate 16 round keys from external key to be used during each round
	roundKeys := GenerateRoundKeys(externalKey)
	
	//**************** FEISTEL NETWORK ****************//
	// Split the block to 64-bit left and right side to prepare to enter the feistel network
	left := blockArray[0:64]
	right := blockArray[64:128]

	// For 16 rounds, use the feistel network (encryption method) to encrypt the block
	newLeft := left
	newRight := right
	for i := 0; i < 16; i++ {
		newLeft, newRight = feistelNetworkEncrypt(newLeft, newRight, roundKeys[i])
	}

	// Merge the results of feistel network back to form 128-bit block
	feistelResult := append(newLeft, newRight...) 

	return feistelResult
}

func NovaDSDecrypt(blockArray []int, externalKey []int) []int {
	//****************** ROUND KEYS *******************//
	// Generate 16 round keys from external key to be used during each round
	roundKeys := GenerateRoundKeys(externalKey)

	//**************** FEISTEL NETWORK ****************//
	// Split the block to 64-bit left and right side to prepare to enter the feistel network
	left := blockArray[0:64]
	right := blockArray[64:128]

	// For 16 rounds, use the feistel network (decryption method) to decrypt the block
	newLeft := left
	newRight := right
	for i := 15; i >= 0; i-- {
		newLeft, newRight = feistelNetworkDecrypt(newLeft, newRight, roundKeys[i])
	}

	// Merge the results of feistel network back to form 128-bit block
	feistelResult := append(newLeft, newRight...) 

	return feistelResult
}
