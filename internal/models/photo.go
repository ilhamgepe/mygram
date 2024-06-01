package models

import "time"

type Photo struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Title     string    `json:"title" binding:"required" gorm:"column:title"`
	Caption   string    `json:"caption" gorm:"column:caption"`
	PhotoUrl  string    `json:"photo_url" binding:"required" gorm:"column:photo_url"`
	UserId    uint64    `json:"user_id" gorm:"column:user_id;not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId;references:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type GetAllPhotoResponse struct {
	ID        uint64       `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Title     string       `json:"title" binding:"required" gorm:"column:title"`
	Caption   string       `json:"caption" gorm:"column:caption"`
	PhotoUrl  string       `json:"photo_url" binding:"required" gorm:"column:photo_url"`
	UserId    uint64       `json:"user_id" gorm:"column:user_id;not null"`
	User      UserForPhoto `json:"user"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}
type UpdatePhotoResponse struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Title     string    `json:"title" binding:"required" gorm:"column:title"`
	Caption   string    `json:"caption" gorm:"column:caption"`
	PhotoUrl  string    `json:"photo_url" binding:"required" gorm:"column:photo_url"`
	UserId    uint64    `json:"user_id" gorm:"column:user_id;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type PhotoCreateReq struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

// for comment
type PhotoForComment struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	UserId   uint64 `json:"user_id"`
	PhotoUrl string `json:"photo_url"`
}
