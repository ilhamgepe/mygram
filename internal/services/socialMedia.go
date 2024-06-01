package services

import (
	"context"

	"github.com/ilhamgepe/mygram/internal/models"
	"github.com/ilhamgepe/mygram/internal/repositories"
)

type SocialMediaService interface {
	AddSocialMedia(ctx context.Context, socialMedia *models.SocialMediaCreateReq) (*models.SocialMediaCreateRes, error)
	GetSocialMedias(ctx context.Context) ([]models.GetAllSocialMediaResponse, error)
	UpdateSocialMedia(ctx context.Context, socialMedia *models.SocialMediaUpdateReq, id uint64) (*models.SocialMediaUpdateRes, error)
	DeleteSocialMedia(ctx context.Context, id uint64) error
}

type socialMediaService struct {
	sr repositories.SocialMediaRepository
}

func NewSocialMediaService(sr repositories.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{sr: sr}
}

func (ss *socialMediaService) AddSocialMedia(ctx context.Context, s *models.SocialMediaCreateReq) (*models.SocialMediaCreateRes, error) {
	return ss.sr.AddSocialMedia(ctx, s)
}

func (ss *socialMediaService) GetSocialMedias(ctx context.Context) ([]models.GetAllSocialMediaResponse, error) {
	return ss.sr.GetSocialMedias(ctx)
}

func (ss *socialMediaService) UpdateSocialMedia(ctx context.Context, s *models.SocialMediaUpdateReq, id uint64) (*models.SocialMediaUpdateRes, error) {
	return ss.sr.UpdateSocialMedia(ctx, s, id)
}

func (ss *socialMediaService) DeleteSocialMedia(ctx context.Context, id uint64) error {
	return ss.sr.DeleteSocialMedia(ctx, id)
}
