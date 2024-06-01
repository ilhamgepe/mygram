package services

import (
	"context"
	"log"

	"github.com/ilhamgepe/mygram/internal/models"
	"github.com/ilhamgepe/mygram/internal/repositories"
)

type UserService interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	RegisterUser(ctx context.Context, user *models.UserCreateReq) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.UserUpdateReq, id uint64) (*models.User, error)
	DeleteUser(ctx context.Context) error
}

type userService struct {
	ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{ur: ur}
}

func (us *userService) RegisterUser(ctx context.Context, user *models.UserCreateReq) (*models.User, error) {
	return us.ur.RegisterUser(ctx, user)
}

func (us *userService) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return us.ur.FindByEmail(ctx, email)
}

func (us *userService) UpdateUser(ctx context.Context, user *models.UserUpdateReq, id uint64) (*models.User, error) {
	log.Println("update user SERVICE")
	return us.ur.UpdateUser(ctx, user, id)
}

func (us *userService) DeleteUser(ctx context.Context) error {
	return us.ur.DeleteUser(ctx)
}
