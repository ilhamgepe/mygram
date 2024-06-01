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

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userHandler struct {
	us services.UserService
}

func NewUserHandler(us services.UserService) UserHandler {
	return &userHandler{
		us: us,
	}
}

func (uh *userHandler) Register(ctx *gin.Context) {
	var body *models.UserCreateReq
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	// hasing password
	hashedPass, err := helper.HashPassword(body.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	body.Password = hashedPass
	user, err := uh.us.RegisterUser(ctx, body)

	if err != nil {
		if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": "user already exist",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var body *models.UserLoginReq
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	user, err := uh.us.FindByEmail(ctx, body.Email)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": models.ErrInvalidCredentials,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	if !helper.ComparePassword(user.Password, body.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": models.ErrInvalidCredentials,
		})
		return
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (uh *userHandler) Update(ctx *gin.Context) {
	var body *models.UserUpdateReq
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	id := ctx.Param("id")
	iduint64, err := strconv.ParseUint(id, 10, 64)
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

	user, err := uh.us.UpdateUser(stdCtx, body, iduint64)
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
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

func (uh *userHandler) Delete(ctx *gin.Context) {
	stdCtx, err := helper.CreateContextString(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	err = uh.us.DeleteUser(stdCtx)
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
		"message": "your account has been successfully deleted",
	})
}
