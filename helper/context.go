package helper

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/mygram/internal/models"
)

func CreateContextString(ctx *gin.Context, key string) (context.Context, error) {
	value, ok := ctx.Get(key)
	if !ok {
		return nil, errors.New(models.ErrUnauthorized)
	}
	return context.WithValue(ctx, key, value), nil
}
