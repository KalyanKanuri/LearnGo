package services

import (
	"fmt"

	"github.com/KalyanKanuri/GoCart/internal/auth"
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/models"
	"gorm.io/gorm"
)

func (as *AuthService) genAUthResp(tx *gorm.DB, user *models.User) (*dto.AuthResponse, error) {
	claims := auth.JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   string(user.Role),
	}

	accessToken, refreshToken, err := claims.GenTokenPair(&as.cfg.JWT)
	if err != nil {
		return nil, fmt.Errorf("error generating token pair %w", err)
	}

	reftokenModel := models.RefreshToken{
		UserID:    claims.UserID,
		Token:     refreshToken,
		ExpiresAt: claims.ExpiresAt.Time,
	}

	if err := tx.Create(&reftokenModel).Error; err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		User: &dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      string(user.Role),
			Phone:     user.Phone,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func convToCategoryResponse(category *models.Category) *dto.CategoryResponse {
	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

func convToProductResponse(product *models.Product) *dto.ProductResponse {
	images := make([]dto.ProductImagesResponse, len(product.Images))

	for i := range product.Images {
		images[i] = dto.ProductImagesResponse{
			ID:        product.Images[i].ID,
			URL:       product.Images[i].URL,
			CreatedAt: product.Images[i].CreatedAt,
			UpdatedAt: product.Images[i].UpdatedAt,
		}
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		SKU:         product.SKU,
		IsActive:    product.IsActive,
		Category: dto.CategoryResponse{
			ID:          product.CategoryID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
			IsActive:    product.Category.IsActive,
			CreatedAt:   product.Category.CreatedAt,
			UpdatedAt:   product.Category.UpdatedAt,
		},
		Images:    images,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
