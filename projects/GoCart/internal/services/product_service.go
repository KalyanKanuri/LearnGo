package services

import (
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/models"
	"github.com/KalyanKanuri/GoCart/internal/server/utils"
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

func (ps *ProductService) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {
	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := ps.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return convToCategoryResponse(&category), nil
}

func (ps *ProductService) GetCategories() ([]*dto.CategoryResponse, error) {
	var categories []models.Category

	if err := ps.db.Where("is_active=?", true).Find(&categories).Error; err != nil {
		return nil, err
	}

	resp := make([]*dto.CategoryResponse, len(categories))

	for i := range categories {
		resp[i] = convToCategoryResponse(&categories[i])
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

	return convToCategoryResponse(&category), nil
}

func (ps *ProductService) DeleteCategory(cat_id uint) (bool, error) {
	var category *models.Category

	if err := ps.db.Where("id=?", cat_id).Find(&category).Error; err != nil {
		return false, err
	}

	if err := ps.db.Delete(&category).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (ps *ProductService) GetProducts(page int, limit int) ([]*dto.ProductResponse, *utils.PaginationMeta, error) {
	if page < 1 {
		page = 0
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	var products []models.Product
	var total int64

	ps.db.Model(&models.Product{}).Where("is_active=?", true).Count(&total)
	if err := ps.db.Preload("Category").Preload("images").
		Where("is_active", true).
		Offset(offset).Limit(limit).
		Find(&products).Error; err != nil {
		return nil, nil, err
	}

	response := make([]*dto.ProductResponse, limit)
	for i := range products {
		response[i] = convToProductResponse(&products[i])
	}

	totalPages := (int64(total) - int64(limit-1)) / int64(limit)
	paginationInfo := &utils.PaginationMeta{
		Page:       uint(page),
		Limit:      uint(limit),
		Offset:     uint(offset),
		Total:      uint(total),
		TotalPages: uint(totalPages),
	}

	return response, paginationInfo, nil
}

func (ps *ProductService) GetProduct(productID uint) (*dto.ProductResponse, error) {
	var product models.Product
	if err := ps.db.Preload("Category").Preload("Images").Where("id=?", productID).First(&product).Error; err != nil {
		return nil, err
	}

	return convToProductResponse(&product), nil
}

func (ps *ProductService) AddProduct(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
		Stock:       req.Stock,
		SKU:         req.SKU,
	}

	if err := ps.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return ps.GetProduct(product.ID)
}

func (ps *ProductService) UpdateProduct(productID uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var product models.Product
	if err := ps.db.Where("id=?", productID).Find(&product).Error; err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.IsActive = req.IsActive

	if err := ps.db.Save(&product).Error; err != nil {
		return nil, err
	}

	return convToProductResponse(&product), nil
}

func (ps *ProductService) DeleteProduct(productID uint) (bool, error) {
	var product models.Product
	if err := ps.db.Where("id=?", product).Find(&product).Error; err != nil {
		return false, err
	}

	if err := ps.db.Delete(&product).Error; err != nil {
		return false, err
	}

	return true, nil
}
