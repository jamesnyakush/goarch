package model

import "github.com/gofrs/uuid"

type Order struct {
	OrderId     uuid.UUID `json:"order_id"`
	Quantity    uint      `json:"quantity"`
	UserId      uuid.UUID `json:"user_id"`
	ProductId   uuid.UUID `json:"product_id"`
	OrderStatus uuid.UUID `json:"order_status"`
}
