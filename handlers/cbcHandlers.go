package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"block-cipher/algorithms"
	"block-cipher/models"
)

func HandleCBCRequest(c *gin.Context) {
	var cbcRequest models.CBCRequest
	err := json.NewDecoder(c.Request.Body).Decode(&cbcRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to decode request body"})
		return
	}

	blockArrays := algorithms.CreateBlockArrays(cbcRequest.TextBitArray)

	if cbcRequest.Encrypt {
		result := CBCEncrypt(cbcRequest, blockArrays)
		c.JSON(http.StatusOK, models.CBCResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	} else {
		result := CBCDecrypt(cbcRequest, blockArrays)
		c.JSON(http.StatusOK, models.CBCResponse{
			Success:        true,
			ResultBitArray: algorithms.DeletePadding(result),
		})
		return
	}
}

func CBCEncrypt(cbcRequest models.CBCRequest, blockArrays [][]int) []int {
	var cipherBlockArrays [][]int
	xorVector := cbcRequest.InitVector

	for i := 0; i < len(blockArrays); i++ {
		result := algorithms.XORBitArray(blockArrays[i], xorVector)
		result = algorithms.NovaDSEncrypt(result, cbcRequest.KeyBitArray)

		xorVector = result
		cipherBlockArrays = append(cipherBlockArrays, result)
	}

	return algorithms.MergeBlockArrays(cipherBlockArrays)
}

func CBCDecrypt(cbcRequest models.CBCRequest, blockArrays [][]int) []int {
	var plainBlockArrays [][]int
	xorVector := cbcRequest.InitVector

	for i := 0; i < len(blockArrays); i++ {
		result := algorithms.NovaDSDecrypt(blockArrays[i], cbcRequest.KeyBitArray)
		result = algorithms.XORBitArray(result, xorVector)

		xorVector = blockArrays[i]
		plainBlockArrays = append(plainBlockArrays, result)
	}

	return algorithms.MergeBlockArrays(plainBlockArrays)
}
