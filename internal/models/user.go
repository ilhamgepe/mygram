package models

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Username  string    `json:"username" binding:"required" gorm:"index,column:username;unique;not null"`
	Email     string    `json:"email" binding:"required,email" gorm:"index,column:email;unique;not null"`
	Password  string    `json:"-" binding:"required,min=6" gorm:"column:password;not null"`
	Age       int       `json:"age" binding:"required,gt=8" gorm:"column:age"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserCreateReq struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,gt=8"`
}

type UserLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateReq struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
}

// for photo
type UserForPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// for comment
type UserForComment struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserForSocialMedia struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
