package repositories

import (
	"context"
	"errors"

	"github.com/ilhamgepe/mygram/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user *models.UserCreateReq) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.UserUpdateReq, id uint64) (*models.User, error)
	DeleteUser(ctx context.Context) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) RegisterUser(ctx context.Context, req *models.UserCreateReq) (*models.User, error) {
	var user *models.User = &models.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Age:      req.Age,
	}
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *models.UserUpdateReq, id uint64) (*models.User, error) {
	var existingUser models.User
	if err := ur.db.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	if existingUser.ID == ctx.Value("id") {
		existingUser.Email = user.Email
		existingUser.Username = user.Username

		if err := ur.db.Save(&existingUser).Error; err != nil {
			return nil, err
		}

		return &existingUser, nil
	}

	return nil, errors.New(models.ErrForbidden)

}

func (ur *userRepository) DeleteUser(ctx context.Context) error {
	var existingUser models.User
	if err := ur.db.Where("id = ?", ctx.Value("id")).First(&existingUser).Error; err != nil {
		return err
	}
	if existingUser.ID != ctx.Value("id") {
		return errors.New(models.ErrForbidden)
	}
	return ur.db.Delete(&existingUser).Error
}
