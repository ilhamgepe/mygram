package repositories

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/ilhamgepe/mygram/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository interface {
	AddPhoto(ctx context.Context, photo *models.PhotoCreateReq) (*models.Photo, error)
	GetPhotos(ctx context.Context) ([]models.GetAllPhotoResponse, error)
	UpdatePhoto(ctx context.Context, photo *models.PhotoCreateReq, id uint64) (*models.UpdatePhotoResponse, error)
	DeletePhoto(ctx context.Context, id uint64) error
}

type photoRepository struct {
	db *gorm.DB
	ur UserRepository
}

func NewPhotoRepository(db *gorm.DB, ur UserRepository) PhotoRepository {
	return &photoRepository{
		db: db,
		ur: ur,
	}
}

func (pr *photoRepository) AddPhoto(ctx context.Context, photo *models.PhotoCreateReq) (*models.Photo, error) {
	var user models.User
	if err := pr.db.Where("id = ?", ctx.Value("id")).First(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, errors.New(models.ErrForbidden)
		}
		return nil, err
	}
	log.Println(user)
	var p *models.Photo = &models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId:   user.ID,
	}
	if err := pr.db.Create(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (pr *photoRepository) GetPhotos(ctx context.Context) ([]models.GetAllPhotoResponse, error) {
	var photos []models.Photo
	if err := pr.db.Joins("User").Find(&photos).Error; err != nil {
		return nil, err
	}

	var res []models.GetAllPhotoResponse
	for _, photo := range photos {
		var p models.GetAllPhotoResponse = models.GetAllPhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserId:    photo.UserId,
			User:      models.UserForPhoto{Username: photo.User.Username},
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
		}
		res = append(res, p)
	}

	return res, nil
}

func (pr *photoRepository) UpdatePhoto(ctx context.Context, photo *models.PhotoCreateReq, id uint64) (*models.UpdatePhotoResponse, error) {
	var p *models.Photo = &models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}
	if err := pr.db.Model(&p).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&p).Error; err != nil {
		return nil, err
	}

	if p.UserId != ctx.Value("id") {
		return nil, errors.New(models.ErrForbidden)
	}

	var res *models.UpdatePhotoResponse = &models.UpdatePhotoResponse{
		ID:        p.ID,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoUrl:  p.PhotoUrl,
		UpdatedAt: p.UpdatedAt,
		UserId:    p.UserId,
	}
	log.Println(p)
	log.Println(res)
	return res, nil
}

func (pr *photoRepository) DeletePhoto(ctx context.Context, id uint64) error {
	var photo models.Photo
	if err := pr.db.Where("id = ?", id).First(&photo).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return errors.New(models.ErrPhotoNotFound)
		}
		return err
	}
	if photo.UserId != ctx.Value("id") {
		return errors.New(models.ErrForbidden)
	}

	return pr.db.Where("id = ?", id).Delete(&models.Photo{}).Error
}
