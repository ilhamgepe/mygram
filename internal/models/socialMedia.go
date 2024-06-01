package models

import "time"

type SocialMedia struct {
	ID             uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Name           string    `json:"name" bind:"required" gorm:"column:name;not null"`
	SocialMediaUrl string    `json:"social_media_url" bind:"required" gorm:"column:social_media_url;not null"`
	UserId         uint64    `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User User `json:"user"`
}

type SocialMediaCreateReq struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

type SocialMediaUpdateReq struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

type SocialMediaCreateRes struct {
	ID             uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Name           string    `json:"name" bind:"required" gorm:"column:name;not null"`
	SocialMediaUrl string    `json:"social_media_url" bind:"required" gorm:"column:social_media_url;not null"`
	UserId         uint64    `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type SocialMediaUpdateRes struct {
	ID             uint64 `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Name           string `json:"name" bind:"required" gorm:"column:name;not null"`
	SocialMediaUrl string `json:"social_media_url" bind:"required" gorm:"column:social_media_url;not null"`
	UserId         uint64 `json:"user_id" gorm:"column:user_id;not null"`
	UpdatedAt      time.Time
}

type GetAllSocialMediaResponse struct {
	ID             uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Name           string    `json:"name" bind:"required" gorm:"column:name;not null"`
	SocialMediaUrl string    `json:"social_media_url" bind:"required" gorm:"column:social_media_url;not null"`
	UserId         uint64    `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User UserForSocialMedia `json:"User"`
}
