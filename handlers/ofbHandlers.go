package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"block-cipher/algorithms"
	"block-cipher/models"
)

func HandleOFBRequest(c *gin.Context) {
	var ofbRequest models.OFBRequest
	err := json.NewDecoder(c.Request.Body).Decode(&ofbRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to decode request body"})
		return
	}

	if ofbRequest.Encrypt {
		result := OFBEncrypt(ofbRequest)
		c.JSON(http.StatusOK, models.OFBResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	} else {
		result := OFBDecrypt(ofbRequest)
		c.JSON(http.StatusOK, models.OFBResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	}
}

func OFBEncrypt(ofbRequest models.OFBRequest) []int {
	var cipherBitArray []int
	shiftRegister := ofbRequest.InitVector

	for i := 0; i < len(ofbRequest.TextBitArray); i += 8 {
		keystream := algorithms.NovaDSEncrypt(shiftRegister, ofbRequest.KeyBitArray)
		keystreamLSB := keystream[0:8]

		result := ofbRequest.TextBitArray[i : i+8]
		result = algorithms.XORBitArray(result, keystreamLSB)

		shiftRegister = shiftRegister[8:128]
		shiftRegister = append(shiftRegister, keystreamLSB...)
		cipherBitArray = append(cipherBitArray, result...)
	}

	return cipherBitArray
}

func OFBDecrypt(ofbRequest models.OFBRequest) []int {
	var plainBitArray []int
	shiftRegister := ofbRequest.InitVector

	for i := 0; i < len(ofbRequest.TextBitArray); i += 8 {
		keystream := algorithms.NovaDSEncrypt(shiftRegister, ofbRequest.KeyBitArray)
		keystreamLSB := keystream[0:8]

		result := ofbRequest.TextBitArray[i : i+8]
		result = algorithms.XORBitArray(result, keystreamLSB)

		shiftRegister = shiftRegister[8:128]
		shiftRegister = append(shiftRegister, keystreamLSB...)
		plainBitArray = append(plainBitArray, result...)
	}

	return plainBitArray
}
