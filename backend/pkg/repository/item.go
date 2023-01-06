package repository

import (
	"github.com/khishh/personal-finance-app/graph/model"
	"github.com/khishh/personal-finance-app/pkg/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(itemInput *model.ItemInput) (*models.Item, error)

	GetOneItem(sub string) (*models.Item, error)
	GetAllItemOfOneUser(sub string) (*[]models.Item, error)
}

type ItemService struct {
	Db *gorm.DB
}

var _ ItemRepository = &ItemService{}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{
		Db: db,
	}
}

func (i *ItemService) CreateItem(itemInput *model.ItemInput) (*models.Item, error) {
	item := &models.Item{
		ID:          itemInput.ID,
		UserSub:     itemInput.UserSub,
		AccessToken: itemInput.AccessToken,
		RequestId:   itemInput.RequestID,
		LastCursor:  nil,
	}

	err := i.Db.Create(item).Error
	return item, err
}

func (i *ItemService) GetOneItem(id string) (*models.Item, error) {
	item := &models.Item{}
	err := i.Db.Where("id = ?", id).First(item).Error
	return item, err
}

func (i *ItemService) GetAllItemOfOneUser(sub string) (*[]models.Item, error) {
	items := &[]models.Item{}
	err := i.Db.Where("user_sub = ?", sub).Find(items).Error
	return items, err
}
