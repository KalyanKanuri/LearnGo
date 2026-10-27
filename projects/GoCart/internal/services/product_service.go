package services

import (
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type ProductService struct {
	db  *gorm.DB
	Log *zerolog.Logger
}

func NewProductService(db *gorm.DB, log *zerolog.Logger) *ProductService {
	return &ProductService{
		db:  db,
		Log: log,
	}
}

func (ps *ProductService) CreateCategroy(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := ps.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          int(category.ID),
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

func (ps *ProductService) GetCategories() ([]*dto.CategoryResponse, error) {
	var categories []*dto.CategoryResponse

	if err := ps.db.Where("is_active=?", true).Find(&categories).Error; err != nil {
		return nil, err
	}

	resp := make([]*dto.CategoryResponse, len(categories))

	for i, category := range categories {
		resp[i] = &dto.CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			IsActive:    category.IsActive,
		}
	}

	return resp, nil
}

func (ps *ProductService) UpdateCategory(cat_id uint, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {
	var category models.Category
	if err := ps.db.First(&category, cat_id).Error; err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Description = req.Description
	category.IsActive = req.IsActive

	if err := ps.db.Save(&category).Error; err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          int(category.ID),
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}
