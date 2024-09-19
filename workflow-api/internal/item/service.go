package item

import (
	"github.com/YumDukFuu/workflow/internal/constant"
	"github.com/YumDukFuu/workflow/internal/model"
	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
}

// func NewService() Service {
// 	return Service{}
// }

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (service Service) Create(req model.RequestItem) (model.Item, error) {
	item := model.Item{
		Title:  req.Title,
		Amount: req.Amount,
		// Price:    req.Price,
		Quantity: req.Quantity,
		Status:   constant.ItemPendingStatus,
		Owner_id: 510680,
	}
	if err := service.Repository.Create(&item); err != nil {
		return model.Item{}, err
	}
	return item, nil
}

func (service Service) Find(query model.RequestFindItem) ([]model.Item, error) {
	return service.Repository.Find(query)
}

func (service Service) UpdateStatus(id uint, status constant.ItemStatus) (model.Item, error) {
	// Find item
	item, err := service.Repository.FindByID(id)
	if err != nil {
		return model.Item{}, err
	}
	// Fill data
	item.Status = status
	// Replace
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}
	return item, nil
}

// 🗨📂///////////////
// func (service Service) EachID(id uint) uint {
// func (service Service) EachID(id uint) (model.Item, error) {
func (service Service) EachID(id uint, param model.RequestFindParam) (model.Item, error) {
	// test := service.Repository.FindEach(id)
	// fmt.Printf("%#v\n", param)
	// test, err := service.Repository.FindEach(id)
	test, err := service.Repository.FindEach(id, param)
	if err != nil {
		return model.Item{}, err
	}
	// return test
	return test, nil
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"data": id,
	// 	})
}

// 🗨✏️///////////////
// func (service Service) EditItemByID(id uint, req model.RequestItem) {
func (service Service) EditItemByID(id uint, req model.RequestItem) (model.Item, error) {

	// fmt.Printf("EDIT-SERVICE ID %#v\n", id)
	// fmt.Printf("EDIT-SERVICE REQ %#v\n", item)
	// service.Repository.EditEach(id)
	item, err := service.Repository.FindEachID(id)
	if err != nil {
		return model.Item{}, err
	}
	// fmt.Println("EDIT-Repo ITEM B4", item)

	// Fill data
	item.Amount = req.Amount
	item.Quantity = req.Quantity

	// fmt.Println("EDIT-Repo ITEM AFT", item)
	// Sent to DB
	if err := service.Repository.Replace(item); err != nil {
		return model.Item{}, err
	}

	return item, nil

}
