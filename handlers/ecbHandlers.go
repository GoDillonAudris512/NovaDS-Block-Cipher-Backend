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

	cipherBitArray = algorithms.NovaDSEncrypt(plaintext, key)

	return cipherBitArray
}

func ECBDecrypt(ecbRequest models.ECBRequest) []int {
	var plainBitArray []int
	ciphertext := ecbRequest.TextBitArray
	key := ecbRequest.KeyBitArray

	plainBitArray = algorithms.NovaDSDncrypt(ciphertext, key)

	return plainBitArray
}
