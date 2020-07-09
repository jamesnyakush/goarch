package model

import "github.com/gofrs/uuid"

type ProductCategory struct {
	ProductCategoryId uuid.UUID `json:"product_category_id"`
	Category          string    `json:"category"`
}
