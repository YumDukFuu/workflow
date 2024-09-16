package model

import "github.com/YumDukFuu/workflow/internal/constant"

type RequestItem struct {
	Title string
	// Price    float64
	// Quantity uint
	Amount   uint
	Quantity uint
}

type RequestFindItem struct {
	Statuses []constant.ItemStatus `form:"status[]"`
	// Statuses []constant.ItemStatus
}

type RequestUpdateItem struct {
	Status constant.ItemStatus
}
