package services

import (
	"context"

	"github.com/ilhamgepe/mygram/internal/models"
	"github.com/ilhamgepe/mygram/internal/repositories"
)

type PhotoService interface {
	AddPhoto(ctx context.Context, photo *models.PhotoCreateReq) (*models.Photo, error)
	GetPhotos(ctx context.Context) ([]models.GetAllPhotoResponse, error)
	UpdatePhoto(ctx context.Context, photo *models.PhotoCreateReq, id uint64) (*models.UpdatePhotoResponse, error)
	DeletePhoto(ctx context.Context, id uint64) error
}

type photoService struct {
	pr repositories.PhotoRepository
}

func NewPhotoService(pr repositories.PhotoRepository) PhotoService {
	return &photoService{pr: pr}
}

func (ps *photoService) AddPhoto(ctx context.Context, photo *models.PhotoCreateReq) (*models.Photo, error) {
	return ps.pr.AddPhoto(ctx, photo)
}

func (ps *photoService) GetPhotos(ctx context.Context) ([]models.GetAllPhotoResponse, error) {
	return ps.pr.GetPhotos(ctx)
}

func (ps *photoService) UpdatePhoto(ctx context.Context, photo *models.PhotoCreateReq, id uint64) (*models.UpdatePhotoResponse, error) {
	return ps.pr.UpdatePhoto(ctx, photo, id)
}

func (ps *photoService) DeletePhoto(ctx context.Context, id uint64) error {
	return ps.pr.DeletePhoto(ctx, id)
}
