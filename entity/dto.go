package entity

import "time"

type GeneralResponse struct {
	Status  int
	Message string
}

type OrderTrackingResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    OrderOverview `json:"data"`
}

type OrderOverview struct {
	Id          uint        `json:"id"`
	Status      OrderStatus `json:"status"`
	ReferenceId uint        `json:"reference_id"`
	CreatedAt   time.Time   `json:"created_at"`
}
