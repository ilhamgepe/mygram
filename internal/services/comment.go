package services

import (
	"context"

	"github.com/ilhamgepe/mygram/internal/models"
	"github.com/ilhamgepe/mygram/internal/repositories"
)

type CommentService interface {
	AddComment(ctx context.Context, comment *models.CommentCreateReq) (*models.CommentCreateRes, error)
	GetComments(ctx context.Context) ([]models.GetAllCommentResponse, error)
	UpdateComment(ctx context.Context, comment *models.CommentUpdateReq, id uint64) (*models.CommentUpdateRes, error)
	DeleteComment(ctx context.Context, id uint64) error
}

type commentService struct {
	cr repositories.CommentRepository
}

func NewCommentService(cr repositories.CommentRepository) CommentService {
	return &commentService{cr: cr}
}

func (cs *commentService) AddComment(ctx context.Context, comment *models.CommentCreateReq) (*models.CommentCreateRes, error) {
	return cs.cr.AddComment(ctx, comment)
}

func (cs *commentService) GetComments(ctx context.Context) ([]models.GetAllCommentResponse, error) {
	return cs.cr.GetComments(ctx)
}

func (cs *commentService) UpdateComment(ctx context.Context, comment *models.CommentUpdateReq, id uint64) (*models.CommentUpdateRes, error) {
	return cs.cr.UpdateComment(ctx, comment, id)
}

func (cs *commentService) DeleteComment(ctx context.Context, id uint64) error {
	return cs.cr.DeleteComment(ctx, id)
}
