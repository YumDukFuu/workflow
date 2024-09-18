package model

import "github.com/YumDukFuu/workflow/internal/constant"

type RequestItem struct {
	Title string `json:"title"`
	// Price    float64
	// Quantity uint
	Amount   uint `json:"amount"`
	Quantity uint `json:"quantity"`
}

type RequestFindItem struct {
	Statuses []constant.ItemStatus `form:"status[]"`
	// Statuses []constant.ItemStatus
}

type RequestUpdateItem struct {
	Status constant.ItemStatus
}

// ////////////////////////////
type RequestFindParam struct {
	Title    string `form:"title"`
	Amount   uint   `form:"amount"`
	Quantity uint   `form:"quantity"`
}
