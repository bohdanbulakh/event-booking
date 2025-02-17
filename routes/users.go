package routes

import (
	"event-booking/models"
	"event-booking/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	exception := context.ShouldBindJSON(&user)

	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "InvalidBodyException"},
		)
		return
	}

	exception = user.Save()
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		_ = fmt.Errorf("InternalServerError\n%w", exception)
		return
	}

	user.Password = ""
	context.JSON(
		http.StatusOK,
		gin.H{"user": user},
	)
}

func login(context *gin.Context) {
	var user models.User

	exception := context.ShouldBindJSON(&user)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "InvalidBodyException"},
		)
		return
	}

	exception = user.ValidateCredentials()
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "UnauthorizedException"},
		)
		return
	}

	token, exception := utils.GenerateToken(
		user.Id,
		user.Email,
	)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": exception.Error()},
		)
		_ = fmt.Errorf("InternalServerError\n%w", exception)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"accessToken": token},
	)
}
