package services

import (
	"github.com/KalyanKanuri/GoCart/internal/dto"
	"github.com/KalyanKanuri/GoCart/internal/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	Log *zerolog.Logger
}

func NewUserService(db *gorm.DB, log *zerolog.Logger) *UserService {
	return &UserService{
		db:  db,
		Log: log,
	}
}

func (us *UserService) GetProfile(userID uint) (*dto.UserResponse, error) {
	var User models.User

	if err := us.db.First(&User, userID).Error; err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        User.ID,
		Email:     User.Email,
		FirstName: User.FirstName,
		LastName:  User.LastName,
		Role:      string(User.Role),
		Phone:     User.Phone,
		IsActive:  User.IsActive,
	}, nil
}

func (us *UserService) UpdateProfile(userID uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	var user models.User

	if err := us.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Phone = req.Phone

	if err := us.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return us.GetProfile(userID)
}
