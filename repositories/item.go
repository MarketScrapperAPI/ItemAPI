package repositories

import (
	"github.com/MarketScrapperAPI/ItemAPI/domain"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetItemById(id uuid.UUID) (domain.Item, error)
	CreateItem(i domain.Item) (domain.Item, error)
	UpdateItem(id uuid.UUID, oldItem domain.Item, newItem domain.Item) (domain.Item, error)
	ListItem(filters map[string]string) ([]domain.Item, error)
	DeleteItem(id uuid.UUID) error
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}
