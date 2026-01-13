package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iiiamnick/EVENT-MANAGEMENT-REST_API/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	context.Next()
	context.Set("userId", userId)
}
