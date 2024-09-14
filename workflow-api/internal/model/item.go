package model

import "github.com/YumDukFuu/workflow/internal/constant"

type Item struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Price    float64
	Quantity uint
	Status   constant.ItemStatus
}
