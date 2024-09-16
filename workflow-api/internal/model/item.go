package model

import "github.com/YumDukFuu/workflow/internal/constant"

type Item struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	// Price    float64             `json:"price"`
	Amount   uint                `json:"amount"`
	Quantity uint                `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
	Owner_id uint                `json:"owner_id"`
	// ID       uint `gorm:"primaryKey"`
	// Title    string
	// Price    float64
	// Quantity uint
	// Status   constant.ItemStatus
}
