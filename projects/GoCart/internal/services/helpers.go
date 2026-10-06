package services

import (
	"fmt"

	"github.com/KalyanKanuri/GoCart/internal/auth"
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/models"
)

func (as *AuthService) genAUthResp(user *models.User) (*dto.AuthResponse, error) {
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
		Token:     accessToken,
		ExpiresAt: claims.ExpiresAt.Time,
	}

	if err := as.db.Create(&reftokenModel).Error; err != nil {
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
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
