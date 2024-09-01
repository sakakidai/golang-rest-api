package repositories

import (
	"golang-rest-api/models"

	"gorm.io/gorm"
)

type IContentItemRepository interface {
	GetAll(contentItems *[]models.ContentItem) error
	Create(contentItem *models.ContentItem) error
}

type contentItemRepository struct {
	db *gorm.DB
}

func NewContentItemRepository(db *gorm.DB) IContentItemRepository {
	return &contentItemRepository{db}
}

func (cir *contentItemRepository) GetAll(contentItems *[]models.ContentItem) error {
	if err := cir.db.Find(contentItems).Error; err != nil {
		return err
	}

	return nil
}

func (cir *contentItemRepository) Create(contentItem *models.ContentItem) error {
	if err := cir.db.Create(contentItem).Error; err != nil {
		return err
	}

	return nil
}
