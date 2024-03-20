package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"fmt"

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

	if counterRequest.Encrypt {
		result := CounterEncrypt(counterRequest)
		c.JSON(http.StatusOK, models.CounterResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	} else {
		result := CounterDecrypt(counterRequest)
		c.JSON(http.StatusOK, models.CounterResponse{
			Success:        true,
			ResultBitArray: result,
		})
		return
	}
}

func CounterEncrypt(counterRequest models.CounterRequest) []int {
	var cipherBlockArray [][]int
	counter := make([]int, 8)

	for i := 0; i < len(plaintext); i += 8 {
		result := algorithms.NovaDSEncrypt(counterRequest.KeyBitArray, counter)
		result = algorithms.XORBitArray(result, counterRequest.TextBitArray)

		cipherBlockArray = append(cipherBlockArray, result)
		for i := len(counter) - 1; i >= 0; i-- {
			if counter[i] == 0 {
				counter[i] = 1
				break
			} else {
				counter[i] = 0
			}
		}
	}

	cipherBitArray := algorithms.MergeBlockArrays(cipherBlockArray)

	return cipherBitArray
}

func CounterDecrypt(cipherRequest models.CipherRequest) []int {
    var plaintextBlockArray [][]int
    counter := make([]int, 8)

    for i := 0; i < len(cipherRequest.CipherBitArray); i += 8 {
        result := algorithms.NovaDSEncrypt(cipherRequest.KeyBitArray, counter)
        plaintextBlock := algorithms.XORBitArray(result, cipherRequest.CipherBitArray[i:i+8])

        plaintextBlockArray = append(plaintextBlockArray, plaintextBlock)

        // Increment counter
        for j := len(counter) - 1; j >= 0; j-- {
            if counter[j] == 1 {
                counter[j] = 0
            } else {
                counter[j] = 1
                break
            }
        }
    }

    plaintextBitArray := algorithms.MergeBlockArrays(plaintextBlockArray)

    return plaintextBitArray
}
