package domain

import (
	pb "github.com/MarketScrapperAPI/ItemAPI/proto/gen"
	"github.com/gofrs/uuid"
)

type Item struct {
	Id               uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();"`
	Name             string    `gorm:"column:name"`
	Brand            string    `gorm:"column:brand"`
	Package          string    `gorm:"column:package"`
	PricePerItem     float32   `gorm:"column:price_per_item"`
	PricePerQuantity float32   `gorm:"column:price_per_quantity"`
	QuantityUnit     string    `gorm:"column:quantity_unit"`
	Url              string    `gorm:"column:url"`
	ImageUrl         string    `gorm:"column:image_url"`
}

func NewItem() Item {
	return Item{}
}

func (i *Item) FromCreateItemRequest(req *pb.CreateItemRequest) {
	i.Name = req.Name
	i.Brand = req.Brand
	i.Package = req.Brand
	i.PricePerItem = req.PricePerItem
	i.PricePerQuantity = req.PricePerQuantity
	i.QuantityUnit = req.QuantityUnit
	i.Url = req.Url
	i.ImageUrl = req.ImageUrl
}

func (i *Item) ToCreateItemResponse() *pb.CreateItemResponse {
	return &pb.CreateItemResponse{
		Uuid:             i.Id.String(),
		Name:             i.Name,
		Brand:            i.Brand,
		Package:          i.Package,
		PricePerItem:     i.PricePerItem,
		PricePerQuantity: i.PricePerQuantity,
		QuantityUnit:     i.QuantityUnit,
		Url:              i.Url,
		ImageUrl:         i.ImageUrl,
	}
}

func (i *Item) ToGetItemResponse() *pb.GetItemResponse {
	return &pb.GetItemResponse{
		Uuid:             i.Id.String(),
		Name:             i.Name,
		Brand:            i.Brand,
		Package:          i.Package,
		PricePerItem:     i.PricePerItem,
		PricePerQuantity: i.PricePerQuantity,
		QuantityUnit:     i.QuantityUnit,
		Url:              i.Url,
		ImageUrl:         i.ImageUrl,
	}
}
