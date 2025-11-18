package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iiiamnick/GOLANG--REST-API-EVENT-BOOKING-.git/models"
	"github.com/iiiamnick/GOLANG--REST-API-EVENT-BOOKING-.git/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not process the request"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save the user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "authentication failed"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate the user"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": token})

}
