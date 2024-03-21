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

	blockArrays := algorithms.CreateBlockArrays(ecbRequest.TextBitArray)

	if ecbRequest.Encrypt {
		result := ECBEncrypt(blockArrays, ecbRequest)
		c.JSON(http.StatusOK, models.ECBResponse{
		Success:        true,
		ResultBitArray: result,
		})
		return
	} else {
		result := ECBDecrypt(blockArrays, ecbRequest)
		c.JSON(http.StatusOK, models.ECBResponse{
		Success:        true,
		ResultBitArray: result,
		})
		return
	}
}

func ECBEncrypt(blockArrays [][]int, ecbRequest models.ECBRequest) []int {
	var cipherBitArray []int

	for _, block := range blockArrays {
		result := algorithms.NovaDSEncrypt(block, ecbRequest.KeyBitArray)
		cipherBitArray = append(cipherBitArray, result...)
	}

	return cipherBitArray
}

func ECBDecrypt(blockArrays [][]int, ecbRequest models.ECBRequest) []int {
	var plainBitArray []int

	for _, block := range blockArrays {
		result := algorithms.NovaDSDecrypt(block, ecbRequest.KeyBitArray)
		plainBitArray = append(plainBitArray, result...)
	}

	return plainBitArray
}
