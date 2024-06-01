package models

const (
	// error message for user
	ErrUserNotFound = "user not found"

	// error message for photo
	ErrPhotoNotFound = "photo not found"

	// error message for comment
	ErrCommentNotFound = "comment not found"

	// error message for social media
	ErrSocialMediaNotFound = "social media not found"

	// error message for authentication & authorized
	ErrUnauthorized       = "unauthorized"
	ErrInvalidToken       = "invalid token"
	ErrExpiredToken       = "expired token"
	ErrInvalidCredentials = "invalid credentials"
	ErrForbidden          = "you are not allowed to access this resource"
)
