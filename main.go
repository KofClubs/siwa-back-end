package main

import (
	"net/http"

	"github.com/KofClubs/siwa-back-end/config"
	"github.com/KofClubs/siwa-back-end/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handlePublicKey(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{
		"public_key": config.Rsa2PublicKeyPem,
	})
}

func handleSign(c *gin.Context) {
	var signRequest handler.SignRequest
	if err := c.BindJSON(&signRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	signResponse, err := signRequest.Handle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "OPTIONS, POST")
	c.JSON(http.StatusOK, gin.H{
		"message": signResponse.Message,
		"sign":    signResponse.Sign,
	})
}

func handleUpload(c *gin.Context) {
	var uploadRequest handler.UploadRequest
	if err := c.BindJSON(&uploadRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	uploadResponse, err := uploadRequest.Handle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "OPTIONS, POST")
	c.JSON(http.StatusOK, gin.H{
		"key": uploadResponse.Key,
	})
}

func main() {
	r := gin.Default()
	// upload trade info
	r.GET("/public_key", handlePublicKey)
	r.POST("/sign", handleSign)
	r.POST("/upload", handleUpload)
	// download trade info
	r.Use(cors.Default())
	err := r.Run()
	if err != nil {
		return
	}
}
