package handlers

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
			Success: true,
			ResultBitArray: result,
		})
		return
	} else {
		result := ECBDecrypt(ecbRequest)
		c.JSON(http.StatusOK, models.ECBResponse{
			Success: true,
			ResultBitArray: result,
		})
		return
	}
}

func ECBEncrypt(ecbRequest models.ECBRequest) []int {
	var cipherBitArray []int
	plaintext := ecbRequest.PlaintextBitArray
	key := ecbRequest.KeyBitArray

	for i := 0; i < len(plaintext); i += len(key) {
		end := i + len(key)
		if end > len(plaintext) {
			end = len(plaintext)
		}
		block := plaintext[i:end]
		encryptedBlock := algorithms.Encrypt(block, key)
		encryptedBlock := append(encryptedBlock[1:3], encryptedBlock[0])

		cipherBitArray = append(cipherBitArray, encryptedBlock...)
	}

	return cipherBitArray
}

func ECBDecrypt(ecbRequest models.ECBRequest) []int {
	var plainBitArray []int
	ciphertext := ecbRequest.CiphertextBitArray
	key := ecbRequest.KeyBitArray

	for i := 0; i < len(ciphertext); i += len(key) {
		end := i + len(key)
		if end > len(ciphertext) {
			end = len(ciphertext)
		}
		block := ciphertext[i:end]
		block := append(block[0], block[1:3])
		decryptedBlock := algorithms.Decrypt(block, key)

		plainBitArray = append(plainBitArray, decryptedBlock...)
	}

	return plainBitArray
}
