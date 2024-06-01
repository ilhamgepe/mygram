package models

import "time"

type Comment struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	UserId    uint64    `json:"user_id" gorm:"column:user_id;not null"`
	PhotoId   uint64    `json:"photo_id" gorm:"column:photo_id;not null"`
	Message   string    `json:"message" binding:"required" gorm:"column:message;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User  User  `json:"User"`
	Photo Photo `json:"Photo"`
}

type CommentCreateReq struct {
	PhotoId uint64 `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}
type CommentUpdateReq struct {
	Message string `json:"message" binding:"required"`
}

type CommentCreateRes struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Message   string    `json:"message" bind:"required" gorm:"column:message;not null"`
	PhotoId   uint64    `json:"photo_id" gorm:"column:photo_id;not null"`
	UserId    uint64    `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type CommentUpdateRes struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Title     string    `json:"title" binding:"required" gorm:"column:title"`
	Caption   string    `json:"caption" gorm:"column:caption"`
	PhotoUrl  string    `json:"photo_url" binding:"required" gorm:"column:photo_url"`
	UserId    uint64    `json:"user_id" gorm:"column:user_id;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type GetAllCommentResponse struct {
	ID        uint64          `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Message   string          `json:"message" binding:"required" gorm:"column:message;not null"`
	PhotoId   uint64          `json:"photo_id" gorm:"column:photo_id;not null"`
	UserId    uint64          `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	User      UserForComment  `json:"user"`
	Photo     PhotoForComment `json:"photo"`
}
