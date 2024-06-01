package repositories

import (
	"context"
	"errors"

	"github.com/ilhamgepe/mygram/internal/models"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	AddSocialMedia(ctx context.Context, socialMedia *models.SocialMediaCreateReq) (*models.SocialMediaCreateRes, error)
	GetSocialMedias(ctx context.Context) ([]models.GetAllSocialMediaResponse, error)
	UpdateSocialMedia(ctx context.Context, socialMedia *models.SocialMediaUpdateReq, id uint64) (*models.SocialMediaUpdateRes, error)
	DeleteSocialMedia(ctx context.Context, id uint64) error
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (pr *socialMediaRepository) AddSocialMedia(ctx context.Context, s *models.SocialMediaCreateReq) (*models.SocialMediaCreateRes, error) {
	var socialMedia *models.SocialMedia = &models.SocialMedia{
		Name:           s.Name,
		SocialMediaUrl: s.SocialMediaUrl,
		UserId:         ctx.Value("id").(uint64),
	}

	if err := pr.db.Create(&socialMedia).Error; err != nil {
		return nil, err
	}

	return &models.SocialMediaCreateRes{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt,
	}, nil
}

func (pr *socialMediaRepository) GetSocialMedias(ctx context.Context) ([]models.GetAllSocialMediaResponse, error) {
	var socialMedia []models.SocialMedia
	if err := pr.db.Preload("User").Find(&socialMedia).Error; err != nil {
		return nil, err
	}

	var res []models.GetAllSocialMediaResponse
	for _, sm := range socialMedia {
		var s models.GetAllSocialMediaResponse = models.GetAllSocialMediaResponse{
			ID:             sm.ID,
			Name:           sm.Name,
			SocialMediaUrl: sm.SocialMediaUrl,
			UserId:         sm.UserId,
			CreatedAt:      sm.CreatedAt,
			UpdatedAt:      sm.UpdatedAt,
			User: models.UserForSocialMedia{
				ID:       sm.User.ID,
				Email:    sm.User.Email,
				Username: sm.User.Username,
			},
		}
		res = append(res, s)
	}
	return res, nil
}

func (pr *socialMediaRepository) UpdateSocialMedia(ctx context.Context, s *models.SocialMediaUpdateReq, id uint64) (*models.SocialMediaUpdateRes, error) {
	var socialMedia *models.SocialMedia
	if err := pr.db.First(&socialMedia, id).Error; err != nil {
		return nil, err
	}

	if socialMedia.UserId != ctx.Value("id").(uint64) {
		return nil, errors.New(models.ErrForbidden)
	}

	socialMedia.Name = s.Name
	socialMedia.SocialMediaUrl = s.SocialMediaUrl
	if err := pr.db.Save(&socialMedia).Error; err != nil {
		return nil, err
	}
	return &models.SocialMediaUpdateRes{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		UpdatedAt:      socialMedia.UpdatedAt,
	}, nil
}

func (pr *socialMediaRepository) DeleteSocialMedia(ctx context.Context, id uint64) error {
	var socialMedia *models.SocialMedia
	if err := pr.db.First(&socialMedia, id).Error; err != nil {
		return err
	}

	if socialMedia.UserId != ctx.Value("id").(uint64) {
		return errors.New(models.ErrForbidden)
	}

	if err := pr.db.Delete(&socialMedia).Error; err != nil {
		return err
	}
	return nil
}
