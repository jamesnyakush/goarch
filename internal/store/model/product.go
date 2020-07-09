package model

import "github.com/gofrs/uuid"

type Product struct {
	ProductId         uuid.UUID `json:"product_id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Quantity          string    `json:"quantity"`
	Cost              string    `json:"cost"`
	Images            []Image   `json:"images"`
	UserId            uint      `json:"user_id"`
	ProductCategoryId uuid.UUID `json:"product_category_id"`
}

type Image struct {
	ImageId   uuid.UUID `json:"image_id"`
	ProductId uuid.UUID `json:"product_id"`
	ImageUrl  string    `json:"image_url"`
}
