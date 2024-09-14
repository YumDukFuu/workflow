package item

import (
	"fmt"
	"net/http"

	"github.com/YumDukFuu/workflow/internal/model"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	// Setup Service
	Service Service
}

// For Contract with main.go
func NewController() Controller {
	return Controller{
		Service: NewService(),
	}
}

// Create item logic
func (controller Controller) CreateItem(ctx *gin.Context) {
	//Run Test [GIN-debug] POST   /items
	// ctx.JSON(http.StatusCreated, gin.H{
	// 	"message": "Success",
	// })
	//POST http://localhost:8080/items = Success

	// Binding with model
	var request model.RequestItem

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	//Check value Request
	fmt.Printf("%#v\n", request)
	// model.RequestItem{Title:"", Price:0, Quantity:0x0}

	// Create item
	item, err := controller.Service.Create(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"data": item,
	})

}
