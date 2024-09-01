package usecases

import (
	"golang-rest-api/models"
	"golang-rest-api/repositories"
)

type IContentItemUsecase interface {
	GetAll() (map[string][]models.ContentItemResponse, error)
	Create(ci models.ContentItem) (models.ContentItemResponse, error)
}

type contentItemUsecase struct {
	cir repositories.IContentItemRepository
}

func NewContentItemUsecase(cir repositories.IContentItemRepository) IContentItemUsecase {
	return &contentItemUsecase{cir}
}

func (ciu *contentItemUsecase) GetAll() (map[string][]models.ContentItemResponse, error) {
	contentItems := []models.ContentItem{}
	if err := ciu.cir.GetAll(&contentItems); err != nil {
		return nil, err
	}

	resContentItems := []models.ContentItemResponse{}
	for _, ci := range contentItems {
		resContentItem := models.ContentItemResponse{
			ID:          ci.ID,
			Name:        ci.Name,
			Description: ci.Description,
			CreatedAt:   ci.CreatedAt,
			UpdatedAt:   ci.UpdatedAt,
		}
		resContentItems = append(resContentItems, resContentItem)
	}

	response := map[string][]models.ContentItemResponse{
		"contentItems": resContentItems,
	}
	return response, nil
}

func (ciu *contentItemUsecase) Create(ci models.ContentItem) (models.ContentItemResponse, error) {
	if err := ciu.cir.Create(&ci); err != nil {
		return models.ContentItemResponse{}, err
	}

	response := models.ContentItemResponse{
		ID:          ci.ID,
		Name:        ci.Name,
		Description: ci.Description,
		CreatedAt:   ci.CreatedAt,
		UpdatedAt:   ci.UpdatedAt,
	}
	return response, nil
}
