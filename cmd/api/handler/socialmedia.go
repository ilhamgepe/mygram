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

type SocialMediaHandler interface {
	AddSocialMedia(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHandler struct {
	ss services.SocialMediaService
}

func NewSocialMediaHandler(ss services.SocialMediaService) SocialMediaHandler {
	return &socialMediaHandler{
		ss: ss,
	}
}

func (sh *socialMediaHandler) AddSocialMedia(ctx *gin.Context) {
	var body *models.SocialMediaCreateReq
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

	socialMedia, err := sh.ss.AddSocialMedia(stdCtx, body)
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

	ctx.JSON(http.StatusOK, socialMedia)
}

func (sh *socialMediaHandler) GetSocialMedias(ctx *gin.Context) {
	socialMedias, err := sh.ss.GetSocialMedias(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"social_medias": socialMedias,
	})
}

func (sh *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	var body *models.SocialMediaUpdateReq
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

	socialMedia, err := sh.ss.UpdateSocialMedia(stdCtx, body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func (sh *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
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

	err = sh.ss.DeleteSocialMedia(stdCtx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
