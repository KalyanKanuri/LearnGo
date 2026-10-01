package auth

import (
	"errors"
	"time"

	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func ValidateToken(tokenStr, secret, expectedUse string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), jwt.WithExpirationRequired())
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if claims.TokenUse != expectedUse {
		return nil, errors.New("unexpected token usage")
	}

	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenPwdHash(pwd []byte) (pwdHash []byte, err error) {
	pwdHash, err = bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return pwdHash, nil
}

func CheckPwdHash(pwdHash, pwd []byte) (is_valid bool, err error) {
	err = bcrypt.CompareHashAndPassword(pwdHash, pwd)
	if err != nil {
		return false, err
	}
	return true, nil
}
