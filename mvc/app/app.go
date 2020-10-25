package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var(
	router *gin.Engine
)

func init(){
	router = gin.Default()
}

func StartApp(){
	mapUrls()

	log.Printf("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
