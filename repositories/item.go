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
	ListItem(filters map[string]interface{}) ([]domain.Item, error)
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
	var item domain.Item

	if err := ir.db.First(&item, id).Error; err != nil {
		return domain.Item{}, err
	}

	return item, nil
}

func (ir *ItemRepository) CreateItem(i domain.Item) (domain.Item, error) {
	if err := ir.db.Create(&i).Error; err != nil {
		return domain.Item{}, err
	}

	return i, nil
}

func (ir *ItemRepository) UpdateItem(id uuid.UUID, oldItem domain.Item, newItem domain.Item) (domain.Item, error) {
	if err := ir.db.Save(&newItem).Error; err != nil {
		return domain.Item{}, err
	}

	return newItem, nil
}

func (ir *ItemRepository) ListItem(filters map[string]interface{}) ([]domain.Item, error) {

	var items []domain.Item

	q := ir.db.Where(filters)

	if err := q.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (ir *ItemRepository) DeleteItem(id uuid.UUID) error {
	tx := ir.db.Delete(&domain.Item{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
