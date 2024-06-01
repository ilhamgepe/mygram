package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/mygram/helper"
	"github.com/ilhamgepe/mygram/internal/models"
	"github.com/ilhamgepe/mygram/internal/services"
)

type PhotoHandler interface {
	AddPhoto(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHandler struct {
	ps services.PhotoService
}

func NewPhotoHandler(ps services.PhotoService) PhotoHandler {
	return &photoHandler{
		ps: ps,
	}
}

func (ph *photoHandler) AddPhoto(ctx *gin.Context) {
	var body *models.PhotoCreateReq
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	stdCtx, err := helper.CreateContextString(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	photo, err := ph.ps.AddPhoto(stdCtx, body)
	if err != nil {
		if strings.Contains(err.Error(), "you are not allowed to access this resource") {
			ctx.JSON(http.StatusForbidden, gin.H{
				"errors": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserId,
		"created_at": photo.CreatedAt,
	})
}

func (ph *photoHandler) GetPhotos(ctx *gin.Context) {
	photos, err := ph.ps.GetPhotos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (ph *photoHandler) UpdatePhoto(ctx *gin.Context) {
	var body *models.PhotoCreateReq
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	stdCtx, err := helper.CreateContextString(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	idReq := ctx.Param("id")
	id, err := strconv.ParseUint(idReq, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	photo, err := ph.ps.UpdatePhoto(stdCtx, body, id)
	if err != nil {
		if strings.Contains(err.Error(), "you are not allowed to access this resource") {
			ctx.JSON(http.StatusForbidden, gin.H{
				"errors": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (ph *photoHandler) DeletePhoto(ctx *gin.Context) {
	stdCtx, err := helper.CreateContextString(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	idReq := ctx.Param("id")
	id, err := strconv.ParseUint(idReq, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	err = ph.ps.DeletePhoto(stdCtx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
