package auth

import (
	"time"

	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	TokenUse string `json:"token_use"`

	jwt.RegisteredClaims
}

func (claims JWTClaims) GenTokenPair(jwtCFG *config.JWTConfig) (accessToken, refreshToken string, err error) {
	now := time.Now()

	// Access Token
	claims.TokenUse = "Access"
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(
			now.Add(jwtCFG.JWTExpiration),
		),
		IssuedAt: jwt.NewNumericDate(now),
	}

	accessToken, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString(
		[]byte(jwtCFG.JWTSecret),
	)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	claims.TokenUse = "Refresh"
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(
			now.Add(jwtCFG.RefreshExpiration),
		),
		IssuedAt: jwt.NewNumericDate(now),
	}

	refreshToken, err = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	).SignedString([]byte(jwtCFG.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
