package item

import (
	"net/http"
	"strconv"

	"github.com/YumDukFuu/workflow/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	// Setup Service
	Service Service
}

// For Contract with main.go
func NewController(db *gorm.DB) Controller {
	return Controller{
		Service: NewService(db),
	}
}

// func NewController() Controller {
// 	return Controller{
// 		Service: NewService(),
// 	}
// }

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
	// fmt.Printf("%#v\n", request)
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

//Find items

func (controller Controller) FindItems(ctx *gin.Context) {
	// Bind query parameters
	var (
		request model.RequestFindItem
	)
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	// Find
	items, err := controller.Service.Find(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

// internal/item/controller.go

func (controller Controller) UpdateItemStatus(ctx *gin.Context) {
	// Bind
	var (
		request model.RequestUpdateItem
	)
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	// Path param
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	//  สามารถใช้ด้านล่างแทนได้ ถ้าไม่ต้องการแปลงเลขฐานสอง
	//  id, _ := strconv.Atoi(ctx.Param("id"))
	// Update status
	item, err := controller.Service.UpdateStatus(uint(id), request.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

// 🗨📂///////////////
func (controller Controller) FindEachItem(ctx *gin.Context) {

	// Bind query parameters
	var request model.RequestFindParam
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	// fmt.Printf("%#v\n", request)
	// 	// Find
	// 	items, err := controller.Service.Find(request)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": err,
	// 		})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"data": items,
	// 	})

	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	// var id uint
	// id = 200
	// item, err := controller.Service.EachID(uint(id))

	item, err := controller.Service.EachID(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// EachID(id)
	ctx.JSON(http.StatusOK, gin.H{
		// "data": id,
		"data": item,
		// 	"param": request,
	})
}

// 🗨✏️///////////////
func (controller Controller) EditItem(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var request model.RequestItem

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	// var id uint
	// id = 200
	// item, err := controller.Service.EachID(uint(id))

	item, err := controller.Service.EditItemByID(uint(id), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// EachID(id)
	ctx.JSON(http.StatusOK, gin.H{
		// "data":  id,
		// "data2": request,
		"data": item,
	})
}
