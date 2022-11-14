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

func (ir *ItemRepository) GetItemById(id uuid.UUID) (domain.Item, error) {
	return domain.Item{}, nil
}

func (ir *ItemRepository) CreateItem(i domain.Item) (domain.Item, error) {
	if err := ir.db.Create(&i).Error; err != nil {
		return domain.Item{}, err
	}

	return i, nil
}

func (ir *ItemRepository) UpdateItem(id uuid.UUID, oldItem domain.Item, newItem domain.Item) (domain.Item, error) {
	return domain.Item{}, nil
}

func (ir *ItemRepository) ListItem(filters map[string]string) ([]domain.Item, error) {
	return []domain.Item{}, nil
}

func (ir *ItemRepository) DeleteItem(id uuid.UUID) error {
	return nil
}
