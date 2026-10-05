package services

import (
	"fmt"

	"github.com/KalyanKanuri/GoCart/internal/auth"
	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/models"
)

func genAUthResp(jwtCFG *config.JWTConfig, user *models.User) (*dto.AuthResponse, error) {
	claims := auth.JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   string(user.Role),
	}

	accessToken, refreshToken, err := claims.GenTokenPair(jwtCFG)
	if err != nil {
		return nil, fmt.Errorf("error generating token pair %w", err)
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
