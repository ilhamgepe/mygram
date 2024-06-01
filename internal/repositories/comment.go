package repositories

import (
	"context"
	"errors"

	"github.com/ilhamgepe/mygram/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository interface {
	AddComment(ctx context.Context, comment *models.CommentCreateReq) (*models.CommentCreateRes, error)
	GetComments(ctx context.Context) ([]models.GetAllCommentResponse, error)
	UpdateComment(ctx context.Context, comment *models.CommentUpdateReq, id uint64) (*models.CommentUpdateRes, error)
	DeleteComment(ctx context.Context, id uint64) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (pr *commentRepository) AddComment(ctx context.Context, c *models.CommentCreateReq) (*models.CommentCreateRes, error) {
	var comment *models.Comment = &models.Comment{
		PhotoId: c.PhotoId,
		Message: c.Message,
		UserId:  ctx.Value("id").(uint64),
	}

	if err := pr.db.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &models.CommentCreateRes{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: comment.CreatedAt,
	}, nil
}

func (pr *commentRepository) GetComments(ctx context.Context) ([]models.GetAllCommentResponse, error) {
	// get all comments with join
	var comments []models.Comment
	if err := pr.db.Preload("User").Preload("Photo").Find(&comments).Error; err != nil {
		return nil, err
	}

	var responses []models.GetAllCommentResponse
	for _, comment := range comments {
		responses = append(responses, models.GetAllCommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoId:   comment.PhotoId,
			UserId:    comment.UserId,
			CreatedAt: comment.CreatedAt,
			User: models.UserForComment{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: models.PhotoForComment{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserId:   comment.Photo.UserId,
			},
		})
	}

	return responses, nil
}

func (pr *commentRepository) UpdateComment(ctx context.Context, comment *models.CommentUpdateReq, id uint64) (*models.CommentUpdateRes, error) {
	var c *models.Comment
	if err := pr.db.Preload("User").Preload("Photo").First(&c, id).Error; err != nil {
		return nil, err
	}

	if c.UserId != ctx.Value("id").(uint64) {
		return nil, errors.New(models.ErrForbidden)
	}

	c.Message = comment.Message

	if err := pr.db.Model(&c).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&c).Error; err != nil {
		return nil, err
	}

	return &models.CommentUpdateRes{
		ID:        c.ID,
		Title:     c.Photo.Title,
		Caption:   c.Photo.Caption,
		PhotoUrl:  c.Photo.PhotoUrl,
		UserId:    c.UserId,
		UpdatedAt: c.UpdatedAt,
	}, nil
}

func (pr *commentRepository) DeleteComment(ctx context.Context, id uint64) error {
	var c *models.Comment

	if err := pr.db.First(&c, id).Error; err != nil {
		return err
	}

	if c.UserId != ctx.Value("id").(uint64) {
		return errors.New(models.ErrForbidden)
	}

	if err := pr.db.Delete(&c).Error; err != nil {
		return err
	}
	return nil
}
