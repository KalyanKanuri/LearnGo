package services

import (
	"errors"
	"fmt"

	"github.com/KalyanKanuri/GoCart/internal/auth"
	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(dbConn *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{
		db:  dbConn,
		cfg: cfg,
	}
}

func (as *AuthService) RegisterUser(regReq *dto.RegisterRequest) (*dto.AuthResponse, error) {
	hashedPwd, err := auth.GenPwdHash([]byte(regReq.Password))
	if err != nil {
		return nil, fmt.Errorf("error generating password hash %w", err)
	}

	var response *dto.AuthResponse
	txErr := as.db.Transaction(func(tx *gorm.DB) error {
		var existingUser models.User

		err := tx.Select("id").Where("email=?", regReq.Email).First(&existingUser).Error
		switch {
		case err == nil:
			return errors.New("user already exists")
		case !errors.Is(err, gorm.ErrRecordNotFound):
			return fmt.Errorf("checking for existing user: %w", err)
		}

		requestedUser := models.User{
			Email:     regReq.Email,
			FirstName: regReq.FirstName,
			LastName:  regReq.LastName,
			Password:  hashedPwd,
			Role:      models.UserRoleCustomer,
			Phone:     regReq.Phone,
			IsActive:  true,
		}
		if usrErr := tx.Create(&requestedUser).Error; usrErr != nil {
			return usrErr
		}

		cart := models.Cart{
			UserID: requestedUser.ID,
		}
		if cartErr := tx.Create(&cart).Error; cartErr != nil {
			return cartErr
		}

		response, err = as.genAUthResp(&requestedUser)
		if err != nil {
			return err
		}
		return nil
	})
	if txErr != nil {
		return nil, txErr
	}

	return response, nil
}

func (as *AuthService) LoginUser(loginReq *dto.LoginRequest) (*dto.AuthResponse, error) {
	var response *dto.AuthResponse

	var user models.User
	if loginErr := as.db.Where("email = ?", loginReq.Email).Find(&user).Error; loginErr != nil {
		return nil, loginErr
	}

	is_valid, err := auth.CheckPwdHash([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, err
	}

	if !is_valid {
		return nil, errors.New("invalid credentials")
	}

	response, err = as.genAUthResp(&user)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (as *AuthService) RefreshToken(reftokenReq *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	claims, err := auth.ValidateToken(reftokenReq.RefreshToken, as.cfg.JWT.JWTSecret, "Refresh")
	if err != nil {
		return nil, err
	}

	var refreshToken *models.RefreshToken
	if err := as.db.Where("token=?", reftokenReq.RefreshToken).Find(refreshToken).Error; err != nil {
		return nil, err
	}
	as.db.Delete(refreshToken)

	var user *models.User
	if err := as.db.Where("id=?", claims.ID).Find(user).Error; err != nil {
		return nil, err
	}
	return as.genAUthResp(user)
}

func (as *AuthService) Logout(logoutReq *dto.RefreshTokenRequest) error {
	_, err := auth.ValidateToken(logoutReq.RefreshToken, as.cfg.JWT.JWTSecret, "Refresh")
	if err != nil {
		return err
	}

	var refreshToken *models.RefreshToken
	if err := as.db.Where("token=?", logoutReq.RefreshToken).Find(refreshToken).Error; err != nil {
		return err
	}
	as.db.Delete(refreshToken)

	return nil
}
