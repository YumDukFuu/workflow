package model

import "github.com/YumDukFuu/workflow/internal/constant"

type RequestItem struct {
	Title    string
	Price    float64
	Quantity uint
}

type RequestFindItem struct {
	Statuses []constant.ItemStatus
	// Statuses []constant.ItemStatus `form:"status[]"`
}

// type RequestUpdateItem struct {
// 	Status constant.ItemStatus
// }
