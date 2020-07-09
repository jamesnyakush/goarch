package model

import "github.com/gofrs/uuid"

type OrderStatus struct {
	OrderStatusId uuid.UUID `json:"order_status_id"`
	Name          string    `json:"name"`
}
