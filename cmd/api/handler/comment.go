package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/mygram/helper"
	"github.com/ilhamgepe/mygram/internal/models"
	"github.com/ilhamgepe/mygram/internal/services"
)

type CommentHandler interface {
	AddComment(ctx *gin.Context)
	GetComments(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHandler struct {
	cs services.CommentService
}

func NewCommentHandler(cs services.CommentService) CommentHandler {
	return &commentHandler{
		cs: cs,
	}
}

func (ch *commentHandler) AddComment(ctx *gin.Context) {
	var body *models.CommentCreateReq
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
	log.Println("in handler", body)
	comment, err := ch.cs.AddComment(stdCtx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (ch *commentHandler) GetComments(ctx *gin.Context) {
	comments, err := ch.cs.GetComments(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func (ch *commentHandler) UpdateComment(ctx *gin.Context) {
	var body *models.CommentUpdateReq
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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

	stdCtx, err := helper.CreateContextString(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	comment, err := ch.cs.UpdateComment(stdCtx, body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func (ch *commentHandler) DeleteComment(ctx *gin.Context) {
	idReq := ctx.Param("id")
	id, err := strconv.ParseUint(idReq, 10, 64)
	if err != nil {
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

	err = ch.cs.DeleteComment(stdCtx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
