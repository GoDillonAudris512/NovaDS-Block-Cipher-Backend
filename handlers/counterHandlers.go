package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"block-cipher/algorithms"
	"block-cipher/models"
)

func HandleCounterRequest(c *gin.Context) {
	var counterRequest models.CounterRequest
	err := json.NewDecoder(c.Request.Body).Decode(&counterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to decode request body"})
		return
	}

	blockArrays := algorithms.CreateBlockArrays(counterRequest.TextBitArray)

	result := CounterMode(blockArrays, counterRequest)
	if counterRequest.Encrypt {
		c.JSON(http.StatusOK, models.CounterResponse{
			Success:        true,
			ResultBitArray: result,
		})
	} else {
		c.JSON(http.StatusOK, models.CounterResponse{
			Success:        true,
			ResultBitArray: algorithms.DeletePadding(result),
		})
	}
}

func CounterMode(blockArrays [][]int, counterRequest models.CounterRequest) []int {
	var processedBlockArray []int
	counter := make([]int, 128)

	for _, block := range blockArrays {
		keystream := algorithms.NovaDSEncrypt(counter, counterRequest.KeyBitArray)
		result := algorithms.XORBitArray(block, keystream)
		processedBlockArray = append(processedBlockArray, result...)
		incrementCounter(counter)
	}

	return processedBlockArray
}

func incrementCounter(counter []int) []int {
	carry := 1
	index := len(counter) - 1

	for carry > 0 && index >= 0 {
		counter[index] += carry
		carry = counter[index] / 2
		counter[index] %= 2
		index--
	}

	return counter
}
