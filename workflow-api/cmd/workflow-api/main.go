package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//3 2 1 Let Go
	// fmt.Println("Hello, World!")
	//Sip the GIN
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// Start server - short hand error handling
	if err := r.Run(); err != nil {
		log.Panic(err)
	}

}
