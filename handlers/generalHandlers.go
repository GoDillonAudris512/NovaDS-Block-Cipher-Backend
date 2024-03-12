package handlers

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func HelloHandler(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hello, Welcome to Modern Block Cipher Web Service!")
}