package main

import (
	"log"
	"net/http"
	"os"

	"github.com/YumDukFuu/workflow/internal/item"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	// Test Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Print Foo:Bar
	// fmt.Println("FOO:", os.Getenv("FOO"))
	// fmt.Println("Test:", os.Getenv("TEST"))

	// Connect database
	dbsort := os.Getenv("Postgres")
	// db, err := gorm.Open(dbsort)
	db, err := gorm.Open(postgres.Open(dbsort))

	// db, err := gorm.Open(
	// 	postgres.Open(
	// 		"postgres://postgres:password@localhost:5432/task",
	// 	),
	// )

	if err != nil {
		log.Panic(err)
	}
	// env DB

	// Controller
	controller := item.NewController(db)
	// controller := item.NewController()

	// Router
	r := gin.Default()

	item := r.Group("/items")
	{
		// Register 📝 router
		item.POST("", controller.CreateItem)
		// Register 🔎 router
		item.GET("", controller.FindItems)
		// Register ✍️ router
		item.PATCH("/:id", controller.UpdateItemStatus)
		// Register 📂 router /items/:id?title=ขอเบิกสินค้า&amount=12&quantity=10
		item.GET("/:id", controller.FindEachItem)
		// Register ✏️ router
		item.PUT("/:id", controller.EditItem)
		// Register 🗑️ router
		item.DELETE("/:id", controller.DeleteItem)
	}

	// // Register 📝 router
	// r.POST("/items", controller.CreateItem)
	// // Register 🔎 router
	// r.GET("/items", controller.FindItems)
	// // Register ✍️ router
	// r.PATCH("/items/:id", controller.UpdateItemStatus)
	// // Register 📂 router /items/:id?title=ขอเบิกสินค้า&amount=12&quantity=10
	// r.GET("/items/:id", controller.FindEachItem)
	// // Register ✏️ router
	// r.PUT("/items/:id", controller.EditItem)
	// // Register 🗑️ router
	// r.DELETE("/items/:id", controller.DeleteItem)
	// // Register 🔐 router
	// r.POST("/login", controller.Login)

	///🚧///🚧///🚧///🚧///🚧///🚧///🚧///🚧///🚧///🚧///
	r.POST("/login", func(ctx *gin.Context) {
		// id := ctx.Param("id")
		// id64, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
		// id := uint(id64)
		ctx.JSON(http.StatusOK, gin.H{
			// "data": id,
			"Login": "Success",
		})
		// fmt.Printf("%#v\n", id)
	})

	// Start server
	// if err := r.Run(); err != nil {
	// 	log.Panic(err)
	// }

	// ENV Port
	port := os.Getenv("PORT")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
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
