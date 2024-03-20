package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"block-cipher/algorithms"
	"block-cipher/models"
)

func HandleECBRequest(c *gin.Context) {
	var ecbRequest models.ECBRequest
	err := json.NewDecoder(c.Request.Body).Decode(&ecbRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to decode request body"})
		return
	}

	if ecbRequest.Encrypt {
		result := ECBEncrypt(ecbRequest)
		c.JSON(http.StatusOK, models.ECBResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	} else {
		result := ECBDecrypt(ecbRequest)
		c.JSON(http.StatusOK, models.ECBResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	}
}

func ECBEncrypt(ecbRequest models.ECBRequest) []int {
	var cipherBitArray []int
	plaintext := ecbRequest.TextBitArray
	key := ecbRequest.KeyBitArray

	bitBlocks := algorithms.BinaryArrayToBitBlocks(plaintext)

	for _, block := range bitBlocks {
		encryptedBlock := algorithms.XORBitArray(block, key)
		encryptedBlock = algorithms.CyclicShiftLeft(encryptedBlock, 1)
		cipherBitArray = append(cipherBitArray, encryptedBlock...)
	}

	return cipherBitArray
}

func ECBDecrypt(ecbRequest models.ECBRequest) []int {
	var plainBitArray []int
	ciphertext := ecbRequest.TextBitArray
	key := ecbRequest.KeyBitArray

	bitBlocks := algorithms.BinaryArrayToBitBlocks(plaintext)

	for _, block := range bitBlocks {
		encryptedBlock := algorithms.CyclicShiftRight(encryptedBlock, 1)
		encryptedBlock = algorithms.XORBitArray(block, key)
		plainBitArray = append(plainBitArray, encryptedBlock...)
	}

	return plainBitArray
}
