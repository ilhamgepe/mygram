package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/mygram/helper"
)

func WithAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")
		auth := strings.Split(authorization, " ")
		if len(auth) != 2 {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
			return
		}

		tokenReq := auth[1]

		// validate token
		token, err := helper.VerifyToken(tokenReq)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
			return
		}

		idString, err := token.Claims.GetSubject()
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
			})
			return
		}
		id, err := strconv.ParseUint(idString, 10, 64)
		// set id to context
		ctx.Set("id", id)

		ctx.Next()
	}
}
