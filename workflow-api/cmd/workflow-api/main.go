package main

import (
	"log"

	"github.com/YumDukFuu/workflow/internal/item"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// POST     /items
// GET		/items?status=xxxxx
// GET		/items/:id
// PUT		/items/:id
// PATCH    /items/:id
// DELETE	/items/:id
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

	// Register 📝 router
	r.POST("/items", controller.CreateItem)
	// Register 🔎 router
	r.GET("/items", controller.FindItems)
	// Register ✍️ router
	r.PATCH("/items/:id", controller.UpdateItemStatus)
	// Register  router /items/:id?title=ขอเบิกสินค้า&amount=12&quantity=10
	r.GET("/items/:id", controller.FindEachItem)

	// r.GET("/items/:id", func(ctx *gin.Context) {
	// 	// id := ctx.Param("id")
	// 	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"data": id,
	// 	})
	// 	// fmt.Printf("%#v\n", id)
	// })

	// Start server
	if err := r.Run(); err != nil {
		log.Panic(err)
	}

	// //3 2 1 Let Go
	// fmt.Println("Hello, World!")
	// //Siping GIN
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
