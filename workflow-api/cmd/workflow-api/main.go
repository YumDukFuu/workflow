package main

import (
	"log"

	"github.com/YumDukFuu/workflow/internal/item"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect database
	db, err := gorm.Open(
		postgres.Open(
			"postgres://postgres:password@localhost:5432/task",
		),
	)
	if err != nil {
		log.Panic(err)
	}

	// Controller
	controller := item.NewController(db)
	// controller := item.NewController()

	// Router
	r := gin.Default()

	// Register router
	r.POST("/items", controller.CreateItem)

	// Start server
	if err := r.Run(); err != nil {
		log.Panic(err)
	}

	// //3 2 1 Let Go
	// fmt.Println("Hello, World!")
	// //Sip the GIN
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// // Start server - short hand error handling
	// if err := r.Run(); err != nil {
	// 	log.Panic(err)
	// }
}
