package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"block-cipher/algorithms"
	"block-cipher/models"
)

func HandleCFBRequest(c *gin.Context) {
	var cfbRequest models.CFBRequest
	err := json.NewDecoder(c.Request.Body).Decode(&cfbRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to decode request body"})
		return
	}

	if cfbRequest.Encrypt {
		result := CFBEncrypt(cfbRequest)
		c.JSON(http.StatusOK, models.CFBResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	} else {
		result := CFBDecrypt(cfbRequest)
		c.JSON(http.StatusOK, models.CFBResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	}
}

func CFBEncrypt(cfbRequest models.CFBRequest) []int {
	var cipherBitArray []int
	shiftRegister := cfbRequest.InitVector

	for i := 0; i < len(cfbRequest.TextBitArray); i += 8 {
		keystream := algorithms.NovaDSEncrypt(shiftRegister, cfbRequest.KeyBitArray)
		keystreamLSB := keystream[0:8]

		result := cfbRequest.TextBitArray[i : i+8]
		result = algorithms.XORBitArray(result, keystreamLSB)

		shiftRegister = shiftRegister[8:128]
		shiftRegister = append(shiftRegister, result...)
		cipherBitArray = append(cipherBitArray, result...)
	}

	return cipherBitArray
}

func CFBDecrypt(cfbRequest models.CFBRequest) []int {
	var plainBitArray []int
	shiftRegister := cfbRequest.InitVector

	for i := 0; i < len(cfbRequest.TextBitArray); i += 8 {
		keystream := algorithms.NovaDSEncrypt(shiftRegister, cfbRequest.KeyBitArray)
		keystreamLSB := keystream[0:8]

		result := cfbRequest.TextBitArray[i : i+8]
		result = algorithms.XORBitArray(result, keystreamLSB)

		shiftRegister = shiftRegister[8:128]
		shiftRegister = append(shiftRegister, (cfbRequest.TextBitArray[i : i+8])...)
		plainBitArray = append(plainBitArray, result...)
	}

	return plainBitArray
}
