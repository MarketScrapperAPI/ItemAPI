package domain

import "github.com/gofrs/uuid"

type Item struct {
	Id               uuid.UUID `gorm:"primary_key;"`
	Name             string    `gorm:"column:name"`
	Brand            string    `gorm:"column:brand"`
	Package          string    `gorm:"column:package"`
	PricePerItem     string    `gorm:"column:price_per_item"`
	PricePerQuantity string    `gorm:"column:price_per_quantity"`
	QuantityUnit     string    `gorm:"column:quantity_unit"`
	Url              string    `gorm:"column:url"`
	ImageUrl         string    `gorm:"column:image_url"`
}
