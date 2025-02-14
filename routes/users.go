package routes

import (
	"event-booking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	exception := context.ShouldBindJSON(&user)

	if exception != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "InvalidBodyException"},
		)
		return
	}

	exception = user.Save()
	if exception != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		fmt.Println(exception)
		return
	}

	user.Password = ""
	context.JSON(
		http.StatusOK,
		gin.H{"user": user},
	)
}
