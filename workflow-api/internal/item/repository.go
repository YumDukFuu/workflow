package item

import (
	"github.com/YumDukFuu/workflow/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Database: db,
	}
}

func (repo Repository) Create(item *model.Item) error {
	return repo.Database.Create(item).Error
}

func (repo Repository) Find(query model.RequestFindItem) ([]model.Item, error) {
	var results []model.Item
	db := repo.Database
	if statuses := query.Statuses; len(statuses) > 0 {
		db = db.Where("status IN ?", statuses).Order("id ASC")
	}
	if err := db.Find(&results).Error; err != nil {
		return results, err
	}
	return results, nil
}

func (repo Repository) Replace(item model.Item) error {
	return repo.Database.Model(&item).Updates(item).Error
}

func (repo Repository) FindByID(id uint) (model.Item, error) {
	var result model.Item
	if err := repo.Database.First(&result, id).Error; err != nil {
		return result, err
	}
	return result, nil
}

// 🗨📂///////////////
// func (repo Repository) FindEach(id uint) (model.Item, error) {
func (repo Repository) FindEach(id uint, param model.RequestFindParam) (model.Item, error) {
	var result model.Item
	// fmt.Printf("IN REPO TEST %#v\n", param)
	// if err := repo.Database.First(&result, id).Error; err != nil {
	// 	return result, err
	// }
	db := repo.Database
	// if statuses := query.Statuses; len(statuses) > 0 {
	// 	db = db.Where("status IN ?", statuses)
	// }
	if err := db.Where("title = ?", param.Title).Where("amount = ?", param.Amount).Where("quantity = ?", param.Quantity).First(&result, id).Error; err != nil {
		return result, err
	}
	// return results, nil
	// fmt.Printf("IN REPO TEST %#v\n", result)
	return result, nil
}

// 🗨✏️///////////////
func (repo Repository) FindEachID(id uint) (model.Item, error) {
	var result model.Item
	// fmt.Printf("IN REPO TEST %#v\n", param)
	// if err := repo.Database.First(&result, id).Error; err != nil {
	// 	return result, err
	// }
	db := repo.Database
	// if statuses := query.Statuses; len(statuses) > 0 {
	// 	db = db.Where("status IN ?", statuses)
	// }
	if err := db.First(&result, id).Error; err != nil {
		return result, err
	}
	// return results, nil
	// fmt.Printf("IN REPO TEST %#v\n", result)
	return result, nil
}
